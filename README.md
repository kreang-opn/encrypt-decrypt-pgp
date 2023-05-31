# `Encrypt / Decrypt PGP file`

## If you need to build it yourself, here are the requirements.
- golang 1.17
- install go module
```shell
go mod install
```
- build
```shell
## windows
make build-encrypt-windows
make build-decrypt-windows
```
```shell
## linux
make build-encrypt-linux
make build-decrypt-linux
```
```shell
## mac
make build-encrypt-mac-m1
make build-decrypt-mac-m1
```


## Command
```shell
./encrypt path-to-file path-to-public.key
```

```shell
./decrypt path-to-file.pgp path-to-private.key "passphrase"
```