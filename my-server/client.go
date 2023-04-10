package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/entkit/entkit-demo/ent-project"
	"github.com/entkit/entkit-demo/ent-project/core/ent"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

// This file will not be regenerated automatically.

func InitClient(*cli.Context) (*ent.Client, error) {
	client, err := ent.Open(
		"sqlite3",
		"file:runtime/ent/demo.db?mode=rwc&cache=shared&_fk=1",
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
