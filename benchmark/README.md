### Benchmark

In order to benchmark the service, [apache benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html) has been used.

**nRequests**: 1000

**concurrency**: 10

**Percentage** | **Time(ms)** |
--- | --- |
50% | 1638 |
66% | 1754 |
75% | 1847 |
80% | 1905 |
90% | 2077 |
95% | 2272 |
98% | 2478 |
99% | 2564 |
100% | 2779 |


### Command

```
ab -n 1000 -c 10 -p ./benchmark/body.json -m POST -T application/json localhost:3000/graphql
```