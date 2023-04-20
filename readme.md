# EntKit demo project

# Setup

Initially run `docker compose up -d` to run required services

Then run `go generate` on project root directory

Finally, run built in generated server to serve demo application

```shell
go run ./my-server/*.go serve -u --app-path=/myapp/qweqwe/pwdpwd refine-project
```