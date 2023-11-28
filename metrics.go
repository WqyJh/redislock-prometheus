package rlprome

import (
	"github.com/bsm/redislock"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	DefaultPrefix = "redislock"
)

type statusDesc struct {
	Success *prometheus.Desc
	Failed  *prometheus.Desc
	Error   *prometheus.Desc
	Cancel  *prometheus.Desc
}

type Collector struct {
	obtain  *statusDesc
	release *statusDesc
	refresh *statusDesc

	backoff      *prometheus.Desc
	watchdog     *prometheus.Desc
	watchdogDone *prometheus.Desc
	watchdogTick *prometheus.Desc
}

func NewCollector(prefix string) *Collector {
	if prefix == "" {
		prefix = DefaultPrefix
	}

	return &Collector{
		obtain: &statusDesc{
			Success: prometheus.NewDesc(prefix+"_obtain_success", "The number of successful obtain", nil, nil),
			Failed:  prometheus.NewDesc(prefix+"_obtain_failed", "The number of failed obtain", nil, nil),
			Error:   prometheus.NewDesc(prefix+"_obtain_error", "The number of error obtain", nil, nil),
			Cancel:  prometheus.NewDesc(prefix+"_obtain_cancel", "The number of cancel obtain", nil, nil),
		},
		release: &statusDesc{
			Success: prometheus.NewDesc(prefix+"_release_success", "The number of successful release", nil, nil),
			Failed:  prometheus.NewDesc(prefix+"_release_failed", "The number of failed release", nil, nil),
			Error:   prometheus.NewDesc(prefix+"_release_error", "The number of error release", nil, nil),
			Cancel:  prometheus.NewDesc(prefix+"_release_cancel", "The number of cancel release", nil, nil),
		},
		refresh: &statusDesc{
			Success: prometheus.NewDesc(prefix+"_refresh_success", "The number of successful refresh", nil, nil),
			Failed:  prometheus.NewDesc(prefix+"_refresh_failed", "The number of failed refresh", nil, nil),
			Error:   prometheus.NewDesc(prefix+"_refresh_error", "The number of error refresh", nil, nil),
			Cancel:  prometheus.NewDesc(prefix+"_refresh_cancel", "The number of cancel refresh", nil, nil),
		},
		backoff:      prometheus.NewDesc(prefix+"_backoff", "The number of backoff", nil, nil),
		watchdog:     prometheus.NewDesc(prefix+"_watchdog", "The number of watchdog", nil, nil),
		watchdogDone: prometheus.NewDesc(prefix+"_watchdog_done", "The number of watchdog done", nil, nil),
		watchdogTick: prometheus.NewDesc(prefix+"_watchdog_tick", "The number of watchdog tick", nil, nil),
	}
}

func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.obtain.Success
	ch <- c.obtain.Failed
	ch <- c.obtain.Error
	ch <- c.obtain.Cancel

	ch <- c.release.Success
	ch <- c.release.Failed
	ch <- c.release.Error
	ch <- c.release.Cancel

	ch <- c.refresh.Success
	ch <- c.refresh.Failed
	ch <- c.refresh.Error
	ch <- c.refresh.Cancel

	ch <- c.backoff
	ch <- c.watchdog
	ch <- c.watchdogDone
	ch <- c.watchdogTick
}

func (c *Collector) Collect(ch chan<- prometheus.Metric) {
	stats := redislock.GetStats()

	ch <- prometheus.MustNewConstMetric(c.obtain.Success, prometheus.CounterValue, float64(stats.Obtain.Success))
	ch <- prometheus.MustNewConstMetric(c.obtain.Failed, prometheus.CounterValue, float64(stats.Obtain.Failed))
	ch <- prometheus.MustNewConstMetric(c.obtain.Error, prometheus.CounterValue, float64(stats.Obtain.Error))
	ch <- prometheus.MustNewConstMetric(c.obtain.Cancel, prometheus.CounterValue, float64(stats.Obtain.Cancel))

	ch <- prometheus.MustNewConstMetric(c.release.Success, prometheus.CounterValue, float64(stats.Release.Success))
	ch <- prometheus.MustNewConstMetric(c.release.Failed, prometheus.CounterValue, float64(stats.Release.Failed))
	ch <- prometheus.MustNewConstMetric(c.release.Error, prometheus.CounterValue, float64(stats.Release.Error))
	ch <- prometheus.MustNewConstMetric(c.release.Cancel, prometheus.CounterValue, float64(stats.Release.Cancel))

	ch <- prometheus.MustNewConstMetric(c.refresh.Success, prometheus.CounterValue, float64(stats.Refresh.Success))
	ch <- prometheus.MustNewConstMetric(c.refresh.Failed, prometheus.CounterValue, float64(stats.Refresh.Failed))
	ch <- prometheus.MustNewConstMetric(c.refresh.Error, prometheus.CounterValue, float64(stats.Refresh.Error))
	ch <- prometheus.MustNewConstMetric(c.refresh.Cancel, prometheus.CounterValue, float64(stats.Refresh.Cancel))

	ch <- prometheus.MustNewConstMetric(c.backoff, prometheus.CounterValue, float64(stats.Backoff))
	ch <- prometheus.MustNewConstMetric(c.watchdog, prometheus.CounterValue, float64(stats.Watchdog))
	ch <- prometheus.MustNewConstMetric(c.watchdogDone, prometheus.CounterValue, float64(stats.WatchdogDone))
	ch <- prometheus.MustNewConstMetric(c.watchdogTick, prometheus.CounterValue, float64(stats.WatchdogTick))
}
