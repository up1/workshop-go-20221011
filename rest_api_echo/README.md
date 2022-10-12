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