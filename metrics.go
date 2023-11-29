package rlprome

import (
	"github.com/bsm/redislock"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	DefaultPrefix = "redislock_count"
)

const (
	opObtainSuccess = "obtain_success"
	opObtainFailed  = "obtain_failed"
	opObtainError   = "obtain_error"
	opObtainCancel  = "obtain_cancel"

	opReleaseSuccess = "release_success"
	opReleaseFailed  = "release_failed"
	opReleaseError   = "release_error"
	opReleaseCancel  = "release_cancel"

	opRefreshSuccess = "refresh_success"
	opRefreshFailed  = "refresh_failed"
	opRefreshError   = "refresh_error"
	opRefreshCancel  = "refresh_cancel"

	opBackoff      = "backoff"
	opWatchdog     = "watchdog"
	opWatchdogDone = "watchdog_done"
	opWatchdogTick = "watchdog_tick"
)

type Collector struct {
	desc *prometheus.Desc
}

func NewDefaultCollector() *Collector {
	return NewCollector(DefaultPrefix)
}

func NewCollector(name string) *Collector {
	return &Collector{
		desc: prometheus.NewDesc(name, "RedisLock stats", []string{"op"}, nil),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.desc
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	stats := redislock.GetStats()

	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Obtain.Success), opObtainSuccess)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Obtain.Failed), opObtainFailed)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Obtain.Error), opObtainError)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Obtain.Cancel), opObtainCancel)

	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Release.Success), opReleaseSuccess)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Release.Failed), opReleaseFailed)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Release.Error), opReleaseError)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Release.Cancel), opReleaseCancel)

	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Refresh.Success), opRefreshSuccess)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Refresh.Failed), opRefreshFailed)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Refresh.Error), opRefreshError)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Refresh.Cancel), opRefreshCancel)

	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Backoff), opBackoff)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.Watchdog), opWatchdog)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.WatchdogDone), opWatchdogDone)
	ch <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, float64(stats.WatchdogTick), opWatchdogTick)
}
