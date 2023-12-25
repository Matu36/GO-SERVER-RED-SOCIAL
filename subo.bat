git add .
git commit -m "Subiendo a lambda"
git push
env GOOS=linux GOARCH=amd64 go build -o main main.go


@REM go build -o main main.go
del main.zip      
tar.exe -a -cf main.zip main 



