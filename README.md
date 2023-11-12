# go-url-shortener

## Deploy

1. clone repo
2. fill `.env` file with appropriate environment variables
3. run `docker-compose up`

Infrastructure consists of two containers:
- Golang application container
- Database container

By default save all data (URLs) in PostgreSQL database running in container.

## Webserver

Webserver is running on 8080 TCP port (HTTP).
Two endpoints available:

- `/shorten-url` supports `POST` method  
  json body: `{"url": string}`

  Generate short URL for passed original URL, save it and return generated.

- `/{short-url}` - supports `GET` method  
  Dynamic endpoint path.  
  Will answer 404 NotFound or redirect to original URL.
