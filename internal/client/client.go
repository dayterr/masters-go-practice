package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

/*func (client Client) Run(ctx context.Context) {
	ticker := time.NewTicker(Interval * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				client.ReadMetrics()
			case <-ctx.Done():
				return
			}
		}
	}()

}*/

func (client Client) Run() {
	for {
		client.ReadMetrics()
	}

}

func (client Client) ReadMetrics() {
	c := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return
	}
	req.Header.Set("Host", Host)

	resp, err := c.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK || resp.Header.Get("Content-Type") != ContentTypeHeader {
		client.Counter++
		if client.Counter >= ErrorAmountThreshold {
			fmt.Println("Unable to fetch server statistic")
		}
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return
		}

		metricsArr := strings.Split(string(body), ",")
		m := NewMetrics(metricsArr)

		if m.CurServerLoadAverage >= ServerLoadAverageThreshold {
			fmt.Printf("Load Average is too high: %d\n", m.CurServerLoadAverage)
		}

		if m.CurRAMUsed >= m.CurRAM*RAMConsumptionThreshold {
			perc := (m.CurRAMUsed / m.CurRAM) * 100
			fmt.Printf("Memory usage too high: %d%%\n", uint64(perc))
		}

		if float64(m.CurDiskUsed) >= float64(m.CurDiskSize)*DiskConsumptionThreshold {
			left := m.CurDiskSize - m.CurDiskUsed
			left /= BytesInMBytes
			fmt.Printf("Free disk space is too low: %d Mb left\n", left)
		}

		if float64(m.CurNetLoad) >= float64(m.CurNetBandwidth)*NetBandwidthThreshold {
			left := m.CurNetBandwidth - m.CurNetLoad
			left /= BitsInMBits
			fmt.Printf("Network bandwidth usage high: %d Mbit/s available\n", left)
		}

	}
}
