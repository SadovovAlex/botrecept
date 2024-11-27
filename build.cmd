SET CGO_ENABLED=0
echo Change to the directory where the Go code is located
cd c:\PROGRAMMING\bot\botrecept\cmd\bot
echo Build the Go application
go build -o ..\..\botrecept.exe main.go setup.go

echo Unix build
SET GOOS=linux
SET GOARCH=amd64
SET CGO_ENABLED=0
go build -o ..\..\botrecept -tags sqlite_omit_load_extension main.go setup.go
