.PHONY: run

run:
	@find . -name '*.go' -or -name '*.templ' | entr -r make _run

_run:
	@templ generate
	@go run cmd/main.go