# POC for communicating with Zitadel

## Prerequisites

### Zitadel

You will need a Zitadel instance to handle the authentication procedures.

On the instance there need to be two clients: 

#### 1. Web application 
A client with authorization code flow with PKCE enabled. 

Ensure that the redirect uri configured correctly in zitadel. By default, 
this is http://localhost:8765.

Ensure the configuration in [frontend/config.js](frontend/config.js) is correct.
The client-id and endpoints will need to be configured depending on your 
Zitadel instance.

#### 2. Backend application jwt authentication enabled (for token introspection)

### Go compiler

At least version 1.21.3 as specified in the [go.mod](go.mod) file.

## How to run

Ensure the location from which you run contains a `config.json` file or the 
required environment variables. An example of the `config.json` file can be 
found at [config.json](config.json). The names of the environment can
be found in [config.go](config.go). (client-id and secret in the config files
are made up)

Once everything is set up run `go run .` in the root of the repository.
