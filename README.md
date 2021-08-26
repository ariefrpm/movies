### this is example of movies microservices using Clean Architecture Principles

![alt text](https://github.com/ariefrpm/movies/raw/master/movies-diagram.png)


### Test

```bash
go test ./...
```

### Build

```bash
go build
```

### Run

```bash
go run .
```

### Rest API 
```bash
[GET] http://localhost:8080/api/search_movie?pagination2=1&searchword=Batman
[GET] http://localhost:8080/api/movie_detail?i=tt0372784
```

### GRPC Client
```bash
go run sandbox/main.go
```
