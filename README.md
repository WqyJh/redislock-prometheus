# redislock-prometheus

[![GoDoc](https://godoc.org/github.com/Wqyjh/redislock-prometheus?status.png)](http://godoc.org/github.com/Wqyjh/redislock-prometheus)


Prometheus metric collector for redislock.

## Usage

```go
import (
	rlprome "github.com/WqyJh/redislock-prometheus"
	"github.com/prometheus/client_golang/prometheus"
)

func init() {
    prometheus.MustRegister(rlprome.NewDefaultCollector())
}
```

The default metric name is `redislock_count`. You can also custom
the metric name by:

```go
    prometheus.MustRegister(rlprome.NewCollector("my_metric_name"))
````
