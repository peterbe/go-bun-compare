# https://github.com/casey/just
# https://just.systems/


go-run:
    go run run.go

go-build:
    go build run.go

go-start: go-build
    ./run

bun-run:
    bun run index.ts

bun-build:
    bun build index.ts --outfile=bun-run --minify --compile

bun-start: bun-build
    ./bun-run

benchmark:
    oha -n 100000 -c 100 http://localhost:3000/

benchmark-low-concurrency:
    oha -n 100000 -c 10 http://localhost:3000/

starlette-run:
    cd python && uv sync
    cd python && source .venv/bin/activate && uvicorn starlette-run:app --port 3000 --log-level error

flask-run:
    cd python && uv sync
    cd python && source .venv/bin/activate && PORT=3000 gunicorn -w 8 'flask-run:app' --preload -k gevent