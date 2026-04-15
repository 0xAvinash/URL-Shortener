# URL-Shortener
URL-Shortener project (Written in Golang)

- Run the commands as follows `docker-compose up -d` to start the Kafka(KRaft)
- Run `go run cmd/main.go`

- Use the [Postman]([text](https://www.postman.com/)) and send a `POST` request to `localhost:3000/shorten` with the following data as an example (just for demo)
```
 {
    "url": "https://google.com"
 }

```

- You'll receive `ShotCode` to your URL