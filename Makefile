.PHONY: run
run:
	docker compose up -d
	eval $$(egrep -v '^#' .env | xargs) go run ./main.go 