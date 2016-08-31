package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/consul/api"
)

func cmdGet(kv *api.KV, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "USAGE: %s get <prefix>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	pairs, _, err := kv.List(args[0], nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	values := make(map[string]string)
	for _, p := range pairs {
		values[p.Key] = string(p.Value)
	}
	b, err := json.MarshalIndent(values, "", "    ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(string(b))
}
