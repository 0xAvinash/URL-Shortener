# URL-Shortener
URL-Shortener project (Written in Golang)

- Run the commands as follows `docker-compose up -d` to start the Kafka(KRaft), Postgres, Redis and the app.

- Use the [Postman](https://www.postman.com/) and send a `POST` request to `localhost:3000/shorten` with the following data as an example (just for demo):
```
 {
    "url": "https://google.com"
 }

```

- You'll receive a `shortcode url` to your original URL(in this case `"https://google.com"`).
#### output
```
{
    "short_url": "http://localhost:3000/KDJPxEsaBbC"
}

```
- Open the above `http://localhost:3000/KDJPxEsaBbC` in a web browser of your choice to see the demo in action.
