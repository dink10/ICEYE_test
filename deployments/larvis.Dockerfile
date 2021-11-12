FROM alpine:3.7

COPY ./cmd/larvis .

RUN chmod +x /larvis

CMD ["/larvis"]
