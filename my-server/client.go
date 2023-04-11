package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/entkit/entkit-demo"
	"github.com/entkit/entkit-demo/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

// This file will not be regenerated automatically.

func InitClient(ctx *cli.Context) (*ent.Client, error) {
	client, err := ent.Open(
		ctx.String("DatabaseDriver"),
		ctx.String("DatabaseSourceName"),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(
		context.Background(),
		//migrate.WithGlobalUniqueID(true),
	); err != nil {
		return nil, err
	}
	return client, nil
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return entproject.NewSchema(client)
}
