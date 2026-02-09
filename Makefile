DESTDIR?=/
PREFIX=/usr

bin/dbus-wait: main.go
	mkdir -p bin
	go build -ldflags '-s'
	mv dbus-wait bin

build: bin/dbus-wait

install: build
	@install -Dm755 bin/dbus-wait ${DESTDIR}${PREFIX}/bin/dbus-wait

.PHONY: build install
