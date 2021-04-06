### Hardware

- MacBook Pro (16-inch, 2019)
- Processor 2.6 GHz 6-Core Intel Core i7
- Memory 16 GB 2667 MHz DDR4

### fastify

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    28.58ms   11.87ms 240.36ms   93.36%
    Req/Sec     1.18k   275.35     1.55k    85.58%
  140920 requests in 10.02s, 26.21MB read
  Socket errors: connect 0, read 601, write 0, timeout 0
Requests/sec:  14060.38
Transfer/sec:      2.61MB
```

### express

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    90.01ms   23.88ms 358.50ms   89.14%
    Req/Sec   361.69     89.18   575.00     72.75%
  43347 requests in 10.06s, 10.75MB read
  Socket errors: connect 0, read 620, write 0, timeout 0
Requests/sec:   4310.95
Transfer/sec:      1.07MB
```

### golang net/http

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.47ms    5.97ms 146.80ms   86.83%
    Req/Sec     3.26k   472.59     5.38k    86.08%
  391575 requests in 10.08s, 53.40MB read
  Socket errors: connect 0, read 380, write 0, timeout 0
Requests/sec:  38842.09
Transfer/sec:      5.30MB
```

### golang fasthttp

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.62ms    3.20ms  60.61ms   89.94%
    Req/Sec     4.32k   654.32     6.62k    82.42%
  519104 requests in 10.07s, 79.70MB read
  Socket errors: connect 0, read 366, write 0, timeout 0
Requests/sec:  51557.88
Transfer/sec:      7.92MB
```

### golang gin

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.67ms   11.42ms 259.82ms   91.24%
    Req/Sec     2.93k   568.13     5.43k    85.75%
  351398 requests in 10.08s, 49.60MB read
  Socket errors: connect 0, read 380, write 0, timeout 0
Requests/sec:  34851.76
Transfer/sec:      4.92MB
```

### golang fiber

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     7.96ms    2.93ms  80.23ms   89.19%
    Req/Sec     4.11k   600.67     7.63k    86.50%
  494331 requests in 10.10s, 62.70MB read
  Socket errors: connect 0, read 382, write 0, timeout 0
Requests/sec:  48930.85
Transfer/sec:      6.21MB
```

### go echo

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    11.49ms    6.98ms 139.80ms   89.01%
    Req/Sec     2.99k   423.17     4.92k    89.83%
  358454 requests in 10.07s, 50.94MB read
  Socket errors: connect 0, read 371, write 0, timeout 0
Requests/sec:  35607.70
Transfer/sec:      5.06MB
```

### golang iris

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    10.48ms    6.01ms 110.68ms   87.45%
    Req/Sec     3.27k   441.86     4.28k    91.17%
  391777 requests in 10.04s, 57.54MB read
  Socket errors: connect 0, read 387, write 0, timeout 0
Requests/sec:  39040.32
Transfer/sec:      5.73MB
```

### express graphql

```shell
wrk -t12 -c400 -d10s -s ./gql.lua http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    79.10ms   17.53ms 318.59ms   92.47%
    Req/Sec   403.24     98.35   585.00     76.25%
  48344 requests in 10.06s, 19.00MB read
  Socket errors: connect 0, read 463, write 0, timeout 0
  Non-2xx or 3xx responses: 48344
Requests/sec:   4803.29
Transfer/sec:      1.89MB
```

### apollo server

```shell
wrk -t12 -c400 -d10s -s ./gql.lua http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   253.75ms   72.29ms 903.80ms   85.39%
    Req/Sec   130.09     54.69   290.00     66.02%
  15486 requests in 10.06s, 4.00MB read
  Socket errors: connect 0, read 440, write 0, timeout 0
Requests/sec:   1539.66
Transfer/sec:    407.47KB
```

### golang graphql

```shell
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     9.18ms    4.06ms  70.83ms   85.39%
    Req/Sec     3.65k   447.38     4.72k    87.42%
  436500 requests in 10.03s, 73.27MB read
  Socket errors: connect 0, read 375, write 0, timeout 0
  Non-2xx or 3xx responses: 436500
Requests/sec:  43526.56
Transfer/sec:      7.31MB
```

### Summary

| framework       | Requests/sec |
| --------------- | ------------ |
| express         | 4310.95      |
| express-graphql | 4803.29      |
| apollo-server   | 1539.66      |
| fastify         | 14060.38     |
| golang graphql  | 43526.56     |
| golang net/http | 38842.09     |
| golang fasthttp | 51557.88     |
| golang gin      | 34851.76     |
| golang fiber    | 48930.85     |
| golang echo     | 35607.70     |
| golang iris     | 39040.32     |
