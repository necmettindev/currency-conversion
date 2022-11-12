package finance

type FinanceAPI struct {
	Spark struct {
		Result []struct {
			Symbol   string `json:"symbol"`
			Response []struct {
				Meta struct {
					Currency             string  `json:"currency"`
					Symbol               string  `json:"symbol"`
					ExchangeName         string  `json:"exchangeName"`
					InstrumentType       string  `json:"instrumentType"`
					FirstTradeDate       int     `json:"firstTradeDate"`
					RegularMarketTime    int     `json:"regularMarketTime"`
					Gmtoffset            int     `json:"gmtoffset"`
					Timezone             string  `json:"timezone"`
					ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
					RegularMarketPrice   float64 `json:"regularMarketPrice"`
					ChartPreviousClose   float64 `json:"chartPreviousClose"`
					PreviousClose        float64 `json:"previousClose"`
					Scale                int     `json:"scale"`
					PriceHint            int     `json:"priceHint"`
					CurrentTradingPeriod struct {
						Pre struct {
							Timezone  string `json:"timezone"`
							Start     int    `json:"start"`
							End       int    `json:"end"`
							Gmtoffset int    `json:"gmtoffset"`
						} `json:"pre"`
						Regular struct {
							Timezone  string `json:"timezone"`
							Start     int    `json:"start"`
							End       int    `json:"end"`
							Gmtoffset int    `json:"gmtoffset"`
						} `json:"regular"`
						Post struct {
							Timezone  string `json:"timezone"`
							Start     int    `json:"start"`
							End       int    `json:"end"`
							Gmtoffset int    `json:"gmtoffset"`
						} `json:"post"`
					} `json:"currentTradingPeriod"`
					TradingPeriods [][]struct {
						Timezone  string `json:"timezone"`
						Start     int    `json:"start"`
						End       int    `json:"end"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"tradingPeriods"`
					DataGranularity string   `json:"dataGranularity"`
					Range           string   `json:"range"`
					ValidRanges     []string `json:"validRanges"`
				} `json:"meta"`
				Timestamp  []int `json:"timestamp"`
				Indicators struct {
					Quote []struct {
						Close []float64 `json:"close"`
					} `json:"quote"`
				} `json:"indicators"`
			} `json:"response"`
		} `json:"result"`
		Error interface{} `json:"error"`
	} `json:"spark"`
}
