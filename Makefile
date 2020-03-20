all: build_x86_64 build_armv7

build_x86_64:
	go build -o build/gpbkdf2 -i .
	
build_armv7:
	GOOS=linux GOARCH=arm GOARM=7 go get . || true
	GOOS=linux GOARCH=arm GOARM=7 go build -o build/gpbkdf2_armv7 -i .

install_dependencies:
	go get -t .

test:
	go test -v
