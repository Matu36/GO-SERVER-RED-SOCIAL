git add .
git commit -m "Subiendo a lambda"
git push
go build -o main main.go
del main.zip      
tar.exe -a -cf main.zip main 



