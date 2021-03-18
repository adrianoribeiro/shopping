# Shopping (WIP)

It is a Product CRUD (simple test) for known about Go frameworks.  

## Frameworks/Tools
[Fiber](https://github.com/gofiber/fiber) (Web Frameworks + Routers)  
[Gorm](https://github.com/go-gorm/gorm) (ORMs)  
[MySQL](https://www.mysql.com/) (Relational Database)  
[Redis](https://redis.io/) (NoSQL, in-memory data structure store)
## How to run
### Launch a docker image
```
docker run --name some-mysql -p3306:3306 -e MYSQL_DATABASE=shoppingDB -e MYSQL_ROOT_PASSWORD=root -d mysql:latest
```

### If you want you can access docker container
```
docker exec -it <CONTAINER_ID> bash
```
and
```
mysql -u root -proot
```
### To launch the redis
```
docker network create -d bridge my-bridge-network
docker run --name some-redis --network my-bridge-network -p6379:6379 -d redis
```

### If you want to check the redis
```
docker run -it --network my-bridge-network --rm redis redis-cli -h some-redis
```

### Finally, to launch the app, run this command in the root folder (Of course you must setup the Go environment before in you computer )
```
go run main.go
```

### Endpoints
#### Login
- (POST) http://localhost:8080/login
#### Product
- (GET) http://localhost:8080/products  
- (GET) http://localhost:8080/products/:id  
- (POST) http://localhost:8080/products  
- (DELETE) http://localhost:8080/products/:id
#### Cart
- (GET) http://localhost:8080/cart
- (POST) http://localhost:8080/cart/add
#### Payment
- (POST) http://localhost:8080/payment

## Docker Redis commands
```
docker network create -d bridge my-bridge-network

docker run --name some-redis --network my-bridge-network -p6379:6379 -d redis

docker run -it --network my-bridge-network --rm redis redis-cli -h some-redis
```
