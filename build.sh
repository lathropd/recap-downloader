#!/bin/bash
GOOS=darwin GOARCH=amd64 go build -o ./bin/recap_downloader_amd64 main.go;
GOOS=darwin GOARCH=arm64 go build -o ./bin/recap_downloader_arm64 main.go;
GOOS=windows GOARCH=amd64 go build -o ./bin/recap_downloader.exe  main.go;

lipo -create -output ./bin/recap_downloader_universal ./bin/recap_downloader_arm64 ./bin/recap_downloader_amd64
