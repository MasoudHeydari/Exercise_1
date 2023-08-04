package store

import (
	"context"
	"log"
	"net/url"

	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/config"
	"github.com/MasoudHeydari/Exercise_1/Task2_imagy/ent"
	_ "github.com/lib/pq"
)

type Database struct {
	Client *ent.Client
}

func New(config config.Config) *Database {
	u := url.URL{
		Scheme:   config.DbSchema,
		User:     url.UserPassword(config.DbUsername, config.DbPassword),
		Host:     config.DbAddress,
		Path:     config.DbName,
		RawQuery: config.DbSslMode,
	}

	client, err := ent.Open(u.Scheme, u.String())
	if err != nil {
		log.Fatal("Failed to open a connection to database ", "error", err)
	}

	if err = client.Schema.Create(context.Background()); err != nil {
		log.Fatal("Failed to create database schema ", "error", err)
	}

	return &Database{
		Client: client,
	}
}
