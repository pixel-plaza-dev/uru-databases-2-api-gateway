@echo off
echo Generating Swagger docs...
swag init --parseDependency --parseInternal
echo Generating Swagger docs... Done.

echo Compiling .go files...
go build -o ./bin/
echo Compiling .go files... Done.
