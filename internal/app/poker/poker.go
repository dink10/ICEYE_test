package poker

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/dink10/poker/internal/pkg/config"
	"github.com/dink10/poker/internal/pkg/logger"
	"github.com/dink10/poker/internal/pkg/poker/domain/entity"
	"github.com/dink10/poker/internal/pkg/poker/domain/service"
)

const (
	newGameMessage = "Please, input cards for a new party (all hands without spaces)"
)

// Run runs application.
func Run() error {
	logrus.Info("Begin game")
	logrus.Info(newGameMessage)
	defer logrus.Info("End game")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var cfg Config
	err := config.LoadConfig(&cfg)
	if err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}

	err = logger.Init(&cfg.Logger)
	if err != nil {
		return fmt.Errorf("failed to initialize logger: %v", err)
	}

	go runPoker(ctx, cancel)

	select {
	case <-stop:
		return nil
	case <-ctx.Done():
		return nil
	}
}

func runPoker(ctx context.Context, cancel context.CancelFunc) {
	srv := service.NewService()

	hands := make([]string, 0, entity.HandsCount)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			return
		default:
			res := strings.Split(scanner.Text(), " ")
			hands = append(hands, res...)

			for _, hand := range hands {
				if strings.EqualFold(hand, "q") || strings.EqualFold(hand, "exit") {
					cancel()
				}
			}

			switch {
			case len(hands) > entity.HandsCount:
				logrus.Errorf("only %d players can be in play", entity.HandsCount)
				logrus.Info(newGameMessage)
				hands = make([]string, 0, entity.HandsCount)
				continue
			case len(hands) == entity.HandsCount:
				hs := make([]entity.Hand, 0, len(hands))
				var (
					valErr error
					h      entity.Hand
				)
				for i := range hands {
					h, valErr = srv.GetHandFromString(hands[i])
					if valErr != nil {
						break
					}

					hs = append(hs, h)
				}

				if valErr != nil {
					logrus.Error(valErr)
					hands = make([]string, 0, entity.HandsCount)
					continue
				}

				printWinner(srv.GetWinner(hs[0], hs[1]))
				hands = make([]string, 0, entity.HandsCount)
				logrus.Info(newGameMessage)
			}
		}
	}
}

func printWinner(position int) {
	if position == 0 {
		logrus.Println("Tie")
		return
	}
	logrus.Printf("Hand %d\n", position)
}
