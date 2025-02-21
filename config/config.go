package config

import (
	"flag"
	"fmt"
)

type Config struct {
	RedisAddr string
	PGConnStr string
}

var Cfg *Config

func init() {
	host := flag.String("host", "localhost", "Redis host")
	port := flag.String("port", "6379", "Redis port")
	pgConnStr := flag.String("pg", "postgres://app:admin@127.0.0.1:5432/postgres?sslmode=disable", "PostgreSQL connection string")

	Cfg = &Config{
		RedisAddr: fmt.Sprintf("%s:%s", *host, *port),
		PGConnStr: *pgConnStr,
	}
}
