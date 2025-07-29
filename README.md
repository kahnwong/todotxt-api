# todotxt-api

## Usage

```bash
docker build -t todotxt-api .
docker run \
  -p 3000:3000 \
  -v $(pwd)/todotxt:/opt/todotxt \
  -e TODO_PATH=/opt/todotxt/todo.txt \
  -e LISTEN_ADDR=:3000 \
  todotxt-api
```
