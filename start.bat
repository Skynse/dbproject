@echo off

rem Run Go server in the background
go build -o api.exe main.go
start /B api.exe > go_output.log 2>&1

rem Run Deno server in the background
cd web\static\js
start /B deno task start --allow-net --allow-read

rem Wait for user input to terminate
pause
