#! /usr/bin/env sh

echo "Formatting go files"
go fmt src/*.go
go fmt src/interpreter/*.go
go fmt src/stdlib/*.go

echo ""
echo "Running unit tests."
cd src/interpreter
echo "  * Intepreter"
go test
cd ../stdlib
echo "  * Standard Library"
go test
cd ../codegen
echo "  * Code Geenerators"
go test
cd ..
echo "  * Unicorn"
go test
cd ..

echo ""
echo "Installing dependencies"
echo "  * Yaml.v2"
go get -u gopkg.in/yaml.v2

echo ""
echo "Building Unicorn"
go build src/unicorn.go
if [ $? -eq 0 ]; then
    echo "[SUCCESS] - Built Unicorn"
else
    echo "[FAILURE] - Could not build Unicorn"
fi
