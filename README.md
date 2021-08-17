## Golang simple API server

A simple Dockerised API server using Golang and PostgreSQL.


### Project Structure
```
    - internal
            |
            /app
                /api_helper.go
                /api.go
                /dal.go
                /models.go
                /repo.go
                /router.go
                /server.go
    - migration
            |
            /initial_migration.sql
    - docker-compose.yml
    - go.mod
    - go.sum
    - main.go

```

### Dependencies
> docker-compose == 1.29.2 \
Golang == 1.16

### Steps to Execute 
#### 1. Git clone golang-simple-api-server 
```bash
git clone {github url here} golang-simple-api-server
```
#### 2. goto golang-simple-api-server 
```bash
cd golang-simple-api-server
```
#### 3. start the postgres and pgweb docker 
```bash
docker-compose up -d
```
#### 4. Build the Application
```bash
go build -o app 
```
#### 5. Run the application
```bash
./app
```

Go to your browser : http://127.0.0.1:3000/api/

### Available API End Points

> GET   : http://127.0.0.1:3000/api/status <br>

> POST  : http://127.0.0.1:3000/api/category
#### Sample Body
```bash
{
    "name": "category1",
    "parent_category": 3
}
```

> GET   : http://127.0.0.1:3000/api/categories <br>

> POST  : http://127.0.0.1:3000/api/product
#### Sample Body
```bash
{
    "name": "product1",
    "price": 100,
    "category_ids": [
        1,
        2
    ]
}
```
> PUT   : http://127.0.0.1:3000/api/product
#### Sample Body
```bash
{
    "id": 1,
    "name": "Product1",
    "price": 1200
}
```
> GET   : http://127.0.0.1:3000/api//products/{category}