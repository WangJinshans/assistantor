package economy

type ProfitInfo struct {
	ReportDate         string  `position:"x"`
	TotalOperateIncome float64 `position:"y"`
	TotalOperateCost   float64 `position:"y"`
	OperateCost        float64 `position:"y"`
	OperateTax         float64 `position:"y"`
	SaleExpense        float64 `position:"y"`
	ManageExpense      float64 `position:"y"`
	ResearchExpense    float64 `position:"y"`
	FinanceExpense     float64 `position:"y"`
	InvestJointIncome  float64 `position:"y"`
}
