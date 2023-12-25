git add .
git commit -m "Subiendo a lambda"
git push
GOOS=linux GOARCH=amd64 go build -o main main.go
del main.go.zip      
tar.exe -a -cf main.zip main 



