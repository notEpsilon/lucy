clean:
	@rm -rf build
	@rm -rf *.exe
	@rm -rf lucy.*
	@rm -rf output_file.*

format:
	@gofmt -s -w .

build: clean format
	@mkdir build
	@CGO_ENABLED=0 go build -ldflags="-s -w" -o build

run: build
	@./build/lucy
