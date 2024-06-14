all:
    bash ./bin/allgen.sh

mocks:
    bash ./bin/mockgen.sh

lint:
    bash ./bin/lint.sh

providers:
    bash ./bin/providergen.sh

ent:
    bash ./bin/dbgen.sh

api:
    bash ./bin/apigen.sh

wire:
    bash ./bin/gogen.sh

openapi-validate:
    @echo "Validating OpenAPI spec..."
    bash ./spec/bin/validate.sh

openapi-build: openapi-validate
    @echo "Building OpenAPI spec..."
    bash ./spec/bin/build.sh

docker-build tag:
    @echo "Building docker container with tag {{tag}}..."
    docker build -t tag . -f docker/Dockerfile

alias openapi := openapi-build
alias docker := docker-build

run:
    go run .

build:
    GOOS=linux GOARCH=amd64 go build -o main .

migration-create name:
    cd migrations
    goose create {{name}}
    cd ..
