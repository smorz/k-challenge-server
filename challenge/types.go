package challenge

import (
	"database/sql"
	"time"
)

type LastTrade struct {
	db   *sql.DB
	stmt *sql.Stmt
}

// ResultRow
//
// one row of result of query
type ResultRow struct {
	Name   string    `json:"name"`
	DateEn time.Time `json:"date_en"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
}

// JsonResponse
//
// api response json
type JsonResponse struct {
	OK     bool        `json:"ok"`
	Result []ResultRow `json:"result"`
}
