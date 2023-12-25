git add .
git commit -m "Subiendo a lambda"
git push
build-main-zip
del main.zip      
tar.exe -a -cf main.zip main 



