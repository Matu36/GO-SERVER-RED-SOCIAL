git add .
git commit -m "Subiendo a lambda"
git push
go build main.go
del main.zip      
tar.exe -a -cf main.zip main 