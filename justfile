build:
	cd agent && go build -ldflags="-s -w" -trimpath main.go