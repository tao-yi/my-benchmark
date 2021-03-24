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
    Latency    13.77ms    3.91ms 104.57ms   97.61%
    Req/Sec     2.42k   373.35     3.12k    90.75%
  289345 requests in 10.01s, 53.81MB read
  Socket errors: connect 0, read 594, write 0, timeout 0
Requests/sec:  28898.26
Transfer/sec:      5.37MB
```

### express

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    51.85ms   13.47ms 221.38ms   92.05%
    Req/Sec   637.38    127.10     0.95k    79.67%
  76299 requests in 10.04s, 18.92MB read
  Socket errors: connect 0, read 461, write 0, timeout 0
Requests/sec:   7599.79
Transfer/sec:      1.88MB
```

### golang net/http

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.36ms    1.97ms  48.49ms   90.73%
    Req/Sec     6.23k     1.04k   22.24k    93.43%
  746547 requests in 10.11s, 101.81MB read
  Socket errors: connect 0, read 378, write 0, timeout 0
Requests/sec:  73876.46
Transfer/sec:     10.07MB
```

### golang fasthttp

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.96ms    0.95ms  35.98ms   88.27%
    Req/Sec     8.31k     0.95k   10.63k    73.42%
  992470 requests in 10.01s, 152.39MB read
  Socket errors: connect 0, read 375, write 0, timeout 0
Requests/sec:  99115.16
Transfer/sec:     15.22MB
```

### golang gin

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.62ms    2.22ms  61.62ms   88.43%
    Req/Sec     5.95k   664.02     7.29k    87.17%
  711631 requests in 10.02s, 100.44MB read
  Socket errors: connect 0, read 373, write 0, timeout 0
Requests/sec:  70996.22
Transfer/sec:     10.02MB
```

### golang fiber

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.97ms  830.82us  22.88ms   85.59%
    Req/Sec     8.28k     0.96k   10.02k    71.17%
  989090 requests in 10.01s, 125.45MB read
  Socket errors: connect 0, read 343, write 0, timeout 0
Requests/sec:  98819.73
Transfer/sec:     12.53MB
```

### go echo

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.67ms    2.56ms  78.78ms   88.40%
    Req/Sec     5.95k     1.47k   52.78k    98.08%
  711261 requests in 10.10s, 101.07MB read
  Socket errors: connect 0, read 382, write 0, timeout 0
Requests/sec:  70402.11
Transfer/sec:     10.00MB
```

### golang iris

```shell
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     5.39ms    2.21ms  51.34ms   89.23%
    Req/Sec     6.22k   690.54     7.83k    82.25%
  742873 requests in 10.01s, 109.10MB read
  Socket errors: connect 0, read 380, write 0, timeout 0
Requests/sec:  74193.07
Transfer/sec:     10.90MB
```

### Summary

| framework       | Requests/sec |
| --------------- | ------------ |
| express         | 7599.79      |
| fastify         | 28898.26     |
| golang net/http | 73876.46     |
| golang fasthttp | 99115.16     |
| golang gin      | 70996.22     |
| golang fiber    | 98819.73     |
| golang echo     | 70402.11     |
| golang iris     | 74193.07     |
