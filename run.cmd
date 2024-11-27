SET CGO_ENABLED=0

echo Change to the directory where the Go code is located
cd c:\PROGRAMMING\bot\botrecept\cmd\bot

echo Build the Go application
go build -o ..\..\botrecept.exe main.go setup.go


echo Run the application with the environment variables file
cd c:\PROGRAMMING\bot\botrecept
botrecept.exe

pause
