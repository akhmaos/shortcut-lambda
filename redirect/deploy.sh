set GOOS=linux

go build -o redirect cmd/main.go
zip -o deployment.zip redirect

aws lambda create-function --function-name RedirectFunction --region us-east-1 --zip-file fileb://./deployment.zip --runtime go1.x --tracing-config Mode=Active --role $ROLE --handler redirect
