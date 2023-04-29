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
	"bytes"
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
	"strconv"
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
				&cli.BoolFlag{
					Name:    "UnionServer",
					Usage:   "Run API server, App server and other servers together",
					Aliases: []string{"union-server", "u"},
					Value:   false,
					EnvVars: []string{"UNION_SERVER"},
				},
				&cli.BoolFlag{
					Name:    "HTTPS",
					Aliases: []string{"https", "ssl"},
					Value:   false,
					EnvVars: []string{"HTTPS"},
				},
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
					Value:   "localhost",
					EnvVars: []string{"HOST"},
				},
				&cli.IntFlag{
					Name:    "Port",
					Aliases: []string{"port", "p"},
					Value:   80,
					EnvVars: []string{"PORT"},
				},
				&cli.StringFlag{
					Name:    "AppPath",
					Aliases: []string{"app-path"},
					Value:   "/",
					EnvVars: []string{"APP_PATH"},
				},
				&cli.BoolFlag{
					Name:    "AppServerEnabled",
					Aliases: []string{"app-server-enabled"},
					Value:   true,
					EnvVars: []string{"APP_SERVER_ENABLED"},
				},
				&cli.BoolFlag{
					Name:    "GraphqlServerEnabled",
					Aliases: []string{"graphql-server-enabled", "gql-server-enabled"},
					Value:   false,
					EnvVars: []string{"GRAPHQL_SERVER_ENABLED"},
				},
				&cli.StringFlag{
					Name:    "GraphqlURL",
					Aliases: []string{"graphql-url", "gql-url"},
					EnvVars: []string{"GRAPHQL_URL"},
				},
				&cli.StringFlag{
					Name:    "GraphqlURIPath",
					Aliases: []string{"graphql-uri-path", "gql-uri-path"},
					Value:   "/query",
					EnvVars: []string{"GRAPHQL_URI_PATH"},
				},
				&cli.StringFlag{
					Name:    "GraphqlPlaygroundURIPath",
					Aliases: []string{"graphql-playground-uri-path", "gql-playground-uri-path"},
					Value:   "/playground",
					EnvVars: []string{"GRAPHQL_PLAYGROUND_URI_PATH"},
				},
				&cli.BoolFlag{
					Name:    "GraphqlPlaygroundServerEnabled",
					Aliases: []string{"graphql-playground-server-enabled", "gql-playground-server-enabled", "gql-pg-server-enabled"},
					Value:   false,
					EnvVars: []string{"GRAPHQL_PLAYGROUND_SERVER_ENABLED"},
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
							log.Fatal("Unsupported log level", zap.String("provided", providedLevel))
						}
						log = log.WithOptions(zap.IncreaseLevel(level))

						defer log.Sync()

						isHttps := context.Bool("HTTPS")
						host := context.String("Host")
						port := ":" + strconv.Itoa(context.Int("Port"))
						appPath := strings.Trim(context.String("AppPath"), "/")
						if appPath == "" {
							appPath = "/"
						} else {
							appPath = "/" + appPath + "/"
						}
						gqlURL := context.String("GraphqlURL")
						gqlURIPath := context.String("GraphqlURIPath")
						gqlPlaygroundURIPath := context.String("GraphqlPlaygroundURIPath")
						proto := "http://"
						appUrl := proto + "localhost" + port + appPath
						gqlPlaygroundServerEnabled := context.Bool("GraphqlPlaygroundServerEnabled")
						appServerEnabled := context.Bool("AppServerEnabled")
						gqlServerEnabled := context.Bool("GraphqlServerEnabled")
						isUnionServer := context.Bool("UnionServer")
						if isUnionServer {
							log.Info("Union server enabled")
							appServerEnabled = true
							gqlServerEnabled = true
							gqlPlaygroundServerEnabled = true
						}

						if isHttps {
							proto = "https://"
						}

						if gqlURL == "" {
							gqlURL += proto + host + port + gqlURIPath
						}

						// Up MUX server
						mux := http.NewServeMux()
						var (
							keycloakHost                = context.String("KeycloakHost")
							keycloakRealm               = context.String("KeycloakRealm")
							keycloakFrontendClientID    = context.String("KeycloakFrontendClientID")
							keycloakBackendClientID     = context.String("KeycloakBackendClientID")
							keycloakBackendClientSecret = context.String("KeycloakBackendClientSecret")
						)

						if gqlServerEnabled {
							start := time.Now()
							client, err := InitClient(context)
							if err != nil {
								log.Fatal("Failed to start Ent client", zap.Error(err))
							}
							log.Info("Ent client initialized", zap.Duration("duration", time.Since(start)))

							srv := handler.NewDefaultServer(NewSchema(client))
							srv.Use(entgql.Transactioner{TxOpener: client})
							//srv.Use(&debug.Tracer{})
							mux.Handle(gqlURIPath, ent.DemoAuthMiddleware(
								srv,
								keycloakHost,
								keycloakRealm,
								keycloakBackendClientID,
								keycloakBackendClientSecret,
							))
							log.Info("Graphql server enabled", zap.String("url", gqlURL))
						}

						if gqlPlaygroundServerEnabled {
							gqlPlaygroundURL := proto + host + port + gqlPlaygroundURIPath
							mux.HandleFunc(gqlPlaygroundURIPath, playground.Handler("refine-project", gqlURIPath))
							log.Info("Graphql playground server enabled", zap.String("url", gqlPlaygroundURL))
						}

						if appServerEnabled {
							mux.HandleFunc(
								appPath,
								func(w http.ResponseWriter, r *http.Request) {
									if r.Method != "GET" {
										http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
										return
									}
									start := time.Now()
									log.Debug("Request starts", zap.String("path", r.URL.Path), zap.Time("time", start))

									path := filepath.Clean(r.URL.Path)
									path = strings.TrimPrefix(path, appPath)

									staticPaths := []string{"static", "images", "favicon.ico", "asset-manifest.json", "environment.json"}
									isStaticPath := false
									for _, v := range staticPaths {
										if path == v || strings.HasPrefix(path, v+"/") {
											isStaticPath = true
											continue
										}
									}

									if isStaticPath {
										w.Header().Set("Cache-Control", "public, max-age=31536000")
									} else {
										path = "index.html"
									}

									file, err := RefineProjectFS.Open(filepath.Join("refine-project", path))
									if err != nil {
										if os.IsNotExist(err) {
											log.Info("File not found", zap.String("path", path), zap.Error(err))
											http.NotFound(w, r)
											return
										}
										log.Info("File cannot be read", zap.String("path", path), zap.Error(err))
										http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
										return
									}

									contentType := mime.TypeByExtension(filepath.Ext(path))
									w.Header().Set("Content-Type", contentType)

									var n int64
									if path == "index.html" {
										b, _ := json.Marshal(entkit.Environment{
											Meta:       map[string]any{},
											GraphqlURL: gqlURL,
											AppPath:    appPath,
											Auth: &entkit.AuthEnvironment{
												Keycloak: &entkit.KeycloakEnvironment{
													URL:             keycloakHost,
													Realm:           keycloakRealm,
													ClientID:        keycloakFrontendClientID,
													BackendClientID: keycloakBackendClientID,
												},
											},
										})
										buf := new(bytes.Buffer)
										_, err := io.Copy(buf, file)
										if err != nil {
											panic(err)
										}
										newStr := buf.String()
										for _, sp := range staticPaths {
											newStr = strings.Replace(newStr, "/"+sp, appPath+sp, -1)
										}
										re := regexp.MustCompile("(<script data-name=\"environment\">window.environment=).*?(</script>)")
										newStr = re.ReplaceAllString(newStr, "$1"+string(b)+"$2")

										r := strings.NewReader(newStr)
										w.Header().Set("Content-Length", fmt.Sprintf("%d", r.Size()))
										n, _ = io.Copy(w, r)
									} else {
										stat, err := file.Stat()
										if err == nil && stat.Size() > 0 {
											w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
										}
										n, _ = io.Copy(w, file)
									}

									log.Debug("Serving file", zap.String("path", path), zap.Int64("bytes", n))
									log.Info("Request handled", zap.String("path", r.URL.Path), zap.Duration("duration", time.Since(start)))
								},
							)
							log.Info("App server enabled", zap.String("url", appUrl))
						}

						if err := http.ListenAndServe(port, mux); err != nil {
							log.Fatal("Server failed", zap.Error(err))
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
