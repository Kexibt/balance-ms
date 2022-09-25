package balance

import (
	"sync"
)

// Balances структура, отвечающая за взаимодействие с балансом пользователей
type Balances struct {
	data  map[string]float64
	mutex sync.RWMutex
}

// GetByUID возвращает актуальный баланс пользователя
func (b *Balances) GetByUID(userID string) float64 {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.data[userID]
}

// Change меняет баланс пользователя
func (b *Balances) Change(userID string, amount float64) (float64, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.data[userID]+amount < 0 {
		return 0, errorNotEnoughMoney(userID)
	}
	b.data[userID] += amount

	return b.data[userID], nil
}

// Transfer переводит деньги с одного баланса на другой
func (b *Balances) Transfer(fromUserID, toUserID string, amount float64) (float64, float64, error) {
	if amount < 0 {
		return 0, 0, errorNegativeAmount("перевода")
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.data[fromUserID] < amount {
		return 0, 0, errorNotEnoughMoney(fromUserID)
	}

	b.data[fromUserID] -= amount
	b.data[toUserID] += amount

	return b.data[fromUserID], b.data[toUserID], nil
}

// NewBalances конструктор для Balances
func NewBalances() *Balances {
	return &Balances{
		data: make(map[string]float64, 100),
	}
}
