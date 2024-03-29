package controller

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Config ...
func (h *HTTPChartHandler) Config(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	mlog.Debugw("config", "tr", trID)

	configResponse := struct {
		SupportedResolution    []string `json:"supported_resolutions"`
		SupportsGroupRequest   bool     `json:"supports_group_request"`
		SupportsMarks          bool     `json:"supports_marks"`
		SupportsSearch         bool     `json:"supports_search"`
		SupportsTimeScaleMarks bool     `json:"supports_timescale_marks"`
	}{
		[]string{"1", "5", "15", "30", "60", "1D", "1W", "1M"},
		true,
		false,
		false,
		false,
	}
	return c.JSON(http.StatusOK, configResponse)
}

// SymbolInfo ...
func (h *HTTPChartHandler) SymbolInfo(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	group := c.QueryParam("group")
	mlog.Debugw("symbol_info", "tr", trID, "group", group)

	symbolInfoResponse := struct {
		Symbol         []string `json:"symbol"`
		Description    []string `json:"description"`
		ExchangeListed string   `json:"exchange-listed"`
		ExchangeTraded string   `json:"exchange-traded"`
		MinMovement    int      `json:"minmovement"`
		MinMovement2   int      `json:"maxmovement2"`
		PriceScale     []int    `json:"pricescale"`
		HasDwm         bool     `json:"has-dwm"`
		HasIntraday    bool     `json:"has-intraday"`
		HasNoVolume    []bool   `json:"has-no-volume"`
		Type           []string `json:"type"`
		Ticker         []string `json:"ticker"`
		Timezone       string   `json:"timezone"`
		SessionRegular string   `json:"session-regular"`
	}{
		[]string{"AAPL", "MSFT", "SPX"},
		[]string{"Apple Inc", "Microsoft corp", "S&P 500 index"},
		"NYSE",
		"NYSE",
		1,
		0,
		[]int{1, 1, 100},
		true,
		true,
		[]bool{false, false, true},
		[]string{"stock", "stock", "index"},
		[]string{"AAPL~0", "MSFT~0", "$SPX500"},
		"America/New_York",
		"0900-1600",
	}
	return c.JSON(http.StatusOK, symbolInfoResponse)
}

// Symbols ...
func (h *HTTPChartHandler) Symbols(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	symbol := c.QueryParam("symbol")
	mlog.Debugw("symbols", "tr", trID, "symbol", symbol)

	symbolResponse := struct {
		Name                 string   `json:"name"`
		ExchangeTraded       string   `json:"exchange-traded"`
		ExchangeListed       string   `json:"exchange-listed"`
		Timezone             string   `json:"timezone"`
		Minmov               int      `json:"minmov"`
		Minmov2              int      `json:"minmov2"`
		Pointvalue           int      `json:"pointvalue"`
		Session              string   `json:"session"`
		HasIntraday          bool     `json:"has_intraday"`
		HasNoVolume          bool     `json:"has_no_volume"`
		Description          string   `json:"description"`
		Type                 string   `json:"type"`
		SupportedResolutions []string `json:"supported_resolutions"`
		Pricescale           int      `json:"pricescale"`
		Ticker               string   `json:"ticker"`
	}{
		"AAL",
		"NasdaqNM",
		"NasdaqNM",
		"America/New_York",
		1,
		0,
		1,
		"0930-1630",
		false,
		false,
		"American Airlines Group Inc.",
		"stock",
		[]string{"D", "2D", "3D", "W", "3W", "M", "6M"},
		100,
		"AAL",
	}
	return c.JSON(http.StatusOK, symbolResponse)
}

// Search ...
func (h *HTTPChartHandler) Search(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	type options struct {
		Query     string `query:"query"`
		TypeParam string `query:"type"`
		Exchange  string `query:"exchange"`
		Limit     string `query:"limit"`
	}
	option := new(options)
	if err := c.Bind(option); err != nil {
		return c.JSON(http.StatusInternalServerError, TiffanyBlueStatus{
			ResultCode: "1000",
			ResultMsg:  err.Error(),
			TRID:       trID,
		})
	}
	mlog.Debugw("search", "tr", trID, "query", option.Query, "type", option.TypeParam, "exchange", option.Exchange, "limit", option.Limit)

	searchResponse := []struct {
		Symbol      string `json:"symbol"`
		FullName    string `json:"full_name"`
		Description string `json:"description"`
		Exchange    string `json:"exchange"`
		Type        string `json:"type"`
	}{
		{
			"AA",
			"AA",
			"Alcoa Inc.",
			"NYSE",
			"stock",
		},
	}
	return c.JSON(http.StatusOK, searchResponse)
}

// History ...
func (h *HTTPChartHandler) History(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	type options struct {
		Symbol     string `query:"symbol"`
		From       string `query:"from"`
		To         string `query:"to"`
		Resolution string `query:"resolution"`
	}
	option := new(options)
	if err := c.Bind(option); err != nil {
		return c.JSON(http.StatusInternalServerError, TiffanyBlueStatus{
			ResultCode: "1000",
			ResultMsg:  err.Error(),
			TRID:       trID,
		})
	}
	mlog.Debugw("history", "tr", trID, "symbol", option.Symbol, "from", option.From, "to", option.To, "resolution", option.Resolution)

	historyResponse := struct {
		T []int     `json:"t"`
		O []float64 `json:"o"`
		H []float64 `json:"h"`
		L []float64 `json:"l"`
		C []float64 `json:"c"`
		V []int     `json:"v"`
		S string    `json:"s"`
	}{
		[]int{1500940800, 1501027200, 1501113600, 1501200000, 1501459200, 1501545600, 1501632000, 1501718400, 1501804800, 1502150400, 1502236800, 1502323200, 1502409600, 1502668800, 1502755200, 1502841600, 1502928000, 1503014400, 1503273600, 1503360000, 1503446400, 1503532800, 1503619200, 1503878400, 1503964800, 1504051200, 1504137600, 1504224000, 1504569600, 1504656000, 1504742400, 1504828800, 1505088000, 1505174400, 1505260800, 1505347200, 1505433600, 1505692800, 1505779200, 1505865600, 1505952000, 1506038400, 1506297600, 1506384000, 1506470400, 1506556800, 1506643200, 1506902400, 1506988800, 1507075200, 1507161600, 1507248000, 1507507200, 1507593600, 1507680000, 1507766400, 1507852800, 1508112000, 1508198400, 1508284800, 1508371200, 1508457600, 1508716800, 1508803200, 1508889600, 1508976000, 1509062400, 1509321600, 1509408000, 1509494400, 1509580800, 1509667200, 1509926400, 1510012800, 1510185600, 1510272000, 1510531200, 1510617600, 1510704000, 1510790400, 1510876800, 1511136000, 1511222400, 1511308800, 1511481600, 1511740800, 1511827200, 1511913600, 1512000000, 1512086400, 1512345600, 1512432000, 1512518400, 1512604800, 1512691200, 1512950400, 1513036800, 1513123200, 1513209600, 1513296000, 1513555200, 1513641600, 1513728000, 1513814400, 1513900800, 1514246400, 1514332800, 1514419200, 1514505600, 1514851200, 1514937600, 1515024000, 1515110400, 1515369600, 1515456000, 1515542400, 1515628800, 1515715200, 1516060800, 1516147200, 1516233600, 1516320000, 1516579200, 1516665600, 1516752000, 1516838400, 1516924800, 1517184000, 1517270400, 1517356800, 1517443200, 1517529600, 1517788800, 1517875200, 1517961600, 1518048000, 1518134400, 1518393600, 1518480000, 1518566400, 1518652800, 1518739200, 1519084800, 1519171200, 1519257600, 1519344000, 1519603200, 1519689600, 1519776000, 1519862400, 1519948800, 1520208000, 1520294400, 1520380800, 1520467200, 1520553600, 1520812800, 1520899200, 1520985600, 1521072000, 1521158400, 1521417600, 1521504000, 1521590400, 1521676800, 1521763200, 1522022400, 1522108800},
		[]float64{151.8, 153.35, 153.75, 149.89, 149.9, 149.1, 159.28, 157.05, 156.07, 158.6, 159.26, 159.9, 156.6, 159.32, 160.66, 161.94, 160.52, 157.86, 157.5, 158.23, 159.07, 160.43, 159.65, 160.14, 160.1, 163.8, 163.64, 164.8, 163.75, 162.71, 162.09, 160.86, 160.5, 162.61, 159.87, 158.99, 158.47, 160.11, 159.51, 157.9, 155.8, 152.02, 149.99, 151.78, 153.8, 153.89, 153.21, 154.26, 154.01, 153.63, 154.18, 154.97, 155.81, 156.055, 155.97, 156.35, 156.73, 157.9, 159.78, 160.42, 156.75, 156.61, 156.89, 156.29, 156.91, 157.23, 159.29, 163.89, 167.9, 169.87, 167.64, 174, 172.365, 173.91, 175.11, 175.11, 173.5, 173.04, 169.97, 171.18, 171.04, 170.29, 170.78, 173.36, 175.1, 175.05, 174.3, 172.63, 170.43, 169.95, 172.48, 169.06, 167.5, 169.03, 170.49, 169.2, 172.15, 172.5, 172.4, 173.63, 174.88, 175.03, 174.87, 174.17, 174.68, 170.8, 170.1, 171, 170.52, 170.16, 172.53, 172.54, 173.44, 174.35, 174.55, 173.16, 174.59, 176.18, 177.9, 176.15, 179.37, 178.61, 177.3, 177.3, 177.25, 174.505, 172, 170.16, 165.525, 166.87, 167.165, 166, 159.1, 154.83, 163.085, 160.29, 157.07, 158.5, 161.95, 163.045, 169.79, 172.36, 172.05, 172.83, 171.8, 173.67, 176.35, 179.1, 179.26, 178.54, 172.8, 175.21, 177.91, 174.94, 175.48, 177.96, 180.29, 182.59, 180.32, 178.5, 178.65, 177.32, 175.24, 175.04, 170, 168.39, 168.07, 173.68},
		[]float64{153.84, 153.93, 153.99, 150.23, 150.33, 150.22, 159.75, 157.21, 157.4, 161.83, 161.27, 160, 158.5728, 160.21, 162.195, 162.51, 160.71, 159.5, 157.89, 160, 160.47, 160.74, 160.56, 162, 163.12, 163.89, 164.52, 164.94, 164.25, 162.99, 162.24, 161.15, 162.05, 163.96, 159.96, 159.4, 160.97, 160.5, 159.77, 158.26, 155.8, 152.27, 151.83, 153.92, 154.7189, 154.28, 154.13, 154.45, 155.09, 153.86, 155.44, 155.49, 156.73, 158, 156.98, 157.37, 157.28, 160, 160.87, 160.71, 157.08, 157.75, 157.69, 157.42, 157.55, 157.8295, 163.6, 168.07, 169.6499, 169.94, 168.5, 174.26, 174.99, 175.25, 176.095, 175.38, 174.5, 173.48, 170.3197, 171.87, 171.39, 170.56, 173.7, 175, 175.5, 175.08, 174.87, 172.92, 172.14, 171.67, 172.62, 171.52, 170.2047, 170.44, 171, 172.89, 172.39, 173.54, 173.13, 174.17, 177.2, 175.39, 175.42, 176.02, 175.424, 171.47, 170.78, 171.85, 170.59, 172.3, 174.55, 173.47, 175.37, 175.61, 175.06, 174.3, 175.4886, 177.36, 179.39, 179.25, 180.1, 179.58, 177.78, 179.44, 177.3, 174.95, 172, 170.16, 167.37, 168.4417, 168.62, 166.8, 163.88, 163.72, 163.4, 161, 157.89, 163.89, 164.75, 167.54, 173.09, 174.82, 174.26, 174.12, 173.95, 175.65, 179.39, 180.48, 180.615, 179.775, 176.3, 177.74, 178.25, 175.85, 177.12, 180, 182.39, 183.5, 180.52, 180.24, 179.12, 177.47, 176.8, 175.09, 172.68, 169.92, 173.1, 175.15},
		[]float64{151.8, 153.06, 147.3, 149.19, 148.13, 148.41, 156.16, 155.02, 155.69, 158.27, 159.11, 154.63, 156.07, 158.75, 160.14, 160.15, 157.84, 156.72, 155.1101, 158.02, 158.88, 158.55, 159.27, 159.93, 160, 162.61, 163.48, 163.63, 160.56, 160.52, 160.36, 158.53, 159.89, 158.77, 157.91, 158.09, 158, 157.995, 158.44, 153.83, 152.75, 150.56, 149.16, 151.69, 153.54, 152.7, 152, 152.72, 153.91, 152.46, 154.05, 154.56, 155.485, 155.1, 155.75, 155.7299, 156.41, 157.65, 159.23, 159.6, 155.02, 155.96, 155.5, 156.2, 155.27, 156.78, 158.7, 163.72, 166.94, 165.61, 165.28, 171.12, 171.72, 173.6, 173.14, 174.27, 173.4, 171.18, 168.38, 170.3, 169.64, 169.56, 170.78, 173.05, 174.6459, 173.34, 171.86, 167.16, 168.44, 168.5, 169.63, 168.4, 166.46, 168.91, 168.82, 168.79, 171.461, 172, 171.65, 172.46, 174.86, 174.09, 173.25, 174.1, 174.5, 169.679, 169.71, 170.48, 169.22, 169.26, 171.96, 172.08, 173.05, 173.93, 173.41, 173, 174.49, 175.65, 176.14, 175.07, 178.25, 177.41, 176.6016, 176.82, 173.2, 170.53, 170.06, 167.07, 164.7, 166.5, 166.76, 160.1, 156, 154, 159.0685, 155.03, 150.24, 157.51, 161.65, 162.88, 169, 171.77, 171.42, 171.01, 171.71, 173.54, 176.21, 178.16, 178.05, 172.66, 172.45, 174.52, 176.13, 174.27, 175.07, 177.39, 180.21, 179.24, 177.81, 178.0701, 177.62, 173.66, 174.94, 171.26, 168.6, 164.94, 166.44, 166.92},
		[]float64{152.74, 153.46, 150.56, 149.5, 148.85, 150.05, 157.14, 155.57, 156.39, 160.08, 161.06, 155.27, 157.48, 159.85, 161.6, 160.95, 157.87, 157.5, 157.21, 159.78, 159.98, 159.27, 159.86, 161.47, 162.91, 163.35, 164, 164.05, 162.08, 161.91, 161.26, 158.63, 161.5, 160.82, 159.65, 158.28, 159.88, 158.67, 158.73, 156.07, 153.39, 151.89, 150.55, 153.14, 154.23, 153.28, 154.12, 153.81, 154.48, 153.4508, 155.39, 155.3, 155.84, 155.9, 156.55, 156, 156.99, 159.88, 160.47, 159.76, 155.98, 156.16, 156.17, 157.1, 156.405, 157.41, 163.05, 166.72, 169.04, 166.89, 168.11, 172.5, 174.25, 174.81, 175.88, 174.67, 173.97, 171.34, 169.08, 171.1, 170.15, 169.98, 173.14, 174.96, 174.97, 174.09, 173.07, 169.48, 171.85, 171.05, 169.8, 169.64, 169.01, 169.452, 169.37, 172.67, 171.7, 172.27, 172.22, 173.87, 176.42, 174.54, 174.35, 175.01, 175.01, 170.57, 170.6, 171.08, 169.23, 172.26, 172.23, 173.03, 175, 174.35, 174.33, 174.29, 175.28, 177.09, 176.19, 179.1, 179.26, 178.46, 177, 177.04, 174.22, 171.11, 171.51, 167.96, 166.97, 167.43, 167.78, 160.37, 157.49, 163.03, 159.54, 155.32, 155.97, 162.71, 164.34, 167.37, 172.99, 172.43, 171.85, 171.07, 172.6, 175.555, 178.97, 178.39, 178.12, 175, 176.21, 176.82, 176.67, 175.03, 176.94, 179.98, 181.72, 179.97, 178.44, 178.65, 178.02, 175.3, 175.24, 171.27, 168.845, 164.94, 172.77, 168.34},
		[]int{18612649, 15172136, 32175875, 16832947, 19422655, 24725526, 69222793, 26000738, 20349532, 35775675, 25640394, 39081017, 25943187, 21754810, 27936774, 27321761, 26925694, 27012525, 26145653, 21297812, 19198189, 19029621, 25015218, 25279674, 29307862, 26973946, 26412439, 16508568, 29317054, 21179047, 21722995, 28183159, 31028926, 71139119, 44393752, 23073646, 48203642, 27939718, 20347352, 51693239, 36643382, 46114424, 43922334, 35470985, 24959552, 21896592, 25856530, 18524860, 16146388, 19844177, 21032800, 16423749, 16200129, 15456331, 16607693, 16045720, 16287608, 23894630, 18816438, 16158659, 42111326, 23612246, 21654461, 17137731, 20126554, 16751691, 43904150, 43923292, 35474672, 33100847, 32710040, 58683826, 34242566, 23910914, 28636531, 25061183, 16828025, 23588451, 28702351, 23497326, 21665811, 15974387, 24875471, 24997274, 14026519, 20536313, 25468442, 40788324, 40172368, 39590080, 32115052, 27008428, 28224357, 24469613, 23096872, 33092051, 18945457, 23142242, 20219307, 37054632, 28831533, 27078872, 23000392, 20356826, 16052615, 32968167, 21672062, 15997739, 25643711, 25048048, 28819653, 22211345, 23016177, 20134092, 21262614, 23589129, 17523256, 25039531, 29159005, 32752734, 30234512, 30827809, 26023683, 31702531, 50562257, 39661804, 37121805, 48434424, 45137026, 30984099, 38099665, 85436075, 66090446, 66625484, 50852130, 49594129, 66723743, 60560145, 32104756, 39669178, 50609595, 39638793, 33531012, 35833514, 30504116, 33329232, 36886432, 38685165, 33604574, 48801970, 38453950, 28401366, 23788506, 31703462, 23163767, 31385134, 32055405, 31168404, 29075469, 22584565, 36836456, 32804695, 19314039, 35247358, 41051076, 40248954, 36272617, 38962839},
		"ok",
	}
	return c.JSON(http.StatusOK, historyResponse)
}

// Marks ...
func (h *HTTPChartHandler) Marks(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	type options struct {
		Symbol     string `query:"symbol"`
		From       string `query:"from"`
		To         string `query:"to"`
		Resolution string `query:"resolution"`
	}
	option := new(options)
	if err := c.Bind(option); err != nil {
		return c.JSON(http.StatusInternalServerError, TiffanyBlueStatus{
			ResultCode: "1000",
			ResultMsg:  err.Error(),
			TRID:       trID,
		})
	}
	mlog.Debugw("marks", "tr", trID, "symbol", option.Symbol, "from", option.From, "to", option.To, "resolution", option.Resolution)

	marksResponse := struct {
		ID             []int    `json:"id"`
		Time           []int    `json:"time"`
		Color          []string `json:"color"`
		Text           []string `json:"text"`
		Label          []string `json:"label"`
		LabelFontColor []string `json:"labelFontColor"`
		MinSize        []int    `json:"minSize"`
	}{
		[]int{0, 1, 2, 3, 4, 5},
		[]int{1531958400, 1531612800, 1531353600, 1531353600, 1530662400, 1529366400},
		[]string{"red", "blue", "green", "red", "blue", "green"},
		[]string{
			"Today",
			"4 days back",
			"7 days back + Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"7 days back once again",
			"15 days back",
			"30 days back",
		},
		[]string{"A", "B", "CORE", "D", "EURO", "F"},
		[]string{"white", "white", "red", "#FFFFFF", "white", "#000"},
		[]int{14, 28, 7, 40, 7, 14},
	}
	return c.JSON(http.StatusOK, marksResponse)
}

// TimeScale ...
func (h *HTTPChartHandler) TimeScale(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	type options struct {
		Symbol     string `query:"symbol"`
		From       string `query:"from"`
		To         string `query:"to"`
		Resolution string `query:"resolution"`
	}
	option := new(options)
	if err := c.Bind(option); err != nil {
		return c.JSON(http.StatusInternalServerError, TiffanyBlueStatus{
			ResultCode: "1000",
			ResultMsg:  err.Error(),
			TRID:       trID,
		})
	}
	mlog.Debugw("timescale_marks", "tr", trID, "symbol", option.Symbol, "from", option.From, "to", option.To, "resolution", option.Resolution)

	timescaleResponse := []struct {
		ID      string   `json:"id"`
		Time    int      `json:"time"`
		Color   string   `json:"color"`
		Label   string   `json:"label"`
		Tooltip []string `json:"tooltip"`
	}{
		{"tsm1", 1531958400, "red", "A", []string{}},
		{"tsm2", 1531612800, "blue", "D", []string{"Dividends: $0.56", "Date: Sun Jul 15 2018"}},
		{"tsm3", 1531353600, "blue", "D", []string{"Dividends: $3.46", "Date: Thu Jul 12 2018"}},
		{"tsm4", 1530662400, "#999999", "E", []string{"Earnings: $3.44", "Estimate: $3.60"}},
		{"tsm7", 1529366400, "red", "E", []string{"Earnings: $5.40", "Estimate: $5.00"}},
	}

	return c.JSON(http.StatusOK, timescaleResponse)
}

// Time ...
func (h *HTTPChartHandler) Time(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)

	mlog.Debugw("time", "tr", trID)

	return c.JSON(http.StatusOK, int32(time.Now().Unix()))
}

// Quotes ...
func (h *HTTPChartHandler) Quotes(c echo.Context) (err error) {

	trID := c.Response().Header().Get(echo.HeaderXRequestID)
	type options struct {
		Symbols []string `query:"symbols"`
	}
	option := new(options)
	if err := c.Bind(option); err != nil {
		return c.JSON(http.StatusInternalServerError, TiffanyBlueStatus{
			ResultCode: "1000",
			ResultMsg:  err.Error(),
			TRID:       trID,
		})
	}
	mlog.Debugw("quotes", "tr", trID, "symbols", option.Symbols)

	/*
		quotesResponse := struct {
			S string `json:"s"`
			D []struct {
				S string `json:"s"`
				N string `json:"n"`
				V struct {
					Ch             int     `json:"ch"`
					Chp            int     `json:"chp"`
					ShortName      string  `json:"short_name"`
					Exchange       string  `json:"exchange"`
					OriginalName   string  `json:"original_name"`
					Description    string  `json:"description"`
					Lp             float64 `json:"lp"`
					Ask            float64 `json:"ask"`
					Bid            float64 `json:"bid"`
					OpenPrice      float64 `json:"open_price"`
					HighPrice      float64 `json:"high_price"`
					LowPrice       float64 `json:"low_price"`
					PrevClosePrice float64 `json:"prev_close_price"`
					Volume         float64 `json:"volume"`
				} `json:"v"`
			} `json:"d"`
			Source string `json:"source"`
		}{
			S: "ok",
			D: []struct{
				{S: "ok", N: "NYSE:AA", V: {0, 0, "AA", "", "NYSE:AA", "NYSE:AA", 46.25, 46.25, 46.25, 46.25, 46.25, 46.25, 45.77, 46.25}},
				{S: "ok", N: "NYSE:F", V: {0, 0, "F", "", "NYSE:F", "NYSE:F", 12.03, 12.03, 12.03, 12.03, 12.03, 12.03, 12.015, 12.03}},
				{S: "ok", N: "NasdaqNM:AAPL", V: {0, 0, "AAPL", "", "NasdaqNM:AAPL", "NasdaqNM:AAPL", 173.68, 173.68, 173.68, 173.68, 173.68, 173.68, 172.77, 173.68}},
			},
			Source: "Quandl",
		}
	*/

	return c.JSON(http.StatusOK, TiffanyBlueStatus{
		ResultCode: "0000",
		ResultMsg:  "Request OK",
		TRID:       trID,
	})
}
