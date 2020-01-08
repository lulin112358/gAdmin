package serializer

import (
	"math"
	"os"
	"strconv"
)

type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
}

type ResponseList struct {
	Items   interface{} `json:"items"`
	Count   int         `json:"count"`
	AllPage int         `json:"all_page"`
	PerPage int         `json:"per_page"`
}

// 构建通用列表结构体
func BuildList(items interface{}, count int) ResponseList {
	perPage, _ := strconv.Atoi(os.Getenv("PER_PAGE"))
	return ResponseList{
		Items:   items,
		Count:   count,
		AllPage: int(math.Ceil(float64(count / perPage))),
		PerPage: perPage,
	}
}
