### Benchmark

In order to benchmark the service, [apache benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html) has been used.

**Percentage** | **Time(ms)** |
--- | --- |
50% | 650 |
66% | 703 |
75% | 739 |
80% | 763 |
90% | 816 |
95% | 858 |
98% | 942 |
99% | 992 |
100% | 1228 |


### Command

```
ab -n 1000 -c 10 localhost:3000/api/v1/cy/population
```