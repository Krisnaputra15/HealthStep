package entity

import (
	"math"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ResponsePagination struct {
	PageNumber int `json:"page"`
	ItemInPage int `json:"item_in_page"`
	TotalPage  int `json:"total_page"`
	TotalItem  int `json:"total_item"`
	Offset     int `json:"-"`
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page == 0 {
		  page = 1
		}
	
		pageSize := 20
		switch {
		case pageSize > 100:
		  pageSize = 100
		case pageSize <= 0:
		  pageSize = 10
		}
	
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	  }
}

func HandlePagination(page, limit, itemCount int) ResponsePagination {
	var (
		totalPages, offset, itemInPage int
		pagination                     ResponsePagination
	)

	if page == 0 {
		page = 1
	}

	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * 5
	}

	if limit == 0 {
		itemInPage = 3
	} else {
		itemInPage = limit
	}

	if limit == 0 {
		totalPages = int(math.Ceil(float64(itemCount) / float64(3)))
	} else {
		totalPages = int(math.Ceil(float64(itemCount) / float64(limit)))
	}

	pagination = ResponsePagination{
		PageNumber: page,
		ItemInPage: itemInPage,
		TotalPage:  totalPages,
		TotalItem:  itemCount,
		Offset:     offset,
	}

	return pagination
}