#!/bin/bash

echo "Starting docker-compose"
docker-compose up --build &


echo "Building integration tests"
pushd integtests
go build
popd

echo "Running integtests"
integtests/integtests -addr=http://127.0.0.1:8081 -wait=40
result=$?

docker-compose down
docker rmi clientapi dmservice porto_psql -f

[ "$result" -ne "0" ] && echo "Tests failed" || echo "Tests passed"
exit $result
