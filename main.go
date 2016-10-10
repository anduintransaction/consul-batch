package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
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

type ConsulConfig struct {
	Scheme   string `json:"scheme"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func readConfig() (*ConsulConfig, error) {
	b, err := ioutil.ReadFile(os.Getenv("HOME") + "/.consul/config.json")
	if err != nil {
		return &ConsulConfig{
			Scheme: "http",
			Host:   "127.0.0.1:8500",
		}, nil
	}
	var config ConsulConfig
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func main() {
	fileConfig, err := readConfig()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	var (
		host     string
		user     string
		password string
		scheme   string
	)
	flag.StringVar(&host, "consul", fileConfig.Host, "Consul API end point")
	flag.StringVar(&scheme, "scheme", fileConfig.Scheme, "Consul API scheme")
	flag.StringVar(&user, "user", fileConfig.User, "Consul API user")
	flag.StringVar(&password, "password", fileConfig.Password, "Consul API password")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		printUsage()
	}

	consulConfig := &api.Config{
		Address: host,
		Scheme:  scheme,
	}
	if user != "" && password != "" {
		consulConfig.HttpAuth = &api.HttpBasicAuth{
			Username: user,
			Password: password,
		}
	}

	client, err := api.NewClient(consulConfig)
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
