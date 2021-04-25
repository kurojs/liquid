# Liquid

This is an assigment for Liquid interview 
The repo is contains two sessions:
- Blockchain get balance at timestamp
- A Server and Client using authentication based on JWT and ecdsa as a signing method  

# Blockchain get balance at timestamp 

[Blockchain GetBalanceAt](./blockchain/get_balance.go)
Implemented using binary search
Time complexity: O(log(N))
Space complexity: O(1)
Unit-test included: [Unit-test](./blockchain/get_balance_test.go)

# Auth Server and Client

- Server handles simple login request form API, then check with existing data in MySQL. If user exists, generate access 
token which will expired in 15 minute. The token is generated by using ECDSA as signing method with a secret-public key pare 
, and only this public key can use to get data from the access token.

- By default, there will be three user in DB, you can use these for testing: `kuro`, `liquid`, `skywalker`. The username and
password are the same. There are *POSTMAN* collection in `./docs` you can import to Postman and test by yourself. 
  
- Client handles simple verify request form API, get the access token and parses it with the given public key.

# How to

- Required: `go`, `docker-compose`

- Start docker compose first: `docker-compose -up -d`

- Start auth server and client: `go run main.go` 

- Can set port for server or client by using env `export SERVER_PORT=8090` and `export SERVER_PORT=9090`