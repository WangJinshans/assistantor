package report

import (
	"assistantor/global"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

var kvData = map[string]interface{}{
	"教授":  15,
	"副教授": 20,
	"讲师":  30,
	"助教":  25,
	"其他":  8,
}

func ShowPie() {
	pie := charts.NewPie()
	pie.SetGlobalOptions(charts.TitleOpts{Title: "Pie-示例图"})
	pie.Add("pie", kvData, charts.LabelTextOpts{Show: true, Formatter: "{b}: {c}%"})
	reportPath, _ := global.GetReportPath()
	fileName := filepath.Join(reportPath, "pie.html")
	f, err := os.Create(fileName)
	if err != nil {
		log.Error().Msgf("fail to create file, error is: %v", err)
		return
	}
	pie.Render(f)
}
