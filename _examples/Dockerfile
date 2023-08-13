FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

ADD https://github.com/michimani/evidently-local/archive/refs/tags/v0.0.4.zip ./

RUN unzip v0.0.4.zip \
  && cd evidently-local-0.0.4 \
  && go install \
  && go build -o evidently-local . \
  && mv evidently-local /app

# for run stage
FROM alpine:3.18.2

WORKDIR /app

COPY --from=builder /app/evidently-local .

RUN mkdir data

# from your local data directory
ADD ./data ./data

EXPOSE 2306

CMD [ "/app/evidently-local" ]