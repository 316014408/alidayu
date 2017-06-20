package alidayu

import (
	"time"
)

func QuerySMS(rec_num, query_date, current_page, page_size string) (success bool, response string) {
	if rec_num == "" || query_date == "" || current_page == "" {
		return false, "Parameter not complete"
	}
	if page_size == "" {
		page_size = "10"
	}

	params := make(map[string]string)
	params["app_key"] = AppKey
	params["format"] = "json"
	params["method"] = Method_SendSMS
	params["sign_method"] = "md5"
	params["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	params["v"] = "2.0"
	params["query_date"] = query_date     // 短信发送日期，支持近30天记录查询，格式yyyyMMdd
	params["rec_num"] = rec_num           // 短信接收号码
	params["current_page"] = current_page // 分页参数,页码
	params["page_size"] = page_size       // 分页参数，每页数量。最大值50

	return DoPost(params)

}
