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
curl -XPOST -H "Content-Type: application/json" -d '{"app_name":"stg-quoinex-11", "url":"http://rc11-quoinex-api.quoinex.com"}'  https://getenv.quoine.me/services
```

```
curl -XGET https://getenv.quoine.me/services/stg-quoinex-11
```

```
curl -XDELETE https://getenv.quoine.me/services/stg-quoinex-11
```

