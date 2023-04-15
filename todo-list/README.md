# todo-list + chatting 

기능 참고  
https://to-do.live.com/ 


go  
https://go.dev/src/


gin   
https://gin-gonic.com/docs/


gorm   
https://gorm.io/docs/ 



```bash 
mkdir todo-list
cd todo-list 

go mod init todo-list 
go get -u github.com/joho/godotenv github.com/gin-gonic/gin gorm.io/gorm github.com/gin-contrib/cors github.com/google/uuid 

go run main.go
# http://localhost:8800/ 
``` 



**구조**
```
- common 
- middleware 
- mvc 
  - controller 
  - model          
  - view        # *.html          
- router 
- main.go
```



**버전**
- v1 : gin template + ajax 
- v2 : svelte.js + api 
