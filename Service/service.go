package Service

import (
	"github.com/guonaihong/gout"
)

const IdoHttpConfig = "http://test2.idchats.com:30004/"

type TIdoList struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func GetSystemHaveIdo() (result []string) {
	var nowExistIdo TIdoList
	err := gout.GET(IdoHttpConfig + "process?id=contract").BindJSON(&nowExistIdo).Do()
	if err != nil || nowExistIdo.Code != 0 {
		return nil
	}
	for _, value := range nowExistIdo.Data {
		result = append(result, value)
	}
	return
}
