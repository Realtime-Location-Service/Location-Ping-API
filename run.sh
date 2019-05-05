#! /bin/bash

PROJECT="ping-api"

# run docker compose for consul
docker-compose up -d

# waiting for consul
until $(curl --output /dev/null --silent --fail http://127.0.0.1:8500/v1/kv); do
    echo 'waiting for consul'
    sleep 5
done

# set consul config key values from example
curl --request PUT --data-binary @config.example.yml http://127.0.0.1:8500/v1/kv/ping-api
 
# build binary
go build -o $PROJECT .

# set consul env
export PING_API_CONSUL_URL="127.0.0.1:8500"
export PING_API_CONSUL_PATH=$PROJECT

echo "ENV: PING_API_CONSUL_URL =" $PING_API_CONSUL_URL
echo "ENV: PING_API_CONSUL_PATH =" $PING_API_CONSUL_PATH

./$PROJECT serve
