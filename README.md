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
    Latency    10.30ms    3.11ms  81.40ms   97.07%
    Req/Sec     3.24k   734.73    17.92k    96.84%
  388119 requests in 10.10s, 72.18MB read
  Socket errors: connect 0, read 656, write 0, timeout 0
Requests/sec:  38411.12
Transfer/sec:      7.14MB
```

### express

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    39.19ms    9.95ms 208.23ms   94.16%
    Req/Sec   848.53    205.89     6.20k    92.76%
  101542 requests in 10.10s, 25.18MB read
  Socket errors: connect 0, read 447, write 0, timeout 0
Requests/sec:  10049.47
Transfer/sec:      2.49MB
```

### golang net/http

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.06ms    1.22ms  30.09ms   87.64%
    Req/Sec     8.14k   728.39     9.92k    83.42%
  972681 requests in 10.01s, 132.65MB read
  Socket errors: connect 0, read 366, write 0, timeout 0
Requests/sec:  97158.53
Transfer/sec:     13.25MB
```

### golang fasthttp

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.57ms    1.22ms  34.65ms   92.08%
    Req/Sec    12.99k     1.64k   20.75k    81.07%
  1551334 requests in 10.01s, 238.19MB read
  Socket errors: connect 0, read 362, write 0, timeout 394
Requests/sec: 154902.58
Transfer/sec:     23.78MB
```

### golang gin

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.90ms    1.39ms  28.09ms   87.04%
    Req/Sec     8.51k     1.18k   23.93k    96.26%
  1019767 requests in 10.10s, 143.93MB read
  Socket errors: connect 0, read 379, write 0, timeout 0
Requests/sec: 100934.48
Transfer/sec:     14.25MB
```

### golang fiber

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     2.52ms    0.86ms  28.46ms   92.28%
    Req/Sec    13.08k     1.31k   16.98k    78.99%
  1573840 requests in 10.10s, 199.62MB read
  Socket errors: connect 0, read 356, write 0, timeout 0
Requests/sec: 155763.58
Transfer/sec:     19.76MB
```

### golang echo

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.08ms    1.99ms  56.14ms   87.10%
    Req/Sec     8.24k     1.96k   69.99k    97.50%
  985456 requests in 10.10s, 140.03MB read
  Socket errors: connect 0, read 372, write 0, timeout 0
Requests/sec:  97546.08
Transfer/sec:     13.86MB
```

### golang iris

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.07ms    1.83ms  41.31ms   86.59%
    Req/Sec     8.25k   675.43    10.11k    86.67%
  985659 requests in 10.02s, 144.76MB read
  Socket errors: connect 0, read 368, write 0, timeout 0
Requests/sec:  98418.16
Transfer/sec:     14.45MB
```

### express graphql

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    33.17ms    8.11ms 182.47ms   96.05%
    Req/Sec     1.00k   144.24     1.23k    88.83%
  120074 requests in 10.02s, 47.06MB read
  Socket errors: connect 0, read 507, write 0, timeout 0
  Non-2xx or 3xx responses: 120074
Requests/sec:  11982.98
Transfer/sec:      4.70MB
```

### apollo server

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    74.45ms   15.05ms 331.53ms   91.44%
    Req/Sec   442.68     75.60   610.00     77.33%
  53002 requests in 10.05s, 13.34MB read
  Socket errors: connect 0, read 765, write 0, timeout 0
  Non-2xx or 3xx responses: 53002
Requests/sec:   5274.37
Transfer/sec:      1.33MB
```

### golang graphql

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.02ms    1.38ms  23.59ms   85.07%
    Req/Sec     8.24k     1.54k   42.11k    98.75%
  986263 requests in 10.10s, 165.54MB read
  Socket errors: connect 0, read 371, write 0, timeout 0
  Non-2xx or 3xx responses: 986263
Requests/sec:  97632.36
Transfer/sec:     16.39MB
```

### python flask

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    39.79ms   50.70ms 397.68ms   98.20%
    Req/Sec   148.35     96.25   747.00     76.73%
  6593 requests in 10.10s, 1.08MB read
  Socket errors: connect 0, read 16391, write 0, timeout 0
Requests/sec:    653.05
Transfer/sec:    109.05KB
```

### Summary

| framework       | Requests/sec |
| --------------- | ------------ |
| python flask    | 653.05       |
| apollo-server   | 5274.37      |
| express         | 10049.47     |
| express-graphql | 11982.98     |
| fastify         | 38411.12     |
| golang net/http | 97158.53     |
| golang echo     | 97546.08     |
| golang graphql  | 97632.36     |
| golang iris     | 98418.16     |
| golang gin      | 100934.48    |
| golang fasthttp | 154902.58    |
| golang fiber    | 155763.58    |
