## Step to run

### Download dependencies
```
$go mod tidy
```

### Start server
```
$go run main.go
```


## Build with docker
```
$docker image build -t api:1.0 .
```


### Testing with tags = integration
```
$go test -tags=integration ./... -v
```
