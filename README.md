## Build :

```
./docker_build.sh
```

## Run as testing environment with docker-compose :

```
./run.sh
```

## Production :

```
cd charts/consul && helm install --namespace devops -n consul -f values.yaml stable/consul
cd ../../charts/getenv && helm install --namespace devops -n getenv .
```

## Test

```
curl -XPOST -H "Content-Type: application/json" -d '{"app_name":"stg-quoinex-11", "url":"http://rc11-quoinex-api.quoinex.com"}'  http://localhost/services
```

```
curl -v -XGET http://localhost/services/stg-quoinex-11
```