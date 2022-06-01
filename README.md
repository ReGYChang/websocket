# websocket

- ✅ Get stream aggtrade data from [websocket API](https://github.com/binance/binance-spot-api-docs/blob/master/web-socket-streams.md)
- ✅ Gin HTTP Server
- ✅ Gin API to get last trade infomation

## Architecture
- client: subscribe stream data and save into redis server
- config: server config info
- global: global variables
- handler: Gin HTTP handler
- internal
  - model: data model
- storage
  - redis: redis func implementation
- config.json

## Startup

### Redis

Default

### Gin API

```shell
curl -L -X GET 'http://localhost:8080/load'
```