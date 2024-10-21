package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Env               string            `json:"env"`
	GrpcItemServer    GrpcItemServer    `json:"grpc_item_server"`
	GrpcStorageServer GrpcStorageServer `json:"grpc_storage_server"`
	PostgresDatabase  PostgresDatabase  `json:"postgres_database"`
}

type PostgresDatabase struct {
	Url string `json:"url"`
}

type GrpcItemServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type GrpcStorageServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

func MustRead(configPath string) *Config {
	file, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var config Config
	if err = json.Unmarshal(b, &config); err != nil {
		panic(err)
	}

	return &config
}
