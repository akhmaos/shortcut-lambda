set GOOS=linux

go build -o shorten cmd/main.go
zip -o deployment.zip shorten

aws lambda create-function --region us-east-1 --function-name ShortenFunction --zip-file fileb://./deployment.zip --runtime go1.x --tracing-config Mode=Active --role $ROLE --handler shorten
