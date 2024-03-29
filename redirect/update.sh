set GOOS=linux

go build -o redirect cmd/main.go
zip -o deployment.zip redirect

aws lambda update-function-code --function-name RedirectFunction --region us-east-1 --zip-file fileb://./deployment.zip
