# Pom
Pom is a simple statusbar application for macOS that implements Pomodoro.

# Installation and usage
You have to options:
1. Start pom from the shell by just building and running the app
```
go build -o pom .
./pom
```
2. Build a macOS bundle and run that
```
make build-macos-bundle
# This will create `Pom.app` bundle. Move it to your applications directory and run from there
```
