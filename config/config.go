package config

import (
	"flag"
	"os"
	"strconv"
)

const DefaultPort int = 5000

var (
	hostname string
	port     int
	dbroot   string
)

func Hostname() string {
	return hostname
}

func Port() int {
	return port
}

func DbRoot() string {
	return dbroot
}

func Parse() {
	envport, _ := strconv.Atoi(os.Getenv("HYPER_PORT"))

	if envport == 0 {
		envport = DefaultPort
	}

	flag.StringVar(
		&hostname,
		"hostname",
		os.Getenv("HYPER_HOSTNAME"),
		"Server Hostname")

	flag.IntVar(
		&port,
		"port",
		envport,
		"Server Port")

	flag.StringVar(
		&dbroot,
		"dbroot",
		os.Getenv("HYPER_DB_ROOT"),
		"Database root dir")

	flag.Parse()
}
