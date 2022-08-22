package helper

import (
	"github.com/shellrean-playground/hotel-be-common-util/constant"
	"net/http"
)

func ExoHttpStatus(errCode string) int {
	statuses := map[string]int{
		constant.SUCCESS:            http.StatusOK,
		constant.INVALID_PARAM:      http.StatusBadRequest,
		constant.ERROR_GENERAL:      http.StatusInternalServerError,
		constant.ROOM_NOT_AVAILABLE: http.StatusBadRequest,
		constant.DUPLICATION_DATA:   http.StatusBadRequest,
	}
	if val, ok := statuses[errCode]; ok {
		return val
	}
	return http.StatusInternalServerError
}
