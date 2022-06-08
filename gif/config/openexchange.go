package config

import (
	"fmt"
	"time"
)

type OpenExchange struct {
	BaseApi string
	AppID   string
}

const basePAth = `https://openexchangerates.org/api/%s?app_id=%s`

func NewOpenExchange() OpenExchange {
	return OpenExchange{
		BaseApi: "https://openexchangerates.org/api/",
		AppID:   "5e408028b96c40fe80bcf04bf1f6bcca",
	}
}

//https://openexchangerates.org/api/latest.json?app_id=5e408028b96c40fe80bcf04bf1f6bcca
func (o OpenExchange) Latest() string {
	url := fmt.Sprintf(basePAth, "latest.json", o.AppID)
	return url
}

//https://openexchangerates.org/api/currencies.json?app_id=5e408028b96c40fe80bcf04bf1f6bcca
func (o OpenExchange) Currencies() string {
	url := fmt.Sprintf(basePAth, "currencies.json", o.AppID)
	return url
}

//https://openexchangerates.org/api/historical/2022-06-01.json?app_id=5e408028b96c40fe80bcf04bf1f6bcca
func (o OpenExchange) Yesterday() string {
	today := time.Now()
	yesterday := today.AddDate(0, 0, -1)
	t := yesterday.Format("2006-01-02")
	url := o.BaseApi + "historical/" + t + ".json?app_id=" + o.AppID
	return url
}
