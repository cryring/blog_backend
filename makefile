export CGO_CXXFLAGS_ALLOW:=.*
export CGO_LDFLAGS_ALLOW:=.*
export CGO_CFLAGS_ALLOW:=.*

app:="blog_backend"

all: build

build:
	@echo "\033[32m <============== making app ${app} =============> \033[0m"
	go build -ldflags='-w -s' $(FLAGS) -o ./${app} ./cmd/backend

build_linux:
	@echo "\033[32m <============== making app ${app} =============> \033[0m"
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s' $(FLAGS) -o ./${app} ./cmd/backend
