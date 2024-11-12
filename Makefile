init:
	go mod tidy
	go generate ./ent

gen:
	go generate ./ent


run:
	go run main.go
