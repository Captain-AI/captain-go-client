# `captain-go-client`

Go client for `captain.ai` REST API in JSON format. 

### Testing Live Requests

If you want to test live requests, create a file with environment variables.

```
CAPTAIN_ACCOUNT_ID=
CAPTAIN_INTEGRATION_KEY=
CAPTAIN_DEVELOPER_KEY=
```

Inject the envrionment variables for testing. 

```sh
$ env $(cat .env | xargs) go test ./...
```

If you don't want to test live requests pass the `-short` testing option. 

```sh
$ go test -short ./...
```
