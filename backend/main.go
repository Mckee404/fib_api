package main

import (
	"context"
	"errors"
	"math/big"
	"net/http"
	"strconv"

	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gin-gonic/gin"
)

// エラー定義
var (
	ErrInvalidInput = errors.New("n must be greater than or equal to 0")
	ErrBadRequest   = errors.New("Bad Request")
)

var ginLambda *ginadapter.GinLambda

func init() {
	router := gin.Default()
	router.GET("/fib", handleFibonacci)
	ginLambda = ginadapter.New(router)
}

func handleFibonacci(c *gin.Context) {
	n, err := strconv.Atoi(c.Query("n"))
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, ErrBadRequest.Error())
		return
	}

	result, err := fibonacci(n)
	if err != nil {
		sendErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func sendErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}

// fibonacci は n 番目のフィボナッチ数を計算する
func fibonacci(n int) (*big.Int, error) {
	if n < 0 {
		return nil, ErrInvalidInput
	}

	fib1, fib2 := big.NewInt(0), big.NewInt(1)
	if n == 0 {
		return fib1, nil
	} else if n == 1 {
		return fib2, nil
	}

	for i := 2; i <= n; i++ {
		fib1, fib2 = fib2, new(big.Int).Add(fib1, fib2)
	}
	return fib2, nil
}
