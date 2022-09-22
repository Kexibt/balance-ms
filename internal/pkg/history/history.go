package history

// import (
// 	"fmt"
// 	"sync"
// )

// type History struct {
// 	data  map[string][]string
// 	mutex sync.RWMutex
// }

// func (h *History) GetByUID(userID string) []string {
// 	h.mutex.RLock()
// 	defer h.mutex.RUnlock()

// 	copy := make([]string, 0, len(h.data[userID]))
// 	copy = append(copy, h.data[userID]...)
// 	return copy
// }

// func (h *History) Add(userID string, oldBalance, newBalance, amount float64) {
// 	h.mutex.Lock()
// 	defer h.mutex.Unlock()

// 	if amount < 0 {
// 		h.data[userID] = append(h.data[userID], fmt.Sprintf("Снято с баланса %v рублей. Изменение баланса: %v -> %v", -amount, oldBalance, newBalance))
// 	} else {
// 		h.data[userID] = append(h.data[userID], fmt.Sprintf("Зачислено на баланс %v рублей. Изменение баланса: %v -> %v", amount, oldBalance, newBalance))
// 	}
// }

// func (h *History) AddTransfer(fromID, toID string, oldBalanceFrom, newBalanceFrom, oldBalanceTo, newBalanceTo, amount float64) {
// 	h.mutex.Lock()
// 	defer h.mutex.Unlock()

// 	h.data[fromID] = append(h.data[fromID], fmt.Sprintf("Снято с баланса %v рублей. Изменение баланса: %v -> %v", -amount, oldBalanceFrom, newBalanceFrom))
// 	h.data[toID] = append(h.data[toID], fmt.Sprintf("Зачислено на баланс %v рублей. Изменение баланса: %v -> %v", -amount, oldBalanceTo, newBalanceTo))
// }

// todo
// почитал задание, понял что не то делаю
// надо переосмыслить идею
