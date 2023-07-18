package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/tetigo/go-lessons-backend/internal/server"
)

var ginLambda *ginadapter.GinLambda

func init(){
	log.Println("Gin Cold Start")
	allowedOrigin, found := os.LookupEnv("ALLOWED_ORIGIN")
	if !found{
		log.Fatalf("env variable ALLOWED_ORIGIN was not found")
	}

	myServer, err := server.New(server.Config{Port: 9090, AllowedOrigin: allowedOrigin})
	if err != nil {
		log.Fatalf("impossible to create the server: %s", err)
	}
	
	// err=myServer.Run() // runs on 0.0.0.0:9090
	// if err!=nil{
	// 	log.Fatalf("impossible to start the server: %s", err)
	// }
	
	ginLambda = ginadapter.New(myServer.Engine)
	
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest)(events.APIGatewayProxyResponse, error){
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
	// myServer, err := server.New(server.Config{Port: 9090})
	// if err != nil {
	// 	log.Fatalf("impossible to create the server: %s", err)
	// }
	// myServer.Run() // runs on 0.0.0.0:9090
}

// linux: sudo lfof -i -P | grep LISTEN
// windows: netstat -ab
