package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func getPagingParams(r *http.Request) (page, pageSize int) {
	pageSizeStr, ok := r.URL.Query()["page_size"]
	if ok && len(pageSizeStr) > 0 {
		pageSizeTemp, err := strconv.Atoi(pageSizeStr[0])
		if err != nil {
			pageSizeTemp = 5
		}
		pageSize = pageSizeTemp
	} else {
		pageSize = 5
	}

	pageStr, ok := r.URL.Query()["page"]
	if ok && len(pageStr) > 0 {
		pageTemp, err := strconv.Atoi(pageStr[0])
		if err != nil {
			pageTemp = 1
		}
		page = pageTemp
	} else {
		page = 1
	}

	return page, pageSize
}
func SendAsJson(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func sendAsJson2(w http.ResponseWriter, status int, payload interface{}) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	encoder.Encode(payload)
}
