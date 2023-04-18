// Code generated by EntKit. DO NOT EDIT.
// ---------------------------------------------------------
//
// Copyright (C) 2023 EntKit. All Rights Reserved.
//
// This code is part of the EntKit library and is generated
// automatically to ensure optimal functionality and maintainability.
// Any changes made directly to this file may be overwritten
// by future code generation, leading to unexpected behavior.
//
// Please refer to the EntKit documentation for instructions on
// how to modify the library, extend its functionality or contribute
// to the project: https://entkit.com
// ---------------------------------------------------------

package main

import (
	"encoding/json"
	"entgo.io/contrib/entgql"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/entkit/entkit-demo/ent"
	"github.com/entkit/entkit/v2"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"mime"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "entkit"
	app.Version = "0.1.0"
	app.Description = "This is EntKit CLI"
	app.Authors = []*cli.Author{
		{
			Name:  "Aaron Yordanyan",
			Email: "aaron.yor@gmail.com",
		},
	}

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "LogLevel",
			Aliases: []string{"log-level"},
			Value:   "debug",
			EnvVars: []string{"LOG_LEVEL"},
		},
		&cli.BoolFlag{
			Name:    "DevMode",
			Aliases: []string{"dev-mode", "dev"},
			Value:   true,
			EnvVars: []string{"DEV_MODE"},
		},
	}

	app.Commands = []*cli.Command{
		{
			Name: "docker",
			Subcommands: []*cli.Command{
				{
					Name: "build",
					Action: func(context *cli.Context) error {
						cmd := exec.Command(
							"docker",
							"build",
							"--file",
							"./my-server/Dockerfile",
							"--tag",
							"demo-my-server",
							".",
						)
						cmd.Env = append(os.Environ(), "DOCKER_BUILDKIT=1")
						cmd.Stderr = os.Stderr
						cmd.Stdout = os.Stdout
						cmd.Stdin = os.Stdin
						err := cmd.Run()

						if err != nil {
							return err
						}

						return nil
					},
				},
				{
					Name: "push",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "registry",
							Required: true,
							Aliases:  []string{"reg"},
						},
						&cli.StringSliceFlag{
							Name:  "tags",
							Value: cli.NewStringSlice("latest"),
						},
					},
					Action: func(context *cli.Context) error {
						tags := context.StringSlice("tags")
						repoName := context.String("registry") + "/demo-my-server"

						for _, tag := range tags {
							remoteImageName := repoName + ":" + tag
							cmd := exec.Command(
								"docker",
								"image",
								"tag",
								"demo-my-server:latest",
								remoteImageName,
							)
							cmd.Stderr = os.Stderr
							cmd.Stdout = os.Stdout
							cmd.Stdin = os.Stdin
							err := cmd.Run()
							if err != nil {
								return err
							}

							cmd = exec.Command(
								"docker",
								"image",
								"push",
								remoteImageName,
							)
							cmd.Stderr = os.Stderr
							cmd.Stdout = os.Stdout
							cmd.Stdin = os.Stdin
							err = cmd.Run()
							if err != nil {
								return err
							}
						}
						return nil
					},
				},
			},
		},
		{
			Name:    "serve",
			Aliases: []string{"s"},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "DatabaseDriver",
					Aliases: []string{"db-driver"},
					Value:   "sqlite3",
					EnvVars: []string{"DB_DRIVER"},
				},
				&cli.StringFlag{
					Name:    "DatabaseSourceName",
					Aliases: []string{"db-dsn"},
					Value:   "file:runtime/ent/demo.db?mode=rwc&cache=shared&_fk=1",
					EnvVars: []string{"DB_DSN"},
				},
				&cli.StringFlag{
					Name:    "Host",
					Aliases: []string{"host"},
					Value:   ":80",
					EnvVars: []string{"HOST"},
				},
				&cli.StringFlag{
					Name:    "GraphqlURL",
					Aliases: []string{"graphql-url", "gql-url"},
					Value:   "http://localhost/query",
					EnvVars: []string{"GRAPHQL_URL"},
				},
				&cli.StringFlag{
					Name:    "KeycloakHost",
					Aliases: []string{"keycloak-host", "kc-host"},
					Value:   "http://localhost:8080",
					EnvVars: []string{"KEYCLOAK_HOST"},
				},
				&cli.StringFlag{
					Name:    "KeycloakRealm",
					Value:   "entkit-demo-3",
					Aliases: []string{"keycloak-realm", "kc-realm"},
					EnvVars: []string{"KEYCLOAK_REALM"},
				},
				&cli.StringFlag{
					Name:    "KeycloakFrontendClientID",
					Value:   "frontend",
					Aliases: []string{"keycloak-frontend-client-id", "kc-frontend-client-id"},
					EnvVars: []string{"KEYCLOAK_FRONTEND_CLIENT_ID"},
				},
				&cli.StringFlag{
					Name:    "KeycloakBackendClientID",
					Value:   "backend",
					Aliases: []string{"keycloak-backend-client-id", "kc-backend-client-id"},
					EnvVars: []string{"KEYCLOAK_BACKEND_CLIENT_ID"},
				},
				&cli.StringFlag{
					Name:    "KeycloakBackendClientSecret",
					Value:   "test-secret",
					Aliases: []string{"keycloak-backend-client-secret", "kc-backend-client-secret"},
					EnvVars: []string{"KEYCLOAK_BACKEND_CLIENT_SECRET"},
				},
			},
			Description: "Serve EntKit APP",
			Subcommands: []*cli.Command{
				{
					Name: "refine-project",
					Action: func(context *cli.Context) error {
						var log *zap.Logger
						if context.Bool("DevMode") {
							log, _ = zap.NewDevelopment()
						} else {
							log, _ = zap.NewProduction()
						}

						providedLevel := context.String("LogLevel")

						level, err := zapcore.ParseLevel(providedLevel)
						if err != nil {
							log.Fatal("unsupported log level", zap.String("provided", providedLevel))
						}
						log = log.WithOptions(zap.IncreaseLevel(level))

						defer log.Sync()

						log.Info("starting server...")

						start := time.Now()
						client, err := InitClient(context)
						if err != nil {
							log.Fatal("failed to start ent client", zap.String("reason", err.Error()))
						}
						log.Info("ent client initialized", zap.Duration("duration", time.Since(start)))

						// Up MUX server
						mux := http.NewServeMux()

						srv := handler.NewDefaultServer(NewSchema(client))
						srv.Use(entgql.Transactioner{TxOpener: client})
						//srv.Use(&debug.Tracer{})
						mux.HandleFunc("/playground", playground.Handler("Example", "/query"))
						var (
							keycloakHost                = context.String("KeycloakHost")
							keycloakRealm               = context.String("KeycloakRealm")
							keycloakFrontendClientID    = context.String("KeycloakFrontendClientID")
							keycloakBackendClientID     = context.String("KeycloakBackendClientID")
							keycloakBackendClientSecret = context.String("KeycloakBackendClientSecret")
						)

						mux.Handle("/query", ent.EntkitAuthMiddleware(
							srv,
							keycloakHost,
							keycloakRealm,
							keycloakBackendClientID,
							keycloakBackendClientSecret,
						))

						mux.HandleFunc("/environment.json", func(writer http.ResponseWriter, request *http.Request) {
							b, _ := json.Marshal(entkit.Environment{
								Meta:       map[string]any{},
								GraphqlURL: context.String("GraphqlURL"),
								Auth: &entkit.AuthEnvironment{
									Keycloak: &entkit.KeycloakEnvironment{
										URL:             keycloakHost,
										Realm:           keycloakRealm,
										ClientID:        keycloakFrontendClientID,
										BackendClientID: keycloakBackendClientID,
									},
								},
							})
							writer.Write(b)
						})

						if os.Getenv("SKIP_EMBED_SERVER") != "true" {
							mux.HandleFunc(
								"/",
								func(w http.ResponseWriter, r *http.Request) {
									if r.Method != "GET" {
										http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
										return
									}
									start := time.Now()
									log.Debug("request starts", zap.String("path", r.URL.Path), zap.Time("time", start))

									path := filepath.Clean(r.URL.Path)
									path = strings.TrimPrefix(path, "/")

									rePattern := `^asset-manifest\.json|^environment\.json|^favicon\.ico|^images/.*|^static/.*`
									isStaticPath, _ := regexp.MatchString(rePattern, path)

									if isStaticPath {
										w.Header().Set("Cache-Control", "public, max-age=31536000")
									} else {
										path = "index.html"
									}

									file, err := RefineProjectFS.Open(filepath.Join("refine-project", path))
									if err != nil {
										if os.IsNotExist(err) {
											log.Warn("file not found", zap.String("path", path), zap.String("reason", err.Error()))
											http.NotFound(w, r)
											return
										}
										log.Warn("file cannot be read", zap.String("path", path), zap.String("reason", err.Error()))
										http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
										return
									}

									contentType := mime.TypeByExtension(filepath.Ext(path))
									w.Header().Set("Content-Type", contentType)

									stat, err := file.Stat()
									if err == nil && stat.Size() > 0 {
										w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
									}

									n, _ := io.Copy(w, file)
									log.Debug("serving file", zap.String("path", path), zap.Int64("bytes", n))
									log.Info("request handled", zap.String("path", r.URL.Path), zap.Duration("duration", time.Since(start)))
								},
							)
						}
						host := context.String("Host")
						log.Info("server started", zap.String("address", host))
						if err := http.ListenAndServe(host, mux); err != nil {
							log.Fatal("server failed", zap.String("reason", err.Error()))
						}
						return nil
					},
				},
				{
					Name: "other-refine-project",
					Action: func(context *cli.Context) error {
						var log *zap.Logger
						if context.Bool("DevMode") {
							log, _ = zap.NewDevelopment()
						} else {
							log, _ = zap.NewProduction()
						}

						providedLevel := context.String("LogLevel")

						level, err := zapcore.ParseLevel(providedLevel)
						if err != nil {
							log.Fatal("unsupported log level", zap.String("provided", providedLevel))
						}
						log = log.WithOptions(zap.IncreaseLevel(level))

						defer log.Sync()

						log.Info("starting server...")

						start := time.Now()
						client, err := InitClient(context)
						if err != nil {
							log.Fatal("failed to start ent client", zap.String("reason", err.Error()))
						}
						log.Info("ent client initialized", zap.Duration("duration", time.Since(start)))

						// Up MUX server
						mux := http.NewServeMux()

						srv := handler.NewDefaultServer(NewSchema(client))
						srv.Use(entgql.Transactioner{TxOpener: client})
						//srv.Use(&debug.Tracer{})
						mux.HandleFunc("/playground", playground.Handler("Example", "/query"))
						var (
							keycloakHost                = context.String("KeycloakHost")
							keycloakRealm               = context.String("KeycloakRealm")
							keycloakFrontendClientID    = context.String("KeycloakFrontendClientID")
							keycloakBackendClientID     = context.String("KeycloakBackendClientID")
							keycloakBackendClientSecret = context.String("KeycloakBackendClientSecret")
						)

						mux.Handle("/query", ent.EntkitAuthMiddleware(
							srv,
							keycloakHost,
							keycloakRealm,
							keycloakBackendClientID,
							keycloakBackendClientSecret,
						))

						mux.HandleFunc("/environment.json", func(writer http.ResponseWriter, request *http.Request) {
							b, _ := json.Marshal(entkit.Environment{
								Meta:       map[string]any{},
								GraphqlURL: context.String("GraphqlURL"),
								Auth: &entkit.AuthEnvironment{
									Keycloak: &entkit.KeycloakEnvironment{
										URL:             keycloakHost,
										Realm:           keycloakRealm,
										ClientID:        keycloakFrontendClientID,
										BackendClientID: keycloakBackendClientID,
									},
								},
							})
							writer.Write(b)
						})

						if os.Getenv("SKIP_EMBED_SERVER") != "true" {
							mux.HandleFunc(
								"/",
								func(w http.ResponseWriter, r *http.Request) {
									if r.Method != "GET" {
										http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
										return
									}
									start := time.Now()
									log.Debug("request starts", zap.String("path", r.URL.Path), zap.Time("time", start))

									path := filepath.Clean(r.URL.Path)
									path = strings.TrimPrefix(path, "/")

									rePattern := `^asset-manifest\.json|^environment\.json|^favicon\.ico|^images/.*|^static/.*`
									isStaticPath, _ := regexp.MatchString(rePattern, path)

									if isStaticPath {
										w.Header().Set("Cache-Control", "public, max-age=31536000")
									} else {
										path = "index.html"
									}

									file, err := OtherRefineProjectFS.Open(filepath.Join("other-refine-project", path))
									if err != nil {
										if os.IsNotExist(err) {
											log.Warn("file not found", zap.String("path", path), zap.String("reason", err.Error()))
											http.NotFound(w, r)
											return
										}
										log.Warn("file cannot be read", zap.String("path", path), zap.String("reason", err.Error()))
										http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
										return
									}

									contentType := mime.TypeByExtension(filepath.Ext(path))
									w.Header().Set("Content-Type", contentType)

									stat, err := file.Stat()
									if err == nil && stat.Size() > 0 {
										w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
									}

									n, _ := io.Copy(w, file)
									log.Debug("serving file", zap.String("path", path), zap.Int64("bytes", n))
									log.Info("request handled", zap.String("path", r.URL.Path), zap.Duration("duration", time.Since(start)))
								},
							)
						}
						host := context.String("Host")
						log.Info("server started", zap.String("address", host))
						if err := http.ListenAndServe(host, mux); err != nil {
							log.Fatal("server failed", zap.String("reason", err.Error()))
						}
						return nil
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err.Error())
	}
}
