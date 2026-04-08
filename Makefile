APP_NAME := termtype

.PHONY: run build clean dev

run:
	go run .

build:
	go build -o $(APP_NAME) .

dev:
	@command -v air >/dev/null 2>&1 || go install github.com/air-verse/air@latest
	air

clean:
	rm -f $(APP_NAME)
