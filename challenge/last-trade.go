package challenge

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func (c *LastTradeServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	rows, err := c.stmt.Query()
	if err != nil {
		log.Panicln(err)
	}
	var (
		result []ResultRow
	)

	for rows.Next() {
		select {
		case <-r.Context().Done(): // request canceled
			return
		default:
			// go ahead!
			var (
				Name   string
				DateEn time.Time
				Open   float64
				High   float64
				Low    float64
				close  float64
			)

			if err := rows.Scan(&Name, &DateEn, &Open, &High, &Low, &close); err != nil {
				log.Panicln(err)
			}

			result = append(result,
				ResultRow{
					Name:   Name,
					DateEn: DateEn,
					Open:   Open,
					High:   High,
					Low:    Low,
				},
			)
		}
	}

	var response = JsonResponse{OK: true, Result: result}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Panicln(err)
	}
}

func NewLastTradeServer(db *sql.DB) *LastTradeServer {
	stmt, err := db.Prepare(`
	select name, DateEn, open, high, low, close
	from Instrument
	  left join (select distinct on (instrumentid) *
					from trade
					order by instrumentid, dateen desc) t
	on Instrument.id=t.instrumentid
	`)
	if err != nil {
		log.Fatal(err)
	}

	return &LastTradeServer{
		stmt: stmt,
	}
}
