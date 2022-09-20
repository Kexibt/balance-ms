package balance

import (
	"sync"
)

type Balances struct {
	data  map[string]uint
	mutex sync.RWMutex
}

func (b *Balances) GetByUID(userID string) uint {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return b.data[userID]
}

func (b *Balances) Change(userID string, amount int) (uint, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	if amount < 0 {
		if b.data[userID] < uint(-amount) {
			return 0, errorNotEnoughMoney(userID)
		}
		b.data[userID] -= uint(-amount)
		// todo database + history

		return b.data[userID], nil
	}

	b.data[userID] += uint(amount)
	// todo database + history

	return b.data[userID], nil
}

func (b *Balances) Transfer(fromUserID, toUserID string, amount int) (uint, uint, error) {
	if amount < 0 {
		return 0, 0, errorNegativeAmount("перевода")
	}

	b.mutex.Lock()
	defer b.mutex.Unlock()

	if b.data[fromUserID] < uint(amount) {
		return 0, 0, errorNotEnoughMoney(fromUserID)
	}

	b.data[fromUserID] -= uint(amount)
	b.data[toUserID] += uint(amount)
	// todo database + history

	return b.data[fromUserID], b.data[toUserID], nil
}

func NewBalances() *Balances {
	return &Balances{
		data: make(map[string]uint, 100),
	}
}
