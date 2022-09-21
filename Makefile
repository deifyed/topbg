.PHONY=fmt
INSTALL_DIR=~/.local/bin

fmt:
	@gofmt -w .

topbg:
	go build -o topbg main.go

build: topbg

install: build
	mkdir -p $(INSTALL_DIR)
	cp topbg $(INSTALL_DIR)

clean:
	rm topbg
