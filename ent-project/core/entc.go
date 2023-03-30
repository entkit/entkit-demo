// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

//#go:build ignore
//# +build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/entkit/entkit"
	"log"
	"os"

	"path/filepath"
)

func main() {
	// The codegen is executed from internal/todo/gen.go.
	// So the path for the config file, ent schema, and the GQL schema
	// starts from internal/todo.
	gqlEx, err := entkit.NewEntgqlExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./ent.graphql"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	graphqlUri := os.Getenv("GRAPHQL_URI")
	if graphqlUri == "" {
		graphqlUri = "http://localhost/query"
	}
	entRefine, err := entkit.NewExtension(
		entkit.WithAppPath(filepath.Join("..", "refine-project")),
		entkit.WithGraphqlURL(graphqlUri),
		entkit.IgnoreUncommittedChanges(),
		//entkit.WithAuth(
		//	entkit.AuthWithKeycloak(
		//		entkit.NewKeycloak(
		//			"http://localhost:8080",
		//			"entkit",
		//			"admin",
		//			"admin",
		//			"entadmin",
		//			"entadmin",
		//			&gocloak.Client{
		//				ClientID: gocloak.StringP("xcontain-backend"),
		//				Secret:   gocloak.StringP("test-secret"),
		//			},
		//			[]*gocloak.Client{
		//				{
		//					ClientID: gocloak.StringP("xcontain-frontend"),
		//					RootURL:  gocloak.StringP("https://console.xcontain.com"),
		//					RedirectURIs: &[]string{
		//						"https://console.xcontain.com/*",
		//						"http://localhost:3000/*",
		//					},
		//					Attributes: &map[string]string{
		//						"post.logout.redirect.uris": "+",
		//					},
		//					WebOrigins: &[]string{
		//						"+",
		//					},
		//				},
		//			},
		//		),
		//	),
		//),
	)
	if err != nil {
		log.Fatalf("creating entkit extension: %v", err)
	}

	err = entc.Generate("./core/schema", &gen.Config{
		Target:  "./core/ent",
		Package: "github.com/entkit/entkit-demo/ent-project/core/ent",
		Header: `
			// Autogenerated by entkit
			//
			// Code generated by entc, DO NOT EDIT.
		`,
	}, entc.Extensions(
		gqlEx,
		entRefine,
	))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
