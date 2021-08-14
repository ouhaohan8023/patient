## Patient Registration Problem

### Live Demo
1. [Click To Register Page](https://done.ohh.ink/register) 
2. [Click To Admin Login](https://done.ohh.ink/login) admin admin

### Preview
![Register](https://raw.githubusercontent.com/ouhaohan8023/patient/main/preview/1.png)
![Register Validate](https://raw.githubusercontent.com/ouhaohan8023/patient/main/preview/2.png)
![Register Successful](https://raw.githubusercontent.com/ouhaohan8023/patient/main/preview/3.png)
![Admin Login](https://raw.githubusercontent.com/ouhaohan8023/patient/main/preview/4.png)
![Registerion Slips](https://raw.githubusercontent.com/ouhaohan8023/patient/main/preview/5.png)

### Run Locally

1. Install Go, Yarn
2. git clone this repo
3. cd in this repo
4. change the setting in the file /configs/app.yaml
```yaml
environment: dev
# Register Frequency Limit (Same Ip) 5/minute
requestLimit: 5
dev:
  database:
    host: 192.168.10.10
    port: 3306
    username: homestead
    password: secret
    db: patient
    charset: utf8
  redis:
    host: 127.0.0.1:6379
    username: ""
    password: ""
  # where the static photo exists
  baseUrl: "http://localhost:8080/"
  # where to save the photo
  uploadPath: "./uploads"
prod:
  database:
    host: 
    port: 
    username: 
    password: 
    db: 
    charset: 
  redis:
    host: 
    username: 
    password: 
  baseUrl: ""
  uploadPath: ""
log:
  http: true
```
5. create database
6. run backend project
```bash
# for local test, will start at port 8080
go run main.go
# build for linux
CGO_ENABLED=0  GOOS=linux  GOARCH=amd64  go build  main.go
```
7. run front project
```bash
yarn && yarn serve
#  will start at port 8081
```

### Already Done

1. Register
   1. Validate at frontend
   2. Validate at backend
   3. Upload photo and preview
   4. Use redis to controller frequency of requests from same ip
2. Admin Panel
   1. Login
   2. Registration List
   3. Search with name, phone and email
   4. Add "status" field to distinguish slip which is "Contacted" or not
   5. Update slip "status"
   
### TODO
1. Get "admin" user from DB
2. Unit test
3. ...