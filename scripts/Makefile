APP_BINNAME:=app
APP_FILES:=../cmd/app/*

CLIENT_BINNAME:=cli
CLIENT_FILES:=../cmd/cli/*

DB_FILES:=../cmd/db/*

BINDIR:=../bin
GOARCH:=amd64

setup_db:
	go run $(DB_FILES)

build: build_app build_cli

build_app:
	GOOS=windows GOARCH=$(GOARCH) go build -o $(BINDIR)/$(APP_BINNAME).exe $(APP_FILES)
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BINDIR)/$(APP_BINNAME) $(APP_FILES)

build_cli:
	GOOS=windows GOARCH=$(GOARCH) go build -o $(BINDIR)/$(CLIENT_BINNAME).exe $(CLIENT_FILES)
	GOOS=linux GOARCH=$(GOARCH) go build -o $(BINDIR)/$(CLIENT_BINNAME) $(CLIENT_FILES)

test: build_cli
	go run $(APP_FILES)

clean:
	rm -rf ../bin

.PHONY: setup_db build build_app build_cli test clean
