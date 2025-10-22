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

## Bonus - Compiled vs not compiled

In the above-mentioned benchmark, it creates a single executable binary, that is then started
and benchmark run against. What if you don't bother? Does it make a difference to the raw
performance?

### `bun run index.ts` vs. `./bun-run`

- `bun run index.ts` - 81811.8460 requests/sec
- `./bun-run` - 84457.0669 requests/sec

Ran it a bunch of times, back and forth, and noticed that numbers settled.
*Conclusion:* the executable binary runs ~3% faster.

### `go run run.go` vs. `./run`

- `go run run.go` - 67251.3725 requests/sec
- `./run` - 67702.8723 requests/sec

Ran it a bunch of times, back and forth, and noticed that numbers settled.
*Conclusion:* the executable binary runs no faster.

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

## Note-to-self; measuring memory use on macOS

I'm not sure if I this the best way to do it but here's how I measured peak memory usage:

1. Start the server. E.g. `just bun-start`
1. Use `ps aux | rg bun` to look for the executable `./bun-run` 's PID.
   Might look like this:

   ```text
   peterbengtsson   90845   0.0  0.0 483810288  21296 s005  S+   10:08AM   0:00.02 ./bun-run
   ```

1. Start `top` for that PID. E.g. `top -pid 90845`
1. Run the benchmark in a different terminal and watch the `top` output in the other.
   Might look something like this:

   ```text
   PID    COMMAND      %CPU TIME     #TH  #WQ  #POR MEM    PURG CMPR PGRP  PPID  STATE    BOOSTS    %CPU_ME %CPU_OTHRS UID  FAULT COW  MSGSE MSGRE SYSBSD  SYSMA CSW   PAGE IDLE POWE INSTRS     CYCLES     JETP
   90845  bun-run      0.0  00:01.19 6    0    28   9441K- 0B   0B   90823 90823 sleeping *0[1]     0.00000 0.00000    502  2305  142  4303+ 2140+ 248134+ 7459+ 5141+ 54   4    0.0  472550     754777     180
   ```

## Bonus - minimal disk I/O reading

Consider this change to the TypeScript code:

```diff
 Bun.serve({
   port: 3000,
-  fetch(_) {
-    return new Response("Hello world (bun)!");
+  async fetch(_) {
+    const jsonData = await Bun.file("data.json").text();
+    return new Response(jsonData, {headers: {'content-type': "application/json"}});
   },
 });
 ```

 I.e. on every request it reads the file `data.json` and returns an `application/json` content-type.

 And for the Go version, we do:

 ```go
 jsonData, err = os.ReadFile("data.json")
 ```

 on every request.

 When you benchmark these two versions the results become:

### Go (with disk I/O)

- low concurrency: 61559.9949 requests/sec
- medium concurrency: 62050.8071 requests/sec

### Bun (with disk I/O)

- low concurrency: 55621.6681 requests/sec
- medium concurrency: 58711.0603 requests/sec

Slight "win" for Go here. But conclusion is that they're both still fast enough.