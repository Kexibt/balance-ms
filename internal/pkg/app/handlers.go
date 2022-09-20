package app

import (
	"encoding/json"
	"net/http"
)

func (a *App) mainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// всю инфу скинуть
	case "POST":
		// todo
	}
}

func (a *App) getBalanceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		value := r.URL.Query()
		if len(value["id"]) < 1 {
			ans, _ := json.Marshal(map[string]string{"error": "задайте id через параметры запроса"})
			w.Header().Add("Content-Type", "application/json")
			w.Write(ans)
			return
		}

		res := make(map[string]uint, len(value["id"]))
		for _, id := range value["id"] {
			res[id] = a.balances.GetByUID(id)
		}

		w.Header().Add("Content-Type", "application/json")
		ans, _ := json.Marshal(res)
		w.Write(ans)
	case "POST":
		// todo
	}
}
