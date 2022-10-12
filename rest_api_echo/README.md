## Step to run

### Download dependencies
```
$go mod tidy
```

### Start server
```
$go run main.go
```
### Testing with tags = integration
```
$go test -tags=integration ./... -v
```

## Build with docker
```
$docker image build -t api:1.0 .
$docker container run -d -p 1323:1323 api:1.0
```

## Working with Docker compose
```
$docker-compose build
$docker-compose up -d
$docker-compose ps
$docker-compose logs --follow
$docker-compose down
```
