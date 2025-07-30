FROM golang:1.24.1-alpine AS build

WORKDIR /app

COPY . ./
COPY .env .env
RUN go mod download

RUN go build -o /bin/app

FROM alpine:latest

COPY --from=build /bin/app /bin
COPY --from=build /app/.env ./.env

EXPOSE 4000

CMD [ "/bin/app" ]