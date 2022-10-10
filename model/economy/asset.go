package economy

type AssetInfo struct {
	ReportDate          string  // 报告日期
	MonetaryFunds       float64 // 货币资金
	NoteAccountsReceive float64 // 应收账款
	Prepayment          float64 // 预收款
	Inventory           float64 // 存货
	FixedAsset          float64 // 固定资产
	Cip                 float64 // 在建工程
	StaffSalaryPayable  float64 // 员工薪酬
	UnAssignProfit      float64 // 未分配利润
}
