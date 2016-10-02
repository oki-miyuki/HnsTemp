@echo off
set GOROOT=D:\go
echo 'set GO install directory : %GOROOT%'

set GOPATH=D:\HnsTemp

echo 'set Build directory : %GOPATH%'
echo 'set GOARCH win32 application'

set GOARCH=386

echo 'set GOARCH : %GOARCH%'

set http_proxy=10.3.0.254:8080

set https_proxy=%http_proxy%

echo 'set proxy settings : %http_proxy%'
echo "-------------------------------------------------"
echo Download external source for windows service build

go get golang.org/x/sys/windows/svc

