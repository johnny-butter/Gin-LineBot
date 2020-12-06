FROM golang:1.15.2 AS builder

WORKDIR /builder

COPY . .

# Build migration tool
RUN go get github.com/gobuffalo/pop/... && \
    go build github.com/gobuffalo/pop/soda

# Build app
RUN go build -o line_bot_app


FROM golang:1.15.2

WORKDIR /app

COPY --from=builder /builder/heroku-release.sh .
COPY --from=builder /builder/soda .
COPY --from=builder /builder/line_bot_app .
COPY --from=builder /builder/database.yml.example database.yml

CMD ["/bin/bash"]
