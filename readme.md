# Link Shortener

Run locally: `env $(cat ./.envs | xargs) go run ./main.go`

Build static binary and cook that into a Docker image:

```
GOOS=linux go build -o shortener ./main.go
docker build -t shortener:master .
docker run -d --env-file ./.envs -p 8080:8080 shortener:master
```

Once running you will have:

 * `POST /create` with a `url` form field.
 * `GET /{code}` where the `code` is returned from the create request.
