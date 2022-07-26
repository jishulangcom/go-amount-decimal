package amountdecimal

const (
	COIN_BTC         = "BTC" // BTC比特币
	COIN_BTC_DECIMAL = 8

	COIN_BTC_USDT         = "BTC_USDT"
	COIN_BTC_USDT_DECIMAL = 8

	COIN_ETH         = "ETH" // ETH以太坊
	COIN_ETH_DECIMAL = 18

	COIN_ETH_USDT         = "ETH_USDT"
	COIN_ETH_USDT_DECIMAL = 6

	COIN_ONT_ONT2         = "ONT_ONT2"
	COIN_ONT_ONT2_DECIMAL = 9

	COIN_ONT_ONG2         = "ONT_ONG2"
	COIN_ONT_ONG2_DECIMAL = 18

	COIN_TONT_ONG2         = "TONT_ONG2"
	COIN_TONT_ONG2_DECIMAL = 18

	COIN_TONT_ONT2         = "TONT_ONT2"
	COIN_TONT_ONT2_DECIMAL = 9

	COIN_TRON         = "TRON"
	COIN_TRON_DECIMAL = 6

	COIN_TETH         = "TETH"
	COIN_TETH_DECIMAL = 18

	COIN_TETH_CTT         = "TETH_CTT"
	COIN_TETH_CTT_DECIMAL = 18

	COIN_XLM         = "XLM" // XLM恒星币
	COIN_XLM_DECIMAL = 7

	COIN_XTN         = "XTN"
	COIN_XTN_DECIMAL = 8
)

type CoinInfo struct {
	Code    string `json:"code"`    // 标识
	Decimal int    `json:"decimal"` // 精度
}

var CoinMap map[string]CoinInfo
var CoinList []CoinInfo

func init() {
	CoinMap = make(map[string]CoinInfo)

	// BTC
	CoinMap[COIN_BTC] = CoinInfo{
		Code:    COIN_BTC,
		Decimal: COIN_BTC_DECIMAL,
	}

	// BTC_USDT
	CoinMap[COIN_BTC_USDT] = CoinInfo{
		Code:    COIN_BTC_USDT,
		Decimal: COIN_BTC_USDT_DECIMAL,
	}

	// ETH
	CoinMap[COIN_ETH] = CoinInfo{
		Code:    COIN_ETH,
		Decimal: COIN_ETH_DECIMAL,
	}

	// ETH_USDT
	CoinMap[COIN_ETH_USDT] = CoinInfo{
		Code:    COIN_ETH_USDT,
		Decimal: COIN_ETH_USDT_DECIMAL,
	}

	// ONT_ONT2
	CoinMap[COIN_ONT_ONT2] = CoinInfo{
		Code:    COIN_ONT_ONT2,
		Decimal: COIN_ONT_ONT2_DECIMAL,
	}

	// ONT_ONG2
	CoinMap[COIN_ONT_ONG2] = CoinInfo{
		Code:    COIN_ONT_ONG2,
		Decimal: COIN_ONT_ONG2_DECIMAL,
	}

	// TONT_ONG2
	CoinMap[COIN_TONT_ONG2] = CoinInfo{
		Code:    COIN_TONT_ONG2,
		Decimal: COIN_TONT_ONG2_DECIMAL,
	}

	// TONT_ONG2
	CoinMap[COIN_TONT_ONT2] = CoinInfo{
		Code:    COIN_TONT_ONT2,
		Decimal: COIN_TONT_ONT2_DECIMAL,
	}

	// TRON
	CoinMap[COIN_TRON] = CoinInfo{
		Code:    COIN_TRON,
		Decimal: COIN_TRON_DECIMAL,
	}

	// TETH
	CoinMap[COIN_TETH] = CoinInfo{
		Code:    COIN_TETH,
		Decimal: COIN_TETH_DECIMAL,
	}

	// TETH_CTT
	CoinMap[COIN_TETH_CTT] = CoinInfo{
		Code:    COIN_TETH_CTT,
		Decimal: COIN_TETH_CTT_DECIMAL,
	}

	// XLM
	CoinMap[COIN_XLM] = CoinInfo{
		Code:    COIN_XLM,
		Decimal: COIN_XLM_DECIMAL,
	}

	// XTN
	CoinMap[COIN_XTN] = CoinInfo{
		Code:    COIN_XTN,
		Decimal: COIN_XTN_DECIMAL,
	}

}
