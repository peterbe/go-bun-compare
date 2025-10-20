# go-bun-compare

Really basic Hello World web server benchmark comparison of Bun vs Go [fasthttp](github.com/valyala/fasthttp).

Inspired by https://medium.com/deno-the-complete-reference/bun-v-s-go-hello-world-performance-comparison-1f5418945112 from 2023.
It's Oct 2025 at the time of writing.

## Test Go

In one terminal:

```bash
just go-start
```

In another terminal:

```bash
just benchmark-low-concurrency
just benchmark
```

## Test Bun

In one terminal:

```bash
just bun-start
```

In another terminal:

```bash
just benchmark-low-concurrency
just benchmark
```

## My current results

On a M4 MacBook Pro
using `go version go1.25.3 darwin/arm64` and Bun `1.3.0`

### Go

- low concurrency: 67522.5031 requests/sec
- medium concurrency: 67998.6526 requests/sec

Memory peak: 7264K

### Bun

- low concurrency: 77118.8624 requests/sec
- medium concurrency: 76791.7434 requests/sec

Memory peak: 9921K

## File size

When you compile the `run.go` you get 7.1M file called `run`.
When you compile the `index.ts` you get a 58M file called `bun-run`.

## Conclusion(?)

Hello World server apps are almost pointless because they're so unrealistic in that it deals with no I/O, at all.
And raw speed isn't particularly useful.

What matters though is that the server isn't slowing you down in a chain of a bunch of
other services. The raw performance, without the I/O should be good enough.

Something this limited benchmark concludes that, if you like TypeScript, you don't need a statically
compiled language like Go to make a light web server. With Bun, it's fast enough. And in terms of
memory, it's *not* a hog.


## Bonus - Python

### Starlette and Flask

Starlette run by `uvicorn` and Flask run by `gunicorn` (with 8 workers) and `gevent`.

```bash
just starlette-run
```

and

```bash
just flask-run
```

### Current benchmarks

```bash
just benchmark
```

...yields...

- Starlette: 15896.2157 requests/sec
- Flask: 9564.5005 requests/sec