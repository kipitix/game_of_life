
run:
	go run ./cmd --grid-width=100 --grid-height=100 --life-cycle-duration=100ms

help:
	go run ./cmd --help

test:
	go test ./...

benchmark:
	go test ./internal/domain/iterator -bench=. -cpuprofile ./cpu.prof
	go tool pprof ./cpu.prof

build:
	mkdir -p ./bin
	go build -o ./bin/game_of_life ./cmd
