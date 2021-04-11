build = build

all: clean compile

clean:
	rm -rf $(build)
	mkdir $(build)

compile:
	cd protoapp ; \
	GOROOT=/usr/lib/go-1.14 \
	GOPATH=/home/geremde/go \
	/usr/lib/go-1.14/bin/go build -ldflags="-s -w" -o ../$(build)/test main.go

install:
	cd protoapp ; \
	GOROOT=/usr/lib/go-1.14 \
	GOPATH=/home/geremde/go \
	/usr/lib/go-1.14/bin/go install main.go

run:
	$(build)/test
