build:
	goos=linux go build -o main main.go
zip:
	zip archive.zip main