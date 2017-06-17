package main

import (
	"context"
	"fmt"

	"google.golang.org/api/iterator"

	"cloud.google.com/go/logging"
	"cloud.google.com/go/logging/logadmin"
)

func main() {
	// Create a Client
	ctx := context.Background()
	client, err := logging.NewClient(ctx, "api-project-661531736098")
	if err != nil {
		fmt.Println("err settings up client : ", err)
	}

	// Initialize a logger
	lg := client.Logger("my-log")

	// Add entry to log buffer
	lg.Log(logging.Entry{Payload: "something happened!"})

	// Close the client when finished.
	err = client.Close()
	if err != nil {
		fmt.Println("err closing client : ", err)
	}

	adminClient, err := logadmin.NewClient(ctx, "api-project-661531736098")
	if err != nil {
		fmt.Println("err settings up client : ", err)
	}

	fmt.Println("pulling metrics")
	//it := adminClient.Entries(ctx, logadmin.Filter(`logName = "projects/my-project/logs/my-log"`))
	it := adminClient.Entries(ctx)
	//it := adminClient.Metrics(ctx)
	fmt.Println(it)
	for {
		metric, err := it.Next()
		fmt.Println(metric)
		if err == iterator.Done {
			fmt.Println("iterator.Done")
			break
		}
		if err != nil {
			fmt.Println("err metric : ", err)
		}
		//fmt.Println(metric)
	}
}
