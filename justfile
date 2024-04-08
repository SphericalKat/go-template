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

alias openapi := openapi-build

run:
    go run .

build:
    GOOS=linux GOARCH=amd64 go build -o main .