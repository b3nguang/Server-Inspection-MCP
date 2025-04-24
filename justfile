# 默认编译当前系统的版本
build:
	cd agent && go build -ldflags="-s -w" -trimpath main.go

# 编译Windows 64位版本
build-windows-amd64:
	cd agent && set CGO_ENABLED=0 && set GOOS=windows && set GOARCH=amd64 && go build -ldflags="-s -w" -trimpath -o ../build/agent_windows_amd64.exe main.go

# 编译Windows 32位版本
build-windows-386:
	cd agent && set CGO_ENABLED=0 && set GOOS=windows && set GOARCH=386 && go build -ldflags="-s -w" -trimpath -o ../build/agent_windows_386.exe main.go

# 编译Linux 64位版本
build-linux-amd64:
	cd agent && set CGO_ENABLED=0 && set GOOS=linux && set GOARCH=amd64 && go build -ldflags="-s -w" -trimpath -o ../build/agent_linux_amd64 main.go

# 编译Linux 32位版本
build-linux-386:
	cd agent && set CGO_ENABLED=0 && set GOOS=linux && set GOARCH=386 && go build -ldflags="-s -w" -trimpath -o ../build/agent_linux_386 main.go

# 编译MacOS 64位版本
build-darwin-amd64:
	cd agent && set CGO_ENABLED=0 && set GOOS=darwin && set GOARCH=amd64 && go build -ldflags="-s -w" -trimpath -o ../build/agent_darwin_amd64 main.go

# 编译所有平台版本
build-all: build-windows-amd64 build-windows-386 build-linux-amd64 build-linux-386 build-darwin-amd64

# 使用PowerShell编译Linux 64位版本（适用于Windows 10及以上版本）
build-linux-amd64-ps:
	cd agent && powershell -Command "$env:CGO_ENABLED='0'; $env:GOOS='linux'; $env:GOARCH='amd64'; go build -ldflags='-s -w' -trimpath -o ../build/agent_linux_amd64 main.go"