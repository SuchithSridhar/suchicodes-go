run: gen
	@go run cmd/main.go

gen:
	@templ generate
