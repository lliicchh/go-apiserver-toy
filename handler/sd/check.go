package sd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}

func DiskCheck(c *gin.Context) {
	u, _ := disk.Usage("/")
	userdMB := int(u.Used) / MB
	userdGB := int(u.Used) / GB
	totalMB := int(u.Total) / MB
	totalGB := int(u.Total) / GB
	usedPercent := int(u.UsedPercent)

	status := http.StatusOK
	text := "OK"

	if usedPercent >= 95 {
		status = http.StatusTooManyRequests
		text = "CRITICAL"
	} else if usedPercent >= 90 {
		status = http.StatusOK
		text = "WARNING"
	}

	message := fmt.Sprintf("%s - Free space: %dMB (%dGB) / %dMB (%dGB) | Used: %d%%", text, userdMB, userdGB, totalMB, totalGB)
	c.String(status, "\n"+message)
}

func CPUCheck(c *gin.Context) {
	var (
		cores   int
		text    string
		status  int
		message string
		l1      float64
		l5      float64
		l15     float64
		avgstat *load.AvgStat
	)
	cores, _ = cpu.Counts(false)

	avgstat, _ = load.Avg()
	l1 = avgstat.Load1
	l5 = avgstat.Load5
	l15 = avgstat.Load15

	status = http.StatusOK
	text = "OK"

	if l5 >= float64(cores-1) {
		status = http.StatusInternalServerError
		text = "CRITICAL"
	} else if l5 >= float64(cores-2) {
		status = http.StatusTooManyRequests
		text = "WARNING"
	}

	message = fmt.Sprintf("%s - Load average: %.2f, %.2f, %.2f | cores: %d",
		text, l1, l5, l15, cores)

	c.String(status, "\n"+message)
}
