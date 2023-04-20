# v2 todo-list server 

- echo  



```bash 
mkdir v2-todo-list-server 
cd v2-todo-list-server 

go mod init v2-todo-list-server 

go get -u github.com/labstack/echo/v4


# build
make build.local 
``` 



**구조 참고**  
https://github.com/golang-standards/project-layout


**구조**  
```
- api 
- api_docs
- assets 
- cmd
  - prod 
  - dev 
  - local
- configs
- deployments       
- init 
- internal
- pkg 
- scripts 
- test 
- third_party 
- tools 
```