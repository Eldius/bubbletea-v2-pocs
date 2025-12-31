
helpmenu:
	go run ./cmd/cli helpmenu

release:
	goreleaser.exe build --snapshot --clean
