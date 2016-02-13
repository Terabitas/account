# Architecture 

# Development

## Running integration tests

Create docker machine:
```
docker-machine create --driver virtualbox nildev
```

Setup docker environment:
```
docker-machine start nildev
eval $(docker-machine env nildev)
```

Run required containers:
```
docker-compose -f docker-compose-dev.yml --x-networking up -d
```

Execute tests

Provisioning is happening inside `TestMain`
```
ND_IP_PORT=$(dm ip nildev):27017 go test -v -tags integration
```

## Testing locally

Build project container, it will produce docker image:
```
nildev build github.com/nildev.account
```

Create `local.env` file in root dir with content:
```
IRON_TOKEN=__YOUR_VAL__
IRON_PROJECT_ID=__YOUR_VAL__
ND_DATABASE_NAME=nildev
ND_MONGODB_URL=mongodb://mongodb.account.nildev:27017/nildev
ND_ENV=dev
ND_GITHUB_CLIENT_ID=__YOUR_VAL__
ND_GITHUB_SECRETE=__YOUR_VAL__
```

Run containers:
```
docker-compose -f docker-compose.yml --x-networking up -d
```

Execute HTTP request:
```
curl -X POST --data "@register.json" http://$(docker-machine ip nildev):8080/api/v1/Register -v
```

# Running service at [nildev.services](http://nildev.services)

## Project Details

### Release Notes

See the [releases tab](https://github.com/nildev/account/releases) for more information on each release.

### Contributing

See [CONTRIBUTING](CONTRIBUTING.md) for details on submitting patches and contacting developers via IRC and mailing lists.

### License

Project is released under the MIT license. See the [LICENSE](LICENSE) file for details.

Specific components of project use code derivative from software distributed under other licenses; in those cases the appropriate licenses are stipulated alongside the code.