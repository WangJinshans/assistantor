package model

type StockMarkInfo struct {
	StockId   string `json:"stockId"`
	StockName string `json:"stockName"`
	Message   string `json:"message"` // 备注信息
	Type      string `json:"type"`
	// 涨跌类型  up_and_down: 冲高回落上影线 up: 小阳线 uup： 中阳线  uuup:大阳线  down:小阴线 ddown: 中阴线 dddown:大阴线  face_pressure: 前期压力位
	// 涨跌类型  w1: 先上涨,上影线打到前高,回调  w2: 第二次冲击前高  w3: 第三次冲击前高
	// 涨跌类型  reverse: 反向情绪
	TargetPrice     float64 `json:"targetPrice"`     // 理想价格
	IsStockPositive int     `json:"isStockPositive"` // 是否有主动性
	IsBankPositive  int     `json:"isBankPositive"`  // 板块是否活跃
	CurrentPrice    float64 `json:"currentPrice"`
	BankStatus      int     `json:"bankStatus"`    // 板块状态  1： 主升一阶段 2：一阶段退潮 3：主升2阶段 4：二阶段退潮 5：主升三阶段 6：三阶段退潮
	Ready           int     `json:"ready"`         // 是否可立即买入
	BankName        string  `json:"bankName"`      // 板块
	BankId          string  `json:"bankId"`        // 板块Id
	Emotion         int     `json:"emotion"`       // 情绪值
	Link            string  `json:"link"`          // web链接
	StockPosition   string  `json:"stockPosition"` // 板块中的地位
	HeadStockId     string  `json:"headStockId"`   // 龙头
	HeadStockName   string  `json:"headStockName"` // 龙头名称
	HeadStockLink   string  `json:"headStockLink"` // 龙头链接
	TimeString      string  `json:"timeString"`
	TimeStamp       string  `json:"timestamp"`
}

type StockInfo struct {
	StockId      string  `json:"stock_id"`
	StockName    string  `json:"stock_name"`
	StockMarket  string  `json:"stock_market"` // 市场类型
	TimeString   string  `json:"time_string"`  // 时间
	OpenPrice    float64 `json:"open_price"`
	ClosePrice   float64 `json:"close_price"` // 收盘价
	CurrentPrice float64 `json:"current_price"`
	HighestPrice float64 `json:"highest_price"`
	LowestPrice  float64 `json:"lowest_price"`
	HighestRate  float64 `json:"highest_rate"`
	LowestRate   float64 `json:"lowest_rate"` // 通过最高、低价与开盘价计算得出
	CurrentRate  float64 `json:"current_rate"`
	Vol          int64   `json:"vol"`
	Money        float64 `json:"money"`
	ChangeRate   float64 `json:"change_rate"` // 换手率
	Amplitude    float64 `json:"amplitude"`   // 振幅
	Diff         float64 `json:"diff"`        // 距离前高
}

type StockShortInfo struct {
	StockId      string  `json:"stock_id"`
	StockName    string  `json:"stock_name"`
	CurrentPrice float64 `json:"current_price"`
	CurrentRate  float64 `json:"current_rate"`
	Vol          int64   `json:"vol"`
	Money        float64 `json:"money"`
}

type StockMinuteInfo struct {
	StockId      string
	StockName    string
	TimeStamp    int64
	AveragePrice float64 // 均价
	CurrentPrice float64
	HighestPrice float64
	LowestPrice  float64
	HighestRate  float64
	LowestRate   float64 // 通过最高、低价与开盘价计算得出
	CurrentRate  float64
	Diff         float64 // 距离前高
}
