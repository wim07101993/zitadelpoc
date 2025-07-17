# POC for communicating with Zitadel

## Prerequisites

### Zitadel

You will need a Zitadel instance to handle the authentication procedures.

On the instance there need to be two clients: 

#### Web application client

- authorization code flow with PKCE
- redirect-uri (by default, this is http://localhost:8765, this requires
  development mode to be enabled)
- post-logout-redirect-uri (by default, this is http://localhost:8765, this 
  requires development mode to be enabled) 
- refresh tokens should be enabled

Ensure the configuration in [frontend/config.js](frontend/config.js) is correct.
The client-id and endpoints will need to be configured depending on your 
Zitadel instance.

#### Api client

- authorization using client credentials (client-id, client-secret)

### Go compiler

At least version 1.21.3 as specified in the [go.mod](go.mod) file.

## How to run

Ensure the location from which you run contains a `config.json` file or the 
required environment variables. An example of the `config.json` file can be 
found at [config.json](config.json). The names of the environment variables can
be found in [config.go](config.go). (client-id and secret in the config files
are made up)

Once everything is set up run `go run .` in the root of the repository.
