package app

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
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

		res := make(map[string]float64, len(value["id"]))
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

func (a *App) addBalanceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// todo: help
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		r.Body.Close()

		requests := []map[string]interface{}{}
		err = json.Unmarshal(body, &requests)

		if err == nil {
			results := []map[string]interface{}{}
			for _, request := range requests {
				res, err := a.addForOneOperation(request)

				if err != nil {
					results = append(results, map[string]interface{}{fmt.Sprintf("error ->%s", request["id"]): err.Error()})
				} else {
					results = append(results, map[string]interface{}{
						fmt.Sprintf("result ->%s", request["id"]): map[string]interface{}{"balance": res},
					})
				}
			}

			ans, _ := json.Marshal(results)
			w.Header().Add("Content-Type", "application/json")
			w.Write(ans)
			return
		}

		request := make(map[string]interface{})
		err = json.Unmarshal(body, &request)
		if err != nil {
			ans, _ := json.Marshal(map[string]string{"error": "неверный json формат"})
			w.Header().Add("Content-Type", "application/json")
			w.Write(ans)
			return
		}

		result := make(map[string]interface{})
		res, err := a.addForOneOperation(request)
		if err != nil {
			result = map[string]interface{}{fmt.Sprintf("error ->%s", request["id"]): err.Error()}
		} else {
			result = map[string]interface{}{
				fmt.Sprintf("result ->%s", request["id"]): map[string]interface{}{"balance": res},
			}
		}

		ans, _ := json.Marshal(result)
		w.Header().Add("Content-Type", "application/json")
		w.Write(ans)
	}
}

func (a *App) addForOneOperation(request map[string]interface{}) (float64, error) {
	id, exists := request["id"]
	if !exists {
		return 0, errorMissingID
	}

	amount, exists := request["amount"]
	if !exists {
		return 0, errorMissingAmount
	}

	res, err := a.balances.Change(id.(string), amount.(float64))
	return res, err
}

func (a *App) transferBalanceHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// todo
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			return
		}
		r.Body.Close()

		requests := []map[string]interface{}{}
		err = json.Unmarshal(body, &requests)

		if err == nil {
			results := []map[string]string{}
			for _, request := range requests {
				a.transferForOneOperation(request)
				// todo
				results = append(results, map[string]string{fmt.Sprintf("result %s->%s", request["from"], request["to"]): "todo"})
			}
			ans, _ := json.Marshal(results)
			w.Header().Add("Content-Type", "application/json")
			w.Write(ans)
		} else {
			request := make(map[string]interface{})
			err = json.Unmarshal(body, &request)
			if err != nil {
				// todo error
			}
			ans, _ := json.Marshal(map[string]string{fmt.Sprintf("result %s->%s", request["from"], request["to"]): "todo"})
			w.Header().Add("Content-Type", "application/json")
			w.Write(ans)
		}
	}
}

func (a *App) transferForOneOperation(request map[string]interface{}) {

}
