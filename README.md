# URL-Shortener
URL-Shortener project (Written in Golang)

- Start the `Postgres` server and `Redis` for storing the data.
- Create a `urlservice` database in the `Postgres` server
- Run the commands as follows `docker-compose up -d` to start the Kafka(KRaft).
- Run `go run cmd/main.go`.

- Use the [Postman](https://www.postman.com/) and send a `POST` request to `localhost:3000/shorten` with the following data as an example (just for demo):
```
 {
    "url": "https://google.com"
 }

```

- You'll receive a `shortcode url` to your original URL(in this case `"https://google.com"`).
