build:
	GOOS=linux GOARCH=amd64 go build -o main main.go
zip:
	zip archive.zip main
deploy:
	aws lambda update-function-code --function-name GolangLambdaFunction --zip-file fileb://archive.zip
