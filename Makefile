MAKEFLAGS := --silent --always-make

# -r=false makes termination signals work properly
GOW := gow -r=false

dev: install
	$(MAKE) -j generate-watch run-watch

run-watch:
	$(GOW) -g=make run

run:
	go run .

generate-watch:
	$(GOW) -e templ -g make generate

generate:
	templ generate

	go mod tidy && go mod vendor

install:
	go install github.com/mitranim/gow@latest
	go install github.com/a-h/templ/cmd/templ@latest

prod:
	docker build -t 'seats-org' -f ./build/Dockerfile .
	docker run -p 8080:8080 'seats-org'
