package company_report

import (
	"assistantor/global"
	"assistantor/model/economy"
	"assistantor/utils"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/rs/zerolog/log"
	"io"
	"os"
	"path/filepath"
)

// 报告期相关
func getCashFlowReportDateList(dataList []economy.CashFlowInfo) (dateList []opts.BarData) {
	for _, item := range dataList {
		date := item.ReportDate
		dateList = append(dateList, opts.BarData{
			Value: date,
		})
	}
	return
}

func getProfitReportDateList(dataList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range dataList {
		date := item.ReportDate
		dateList = append(dateList, opts.BarData{
			Value: date,
		})
	}
	return
}

func getAssetReportDateList(dataList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range dataList {
		date := item.ReportDate
		dateList = append(dateList, opts.BarData{
			Value: date,
		})
	}
	return
}

// 现金流量相关
func getSalesList(infoList []economy.CashFlowInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.Sales, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getOperateList(infoList []economy.CashFlowInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.Operate, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getInvestList(infoList []economy.CashFlowInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.Sales, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getFinanceList(infoList []economy.CashFlowInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.Sales, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

//func getRequestDataList(dataList interface{}, attributes string) (dataSet []interface{}) {
//
//	switch infoList := dataList.(type) {
//	case []economy.ProfitInfo:
//		for _, info := range infoList {
//			t := reflect.TypeOf(info)
//			v := reflect.ValueOf(info)
//			for i := 0; i < t.NumField(); i++ {
//				fieldName := t.Field(i).Name
//				position := t.Field(i).Tag.Get("position")
//				if position == "x" {
//					data := v.Field(i).String()
//					dataSet = append(dataSet, data)
//					continue
//				}
//				if fieldName == attributes {
//					data, _ := utils.ShowNumberFriendly(v.Field(i).Float())
//					dataSet = append(dataSet, data)
//				}
//			}
//		}
//	}
//	return
//}
//
//func getBarDataList(dataList interface{}, attributes string) (dateList []opts.BarData) {
//	dataSet := getRequestDataList(dataList, attributes)
//	for _, data := range dataSet {
//		dateList = append(dateList, opts.BarData{
//			Value: data,
//			Label: &opts.Label{
//				Show:      true,
//				Formatter: "{a}: {c}",
//			},
//		})
//	}
//	return
//}

// 利润表相关
func getTotalOperateIncomeList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.TotalOperateIncome, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getTotalOperateCostList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.TotalOperateCost, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getOperateCostList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.OperateCost, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getOperateTaxList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		//data, _ := utils.ShowNumberFriendly(item.OperateTax)
		data := utils.ShowNumberFriendlyNumberWithUnit(item.OperateTax, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getSaleExpenseList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.SaleExpense, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getManageExpenseList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.ManageExpense, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getResearchExpenseList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.ResearchExpense, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getFinanceExpenseList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.FinanceExpense, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getInvestJointIncomeList(infoList []economy.ProfitInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.InvestJointIncome, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

// 资产负债相关
func getMonetaryFundList(infoList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.MonetaryFunds, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getNoteAccountsReceiveList(infoList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.NoteAccountsReceive, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getPrepaymentList(infoList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.Prepayment, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getInventoryList(infoList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.Inventory, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getFixedAssetList(infoList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.FixedAsset, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getCipList(infoList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.Cip, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getStaffSalaryPayableList(infoList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.StaffSalaryPayable, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

func getUnAssignProfitList(infoList []economy.AssetInfo) (dateList []opts.BarData) {
	for _, item := range infoList {
		data := utils.ShowNumberFriendlyNumberWithUnit(item.UnAssignProfit, 1)
		dateList = append(dateList, opts.BarData{
			Value: data,
			Label: &opts.Label{
				Show:      true,
				Formatter: "{c}",
			},
		})
	}
	return
}

// 报告
func GenerateProfitReport(stockId string) {

	infoList := GetProfitReport(stockId)
	reportDateItems := getProfitReportDateList(infoList)

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Bar 示例图"}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} 亿"},
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
	)

	bar.SetXAxis(reportDateItems).
		AddSeries("营业收入", getTotalOperateIncomeList(infoList)).
		AddSeries("营业总成本", getTotalOperateCostList(infoList)).
		AddSeries("营业成本", getOperateCostList(infoList)).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)

	bar1 := charts.NewBar()
	bar1.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Bar 示例图"}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} 亿"},
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
	)

	bar1.SetXAxis(reportDateItems).
		AddSeries("税金及附加", getOperateTaxList(infoList)).
		AddSeries("销售费用", getSaleExpenseList(infoList)).
		AddSeries("管理费用", getManageExpenseList(infoList)).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)

	bar2 := charts.NewBar()
	bar2.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Bar 示例图"}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} 亿"},
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
	)

	bar2.SetXAxis(reportDateItems).
		AddSeries("研发费用", getResearchExpenseList(infoList)).
		AddSeries("财务费用", getFinanceExpenseList(infoList)).
		AddSeries("投资收益", getInvestJointIncomeList(infoList)).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)

	page := components.NewPage()
	page.SetLayout(components.PageFlexLayout)
	page.AddCharts(
		bar,
		bar1,
		bar2,
	)

	reportPath, _ := global.GetReportPath()
	fileName := filepath.Join(reportPath, "profit.html")
	f, err := os.Create(fileName)
	if err != nil {
		log.Error().Msgf("fail to create file, error is: %v", err)
		return
	}
	page.Render(io.MultiWriter(f))
}

func GenerateAssetsLiabilityReport(stockId string) {
	infoList := GetAssetsLiabilityReport(stockId)
	reportDateItems := getAssetReportDateList(infoList)

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Bar 示例图"}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} 亿"},
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
	)

	bar.SetXAxis(reportDateItems).
		AddSeries("货币资金", getMonetaryFundList(infoList)).
		AddSeries("应收账款", getNoteAccountsReceiveList(infoList)).
		AddSeries("预收款", getPrepaymentList(infoList)).
		AddSeries("存货", getInventoryList(infoList)).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)

	bar1 := charts.NewBar()
	bar1.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "Bar 示例图"}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} 亿"},
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
	)
	bar1.SetXAxis(reportDateItems).
		AddSeries("固定资产", getFixedAssetList(infoList)).
		AddSeries("在建工程", getCipList(infoList)).
		AddSeries("员工薪酬", getStaffSalaryPayableList(infoList)).
		AddSeries("未分配利润", getUnAssignProfitList(infoList)).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)

	page := components.NewPage()
	page.SetLayout(components.PageFlexLayout)
	page.AddCharts(
		bar,
		bar1,
	)

	reportPath, _ := global.GetReportPath()
	fileName := filepath.Join(reportPath, "asset.html")
	f, err := os.Create(fileName)
	if err != nil {
		log.Error().Msgf("fail to create file, error is: %v", err)
		return
	}
	page.Render(io.MultiWriter(f))
}

func GenerateCashFlowReport(stockId string) {
	infoList := GetCashFlowReport(stockId)

	reportDateItems := getCashFlowReportDateList(infoList)

	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: "现金流量图"}),
		charts.WithYAxisOpts(opts.YAxis{
			AxisLabel: &opts.AxisLabel{Show: true, Formatter: "{value} 亿"},
		}),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
	)
	bar.SetXAxis(reportDateItems).
		AddSeries("销售", getSalesList(infoList)).
		AddSeries("成本", getOperateList(infoList)).
		AddSeries("投资", getInvestList(infoList)).
		AddSeries("金融", getFinanceList(infoList)).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
			}),
		)

	reportPath, _ := global.GetReportPath()
	fileName := filepath.Join(reportPath, "cash_flow.html")
	f, err := os.Create(fileName)
	if err != nil {
		log.Error().Msgf("fail to create file, error is: %v", err)
		return
	}
	bar.Render(f)
}
