package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/consul/api"
)

func cmdPut(kv *api.KV, args []string) {
	if len(args) < 1 {
		fmt.Fprintf(os.Stderr, "USAGE: %s put <file>\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	var values map[string]string
	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = json.Unmarshal(content, &values)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	ops := api.KVTxnOps{}
	for k, v := range values {
		op := &api.KVTxnOp{
			Verb:  string(api.KVSet),
			Key:   k,
			Value: []byte(v),
		}
		ops = append(ops, op)
	}
	_, _, _, err = kv.Txn(ops, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
