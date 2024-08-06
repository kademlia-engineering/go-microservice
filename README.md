# Description

This service exposes 6 api's that perform cryptographic operations. Specifically AES Encryption and Eliptic Curve Cryptography.

# How to run

From the working directory run the following command,

`go run .`

# API Documentation

- Get Version (healthcheck API)

```
curl --location 'http://127.0.0.1:8080/api/v1/version'
```


- Encrypt Data

```
curl --location 'http://127.0.0.1:8080/api/v1/encrypt' \
--header 'cipher: My-Encryption-Key' \
--header 'Content-Type: application/json' \
--data '{
    "data": "Hello World"
}'
```


- Decrypt Data

```
curl --location 'http://127.0.0.1:8080/api/v1/decrypt' \
--header 'cipher: My-Encryption-Key' \
--header 'Content-Type: application/json' \
--data '{
    "data": "Hello World"
}'
```


- Create Ed25519 Keypair

```
curl --location 'http://127.0.0.1:8080/api/v1/ed25519'
```


- Ed25519 Signature

```
curl --location 'http://127.0.0.1:8080/api/v1/ed25519/sign' \
--header 'private-key: My-Private-Key' \
--header 'Content-Type: application/json' \
--data '{
    "data": "Hello World"
}'
```


- Ed25519 Signature Verification

```
curl --location 'http://127.0.0.1:8080/api/v1/ed25519/verify' \
--header 'public-key: My-Public-Key' \
--header 'Content-Type: application/json' \
--data '{
    "data": "Hello World"
}'
```
