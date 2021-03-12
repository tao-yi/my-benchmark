### Hardware

- MacBook Pro (16-inch, 2019)
- Processor 2.6 GHz 6-Core Intel Core i7
- Memory 16 GB 2667 MHz DDR4

### fastify

```shell
$ make benchmark
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    12.24ms    3.60ms  84.58ms   95.29%
    Req/Sec     2.72k   474.17     4.12k    91.65%
  327730 requests in 10.10s, 60.95MB read
  Socket errors: connect 0, read 633, write 0, timeout 0
Requests/sec:  32434.06
Transfer/sec:      6.03MB
```

### express

```shell
$ make benchmark
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    41.92ms    9.54ms 178.08ms   94.45%
    Req/Sec   789.60    135.40     0.99k    83.75%
  94435 requests in 10.03s, 23.42MB read
  Socket errors: connect 0, read 472, write 0, timeout 0
Requests/sec:   9419.35
Transfer/sec:      2.34MB
```

### golang net/http

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.76ms  846.18us  22.47ms   89.14%
    Req/Sec     6.94k   667.05    12.19k    94.37%
  833974 requests in 10.10s, 113.73MB read
  Socket errors: connect 0, read 386, write 0, timeout 0
Requests/sec:  82538.04
Transfer/sec:     11.26MB
```

### golang fasthttp

```shell
$ make benchmark
wrk -t12 -c400 -d10s http://localhost:3999/
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     3.47ms    0.92ms  26.81ms   91.52%
    Req/Sec     9.54k     1.07k   15.56k    82.12%
  1146465 requests in 10.10s, 176.03MB read
  Socket errors: connect 0, read 376, write 0, timeout 0
Requests/sec: 113460.18
Transfer/sec:     17.42MB
```

### golang gin

```shell
$ make benchmark
Running 10s test @ http://localhost:3999/
  12 threads and 400 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency     4.79ms    1.09ms  32.35ms   90.36%
    Req/Sec     6.90k     1.11k   26.61k    98.00%
  826110 requests in 10.10s, 116.60MB read
  Socket errors: connect 0, read 375, write 0, timeout 0
Requests/sec:  81756.17
Transfer/sec:     11.54MB
```
