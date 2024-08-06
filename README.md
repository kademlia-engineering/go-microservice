# Description

This service exposes 6 api's that perform cryptographic operations. Specifically AES Encryption and Eliptic Curve Cryptography.

# How to run

From the working directory run the following command,

`go run .`

# API Documentation

### Get Version (healthcheck API)

**Request**
```
curl --location 'http://127.0.0.1:8080/api/v1/version'
```

**Response**
```
{
    "version": "0.1.0"
}
```


### Encrypt Data

**Request**
```
curl --location 'http://127.0.0.1:8080/api/v1/encrypt' \
--header 'cipher: My-Encryption-Key' \
--header 'Content-Type: application/json' \
--data '{
    "data": "Hello World"
}'
```

**Response**
```
{
    "payload": "5bd6b3b984d20501fa7d7bc35559a7ab.8a6d0cbf164c14dbedfbb60d82621730.319e2a93eac14394e91aa8891b4a2839.f635ed12696cbaa974e4c0"
}
```


### Decrypt Data

**Request**
```
curl --location 'http://127.0.0.1:8080/api/v1/decrypt' \
--header 'cipher: My-Encryption-Key' \
--header 'Content-Type: application/json' \
--data '{
    "data": "5bd6b3b984d20501fa7d7bc35559a7ab.8a6d0cbf164c14dbedfbb60d82621730.319e2a93eac14394e91aa8891b4a2839.f635ed12696cbaa974e4c0"
}'
```

**Response**
```
{
    "payload": "Hello World"
}
```


### Create Ed25519 Keypair

**Request**
```
curl --location 'http://127.0.0.1:8080/api/v1/ed25519'
```

**Response**
```
{
    "curve": "ed23319",
    "public_key": "gCwsXBov1BL7SvRgFjtzSgiJbMUSSYVsT8M6Uy4MsRY",
    "private_key": "8aJQw2h6waWt7rDxwZ1u3JaBqdWcfh2HxaRRzhj43iwg"
}
```


### Ed25519 Signature

**Request**
```
curl --location 'http://127.0.0.1:8080/api/v1/ed25519/sign' \
--header 'private-key: 8aJQw2h6waWt7rDxwZ1u3JaBqdWcfh2HxaRRzhj43iwg' \
--header 'public-key: gCwsXBov1BL7SvRgFjtzSgiJbMUSSYVsT8M6Uy4MsRY' \
--header 'Content-Type: application/json' \
--data '{
    "data": "Hello World"
}'
```

**Response**
```
{
    "payload": "ec5b5d490fe02f25caa27ee07cd7402cb87bcf1888a9914276c890126896fdc2297d45e7f513ca3b6451ef1573d15969c6ac5b6889c4930d558a02a70b12af02"
}
```


### Ed25519 Signature Verification

**Request**
```
curl --location 'http://127.0.0.1:8080/api/v1/ed25519/verify' \
--header 'public-key: gCwsXBov1BL7SvRgFjtzSgiJbMUSSYVsT8M6Uy4MsRY' \
--header 'signature: ec5b5d490fe02f25caa27ee07cd7402cb87bcf1888a9914276c890126896fdc2297d45e7f513ca3b6451ef1573d15969c6ac5b6889c4930d558a02a70b12af02' \
--header 'Content-Type: application/json' \
--data '{
    "data": "Hello World"
}'
```

**Response**
```
{
    "payload": "true"
}
```
