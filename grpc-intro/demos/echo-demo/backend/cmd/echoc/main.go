package main

import (
	"backend/api"
	"context"
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

var reverse bool
var host string

func init() {
	{ // -h and -host
		const usage = "host"
		flag.StringVar(&host, "host", "localhost:9000", usage)
		flag.StringVar(&host, "h", "localhost:9000", usage+" (shortcut)")
	}

	{ // -r and -reverse
		const usage = "reverse the input"
		flag.BoolVar(&reverse, "reverse", false, usage)
		flag.BoolVar(&reverse, "r", false, usage+" (shortcut)")
	}

	flag.Parse()
}

func main() {
	var opts []grpc.DialOption // TODO?
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial(host, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := api.NewEchoServiceClient(conn)

	msg := &api.EchoMessage{
		Value:   "Hello",
		Reverse: reverse,
	}

	fmt.Println("send:", msg.Value)

	msg, err = c.Echo(context.Background(), msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("recd:", msg.Value)
}
