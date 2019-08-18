test:
	GOARM=7 GOARCH=arm GOOS=linux go test examples/blink_test.go