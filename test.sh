#!/bin/bash
docker network create porto
echo "Starting docker-compose"
docker-compose up -d


echo "Building integration tests"
pushd integtests
docker-compose up --build
#popd

#echo "Running integtests"
#integtests/integtests -addr=http://127.0.0.1:8081 -wait=50
result=$?
docker-compose down
popd
docker-compose down
docker rmi clientapi dmservice porto_psql integtests -f
docker network rm porto
[ "$result" -ne "0" ] && echo "Tests failed" || echo "Tests passed"
exit $result
