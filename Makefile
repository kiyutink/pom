.PHONY: build-macos-bundle
build-macos-bundle: 
	cp -r ./Pom.app.example ./Pom.app
	go build -o "./Pom.app/Contents/MacOS/Pom" .
