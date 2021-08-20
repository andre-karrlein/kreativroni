build:
	@go build -o kreativroni ./app

wasm:
	@GOARCH=wasm GOOS=js go build -o web/app.wasm ./app

run: build wasm
	export GOOGLE_APPLICATION_CREDENTIALS="./kreativroni.json" && ./kreativroni