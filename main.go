package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

type HostStats struct {
	CPUPercent  float64 `json:"cpu_percent"`
	MemoryUsed  uint64  `json:"memory_used"`
	MemoryTotal uint64  `json:"memory_total"`
	DiskUsed    uint64  `json:"disk_used"`
	DiskTotal   uint64  `json:"disk_total"`
}

func getHostStats() (*HostStats, error) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	diskStat, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	return &HostStats{
		CPUPercent:  cpuPercent[0],
		MemoryUsed:  vmStat.Used,
		MemoryTotal: vmStat.Total,
		DiskUsed:    diskStat.Used,
		DiskTotal:   diskStat.Total,
	}, nil
}

func main() {
	r := gin.Default()

	r.GET("/stats", func(c *gin.Context) {
		stats, err := getHostStats()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, stats)
	})

	r.Run(":8080") // listen on port 8080
}
