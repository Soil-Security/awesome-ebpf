package main

import (
	"context"
	"fmt"
	"time"

	"github.com/falcosecurity/client-go/pkg/api/outputs"
	"github.com/falcosecurity/client-go/pkg/client"
)

func main() {
	// Set up a connection to Falco via a Unix socket
	c, err := client.NewForConfig(context.Background(), &client.Config{
		UnixSocketPath: "unix:///run/falco/falco.sock",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()
	// Subscribe to a stream of Falco notifications
	err = c.OutputsWatch(context.Background(), func(res *outputs.Response) error {
		// Put your business logic here
		fmt.Println(res.Output, res.OutputFields)
		return nil
	}, time.Second)
	if err != nil {
		panic(err)
	}
}
