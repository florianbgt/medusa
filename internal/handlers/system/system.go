package system

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/mem"
)

func getWiFiSSID() string {
	out, err := exec.Command("iwgetid", "--raw").Output()
	if err != nil {
		panic(err)
	}

	name := string(out)

	return name
}

func getCPULoad() float64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}

	return v.UsedPercent
}

func getCPUTemp() float64 {
	temps := []float64{}

	index := 0
	for {
		out, err := exec.Command("cat", "/sys/class/thermal/thermal_zone"+strconv.Itoa(index)+"/temp").Output()
		if err != nil {
			break
		}

		cleanOut := strings.Split(string(out), "\n")[0]

		temp, err := strconv.ParseFloat(cleanOut, 32)
		if err != nil {
			panic(err)
		}

		temps = append(temps, temp/1000)

		index++
	}

	temp := 0.0
	for _, t := range temps {
		if t > temp {
			temp = t
		}
	}

	return temp
}

func SystemInfo(c *gin.Context) {
	c.JSON(200, gin.H{
		"ssid": getWiFiSSID(),
	})
}

func SystemMetrics(c *gin.Context) {
	c.JSON(200, gin.H{
		"cpu_load": getCPULoad(),
		"cpu_temp": getCPUTemp(),
	})
}
