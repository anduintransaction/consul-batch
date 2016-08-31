package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
)

func printUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s OPTIONS ACTION ARGS\n", os.Args[0])
	fmt.Fprintln(os.Stderr, "ACTION: put get version")
	fmt.Fprintln(os.Stderr, "OPTIONS: ")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	var host string
	flag.StringVar(&host, "consul", "127.0.0.1:8500", "Consul API end point")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		printUsage()
	}

	client, err := api.NewClient(&api.Config{
		Address: host,
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	kv := client.KV()
	switch args[0] {
	case "put":
		cmdPut(kv, args[1:])
	case "get":
		cmdGet(kv, args[1:])
	case "version":
		cmdVersion()
	default:
		printUsage()
	}
}
