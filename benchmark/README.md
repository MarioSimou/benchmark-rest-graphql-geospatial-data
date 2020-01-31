### Benchmark

In order to benchmark the service, [apache benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html) has been used.

**Percentage** | **Time(ms)** |
--- | --- |
50% | 1292 |
66% | 1449 |
75% | 1539 |
80% | 1606 |
90% | 1803 |
95% | 1966 |
98% | 2143 |
99% | 2305 |
100% | 2400 |


### Command

```
ab -n 1000 -c 10 localhost:3000/api/v1/cy/population
```