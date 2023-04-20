.PHONY: test build clean

EXECUTABLE=v2-todo-list-server

LINUX=$(EXECUTABLE)_linux
DARWIN=$(EXECUTABLE)_darwin
WINDOWS=$(EXECUTABLE)_windows.exe

TARGET=build

LOCAL=cmd/local/main.go

test:
	scripts/test.sh

build.local: clean
	@echo ===== Build Start =====
	env GOOS=linux GOARCH=amd64 go build -o ${TARGET}/${LINUX} $(LOCAL)
	env GOOS=darwin GOARCH=amd64 go build -o ${TARGET}/${DARWIN} $(LOCAL)
	env GOOS=windows GOARCH=amd64 go build -o ${TARGET}/${WINDOWS} $(LOCAL)
	go build -o bin/main $(LOCAL)
	@echo ===== Build End =====

clean:
	go clean
	rm -f ${TARGET}/${LINUX}
	rm -f ${TARGET}/${DARWIN}
	rm -f ${TARGET}/${WINDOWS}

run:
	${EXECUTABLE}/${TARGET}

all: build run 