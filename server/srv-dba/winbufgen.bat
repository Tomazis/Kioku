@echo off
setlocal
set SERVICE_NAME=srv-dba
set SERVICE_PATH=tomazis/kioku/server/srv-dba

echo Generate protobuf
buf generate

echo Move generated files
move pkg\%SERVICE_NAME%\github.com\%SERVICE_PATH%\pkg\%SERVICE_NAME%\* pkg\%SERVICE_NAME%\

echo Delete unnecessary generated files
rmdir /s /q pkg\%SERVICE_NAME%\github.com\

cd pkg\%SERVICE_NAME%
IF NOT EXIST go.mod (
    echo Make go.mod
    go mod init github.com/%SERVICE_PATH%/pkg/%SERVICE_NAME%
    go mod tidy
)

endlocal