FROM golang:1.24-bullseye AS build

COPY . /app

WORKDIR /app

RUN go build

FROM gcr.io/distroless/base AS deploy

COPY --from=build /app/notes-backend /app/notes-backend

WORKDIR /app

CMD ["/app/notes-backend"]

EXPOSE 8080