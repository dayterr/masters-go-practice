package client

import (
	"strconv"
	"time"
)

type Client struct {
	Counter      uint64
	PollInterval time.Duration
}

func NewClient() Client {
	c := Client{
		Counter:      0,
		PollInterval: Interval * time.Second,
	}

	return c
}

type Metrics struct {
	CurServerLoadAverage uint64
	CurRAM               float64
	CurRAMUsed           float64
	CurDiskSize          uint64
	CurDiskUsed          uint64
	CurNetBandwidth      uint64
	CurNetLoad           uint64
}

func NewMetrics(metricsArr []string) Metrics {
	m := Metrics{}

	metrInt, err := strconv.ParseInt(metricsArr[0], 10, 64)
	if err != nil {
		m.CurServerLoadAverage = 1
	}
	m.CurServerLoadAverage = uint64(metrInt)

	metrFl, err := strconv.ParseFloat(metricsArr[1], 64)
	if err != nil {
		m.CurRAM = 1
	}
	m.CurRAM = metrFl

	metrFl, err = strconv.ParseFloat(metricsArr[2], 64)
	if err != nil {
		m.CurRAMUsed = 1
	}
	m.CurRAMUsed = metrFl

	metrInt, err = strconv.ParseInt(metricsArr[3], 10, 64)
	if err != nil {
		m.CurDiskSize = 1
	}
	m.CurDiskSize = uint64(metr)

	metr, err = strconv.ParseInt(metricsArr[4], 10, 64)
	if err != nil {
		m.CurDiskUsed = 1
	}
	m.CurDiskUsed = uint64(metr)

	metr, err = strconv.ParseInt(metricsArr[5], 10, 64)
	if err != nil {
		m.CurNetBandwidth = 1
	}
	m.CurNetBandwidth = uint64(metr)

	metr, err = strconv.ParseInt(metricsArr[6], 10, 64)
	if err != nil {
		m.CurNetLoad = 1
	}
	m.CurNetLoad = uint64(metr)

	return m
}
