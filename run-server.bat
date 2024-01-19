@REM build and run go server

cd %~dp0

cd server
go build -o server.exe server.go
server.exe

cd %~dp0