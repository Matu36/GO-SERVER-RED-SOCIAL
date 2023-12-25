git add .
git commit -m "Subiendo a lambda"
git push
GOARCH=amd64 GOOS=linux go build -a -o ./build/main main.go
@REM go build -o main main.go
del main.zip      
tar.exe -a -cf main.zip main 



