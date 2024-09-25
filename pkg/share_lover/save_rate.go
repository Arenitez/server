package share_lover

import (
	"fmt"
	"log"
	"server/pkg/rate"
)

func SaveRate() {
	// ドル円のレートを取得
	rate, err := rate.CallExchangeRateAPI(rate.USD, rate.JPY)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rate)
	return
}
