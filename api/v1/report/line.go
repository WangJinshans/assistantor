package report

import (
	"assistantor/global"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

func ShowLine(){
	line := charts.NewLine()
	line.SetGlobalOptions(charts.TitleOpts{Title: "Line-显示 Label"})
	line.AddXAxis(nameItems).AddYAxis("商家A", randInt(), charts.LabelTextOpts{Show: true})
	reportPath, _ := global.GetReportPath()
	fileName := filepath.Join(reportPath, "line.html")
	f, err := os.Create(fileName)
	if err != nil {
		log.Error().Msgf("fail to create file, error is: %v", err)
		return
	}
	line.Render(f)
}
