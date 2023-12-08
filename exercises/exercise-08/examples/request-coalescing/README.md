# Benchmark: Request Coalescing in Go

## Setup
First start a new postgres container using

    docker run --rm \
        --name postgres \
        -dp 5432:5432 \
        -e POSTGRES_PASSWORD=postgres \
        -e POSTGRES_DB=postgres \
        postgres

Then start the test server using

    go run main.go

This will setup the database and make the server listen on
`http://localhost:8080`. Now generate the test data using

    curl http://localhost:8080/generate

## Run the benchmark
The server provides two routes. 

* `GET` `/benchmark/without-request-coalescing` \
Returns a random message from the database without request coalescing.

* `GET` `/benchmark/with-request-coalescing` \
Returns a random message using request coalescing.

You may want to benchmark those routes using a HTTP benchmarking tool of your
choice. The results below are outputs of
[ab](https://httpd.apache.org/docs/2.4/programs/ab.html) by running

    ab -n 100000 -c 1000 http://localhost:8080/benchmark/without-request-coalescing

and

    ab -n 100000 -c 1000 http://localhost:8080/benchmark/with-request-coalescing

## Benchmark results

### Without request coalescing
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /benchmark/without-request-coalescing
Document Length:        146 bytes

Concurrency Level:      1000
Time taken for tests:   15.742 seconds
Complete requests:      50000
Failed requests:        6424
   (Connect: 0, Receive: 0, Length: 6424, Exceptions: 0)
Non-2xx responses:      1591
Total transferred:      12924697 bytes
HTML transferred:       7062881 bytes
Requests per second:    3176.16 [#/sec] (mean)
Time per request:       314.846 [ms] (mean)
Time per request:       0.315 [ms] (mean, across all concurrent requests)
Transfer rate:          801.77 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   1.4      0      15
Processing:     0  205 1145.2     55   15719
Waiting:        0  205 1145.1     55   15719
Total:          0  206 1145.4     55   15728

Percentage of the requests served within a certain time (ms)
  50%     55
  66%     84
  75%     93
  80%    100
  90%    122
  95%   1063
  98%   1338
  99%   3497
 100%  15728 (longest request)
```

### With request coalescing
```
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /benchmark/with-request-coalescing
Document Length:        146 bytes

Concurrency Level:      1000
Time taken for tests:   1.530 seconds
Complete requests:      50000
Failed requests:        4958
   (Connect: 0, Receive: 0, Length: 4958, Exceptions: 0)
Total transferred:      13195042 bytes
HTML transferred:       7295042 bytes
Requests per second:    32669.38 [#/sec] (mean)
Time per request:       30.610 [ms] (mean)
Time per request:       0.031 [ms] (mean, across all concurrent requests)
Transfer rate:          8419.41 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    5  29.3      4    1039
Processing:     0   26  15.1     23     106
Waiting:        0   24  15.0     22     106
Total:          0   30  33.2     27    1093

Percentage of the requests served within a certain time (ms)
  50%     27
  66%     35
  75%     39
  80%     42
  90%     50
  95%     57
  98%     66
  99%     74
 100%   1093 (longest request)
```
