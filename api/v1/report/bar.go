package report

import (
	"assistantor/global"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/rs/zerolog/log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

var nameItems = []string{"衬衫", "牛仔裤", "运动裤", "袜子", "冲锋衣", "羊毛衫"}
var foodItems = []string{"面包", "牛奶", "奶茶", "棒棒糖", "加多宝", "可口可乐"}

var seed = rand.NewSource(time.Now().UnixNano())

func randInt() []int {
	cnt := len(nameItems)
	r := make([]int, 0)
	for i := 0; i < cnt; i++ {
		r = append(r, int(seed.Int63())%50)
	}
	return r
}

func ShowBar() {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.TitleOpts{Title: "Bar-示例图"}, charts.ToolboxOpts{Show: true})
	bar.AddXAxis(nameItems).
		AddYAxis("商家A", randInt()).
		AddYAxis("商家B", randInt())
	reportPath, _ := global.GetReportPath()
	fileName := filepath.Join(reportPath, "bar.html")
	f, err := os.Create(fileName)
	if err != nil {
		log.Error().Msgf("fail to create file, error is: %v", err)
		return
	}
	bar.Render(f)
}
