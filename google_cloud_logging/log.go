package main

import (
	"context"
	"os"

	"cloud.google.com/go/logging"

	"github.com/codercom/go-util/loggers"

	"fmt"

	"go.uber.org/zap"
)

// follow these to setup
// https://developers.google.com/identity/protocols/application-default-credentials
// https://unix.stackexchange.com/questions/117467/how-to-permanently-set-environmental-variables
const (
	defaultEnvCredentialFilePath = "GOOGLE_APPLICATION_CREDENTIALS"
	defaultEnvPrivateKey         = "GOOGLE_API_GO_PRIVATEKEY"
	defaultEnvEmail              = "GOOGLE_API_GO_EMAIL"
)

var (
	envCredential string
	envEmail      string
	envPrivateKey string
)

func init() {
	envCredential = os.Getenv(defaultEnvCredentialFilePath)
	envPrivateKey = os.Getenv(defaultEnvPrivateKey)
	envEmail = os.Getenv(defaultEnvEmail)
}

func main() {
	fmt.Println("verson : ", os.Getenv("HOME"))
	log, err := zap.NewDevelopment(zap.Hooks())
	if err != nil {
		fmt.Println(err)
		return
	}

	client, err := logging.NewClient(context.Background(), "api-project-661531736098")
	if err != nil {
		panic(err)
	}

	log = log.WithOptions(zap.Hooks((&loggers.ZapGoogle{
		Logger: client.Logger("clide"),
	}).Log))

	log.Info("starting up1", zap.String("version", os.Getenv("HOME")))

	log.Sync()
	//client.Close()
	// fmt.Println(envCredential)
	// fmt.Println(envPrivateKey)
	// fmt.Println(envEmail)

	// // Create a Client
	// ctx := context.Background()
	// client, err := logging.NewClient(ctx, "api-project-661531736098")
	// if err != nil {
	// 	fmt.Println("err settings up client : ", err)
	// }

	// // Initialize a logger
	// lg := client.Logger("my-log")

	// // Add entry to log buffer

	// lg.Log(logging.Entry{Payload: "something happened!"})

	// // Close the client when finished.
	// err = client.Close()
	// if err != nil {
	// 	fmt.Println("err closing client : ", err)
	// }

	// adminClient, err := logadmin.NewClient(ctx, "api-project-661531736098")
	// if err != nil {
	// 	fmt.Println("err settings up client : ", err)
	// }

	// fmt.Println("pulling metrics")
	// //it := adminClient.Entries(ctx, logadmin.Filter(`logName = "projects/my-project/logs/my-log"`))
	// it := adminClient.Entries(ctx)
	// //it := adminClient.Metrics(ctx)
	// fmt.Println(it)
	// for {
	// 	metric, err := it.Next()
	// 	fmt.Println(metric)
	// 	if err == iterator.Done {
	// 		fmt.Println("iterator.Done")
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println("err metric : ", err)
	// 	}
	// 	//fmt.Println(metric)
	// }
}
