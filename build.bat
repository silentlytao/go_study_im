rd /s/q release
del /f chat.exe
md release
::go build -ldflags "-H windowsgui" -O chat.exe
set GO111MODULE=off
go build -o chat.exe

COPY chat.exe release\
COPY favicon.ico release\favicon.ico

XCOPY asset\*.* release\asset\ /s /e
XCOPY view\*.* release\view\ /s /e