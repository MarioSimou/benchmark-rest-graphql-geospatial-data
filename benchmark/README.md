### Benchmark

In order to benchmark the service, [apache benchmark](https://httpd.apache.org/docs/2.4/programs/ab.html) has been used.

**Percentage** | **Time(ms)** |
--- | --- |
50% | 2277 |
66% | 2409 |
75% | 2480 |
80% | 2551 |
90% | 2797 |
95% | 3076 |
98% |  3343 |
99% |  3503 |
100% | 3951 |

### Command

```
ab -n 1000 -c 10 -p ./benchmark/body.json -m POST -T application/json localhost:3000/graphql
```