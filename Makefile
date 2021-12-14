build:
	@go build -o kreativroni ./app

wasm:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm ./app

run: build wasm
	export API_KEY="" && ./kreativroni