package report

import (
	"assistantor/global"
	"github.com/go-echarts/go-echarts/charts"
	"os"
	"path/filepath"
)

var radarDataBJ = [][]float32{
	{55, 9, 56, 46, 18, 6, 1},
}

func ShowRadar() {
	var indicators = []charts.IndicatorOpts{
		{Name: "AQI", Max: 100},
		{Name: "PM2.5", Max: 100},
		{Name: "PM10", Max: 100},
		{Name: "CO", Max: 100},
		{Name: "NO2", Max: 100},
		{Name: "SO2", Max: 100},
	}
	radar := charts.NewRadar()
	radar.SetGlobalOptions(
		charts.TitleOpts{Title: "Radar-示例图"},
		charts.RadarComponentOpts{
			Indicator: indicators,
			SplitLine: charts.SplitLineOpts{Show: true},
			SplitArea: charts.SplitAreaOpts{Show: true},
		},
	)
	radar.Add("北京", radarDataBJ)

	reportPath, _ := global.GetReportPath()
	fileName := filepath.Join(reportPath, "radar.html")
	f, _ := os.Create(fileName)
	radar.Render(f)
	f.Close()
}