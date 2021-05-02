build: internal/*
	go build -o build/hentri-client internal/main.go

run:
	go run internal/main.go

clean:
	rm -rf build/*
	rm -rf output/*
