run:
	echo "Run triggered"
	echo "Golang Rocks"
	go run ./bin/api.exe

build:
	echo "Building for Windows"
	# env GOOS="windows" go build -o bin/api.exe ./api/main.go
	env GOOS="linux" go build -o bin/api ./api/main.go
	env GOOS="linux" go build -o bin/hello ./testLambda/main.go

deploy-dev: build
	serverless deploy --aws-profile TiagoDeveloper

deploy-stage: build
	serverless deploy --aws-profile TiagoDeveloper --stage staging 

deploy-prod: build
	serverless deploy --aws-profile TiagoDeveloper --stage production
