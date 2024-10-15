build-webapp:
	cd webapp && npm run build && cd ..

build-arm64:
	go generate ./...
	GOOS=linux GOARCH=arm64 go build ${LDFLAGS} -o observability-rest-server-linux-arm64 cmd/observability-rest-server/main.go
	shasum -a 256 observability-rest-server-linux-arm64 > observability-rest-server-linux-arm64.sha256
	shasum -a 256 restserver-linux-arm64 > restserver-linux-arm64.sha256
	shasum -a 256 reset-admin-password-linux-arm64 > reset-admin-password-linux-arm64.sha256
