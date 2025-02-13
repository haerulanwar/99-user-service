package utils

import (
	"errors"
	"strconv"
)

func ParsePaginationParams(pageNumStr, pageSizeStr string) (int, int, error) {
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil || pageNum < 1 {
		return 0, 0, errors.New("invalid page number")
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		return 0, 0, errors.New("invalid page size")
	}

	offSet := (pageNum - 1) * pageSize
	return offSet, pageSize, nil
}
