FROM golang:1.16-alpine as builder

WORKDIR /go/build/
RUN mkdir app
COPY app/* app/
COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go build -o kreativroni ./app
RUN GOARCH=wasm GOOS=js go build -o web/app.wasm ./app

FROM golang:1.16-alpine
RUN apk --no-cache add ca-certificates

WORKDIR /app/
RUN mkdir -p web/css
RUN mkdir -p web/images

COPY web/css/main.css web/css/.
COPY web/images/logo_*.png web/images/
COPY --from=builder /go/build/kreativroni .
COPY --from=builder /go/build/web/app.wasm web/.

EXPOSE 8080

CMD ["./kreativroni"]
