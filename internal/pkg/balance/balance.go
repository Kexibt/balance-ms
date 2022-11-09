package balance

// Balances структура, отвечающая за взаимодействие с балансом пользователей
type Balances struct {
	db Database
}

// GetByUID возвращает актуальный баланс пользователя
func (b *Balances) GetByUID(userID string) float64 {
	amount, _ := b.db.Get(userID)
	if amount == 0 {
		b.db.Create(userID, 0)
	}

	return amount
}

// Change меняет баланс пользователя
func (b *Balances) Change(userID string, amount float64) (float64, error) {
	old_amount, _ := b.db.Get(userID)
	if old_amount == 0 {
		b.db.Create(userID, 0)
	}

	new_balance := old_amount + amount
	if new_balance < 0 {
		return 0, errorNotEnoughMoney(userID)
	}

	err := b.db.Change(userID, new_balance)
	if err != nil {
		return 0, err
	}
	return new_balance, nil
}

// Transfer переводит деньги с одного баланса на другой
func (b *Balances) Transfer(fromUserID, toUserID string, amount float64) (float64, float64, error) {
	if amount < 0 {
		return 0, 0, errorNegativeAmount("перевода")
	}

	usr1_amount, _ := b.db.Get(fromUserID)
	if usr1_amount == 0 {
		b.db.Create(fromUserID, 0)
	}

	usr2_amount, _ := b.db.Get(toUserID)
	if usr2_amount == 0 {
		b.db.Create(toUserID, 0)
	}

	if usr1_amount-amount < 0 {
		return 0, 0, errorNotEnoughMoney(fromUserID)
	}
	if usr2_amount+amount < 0 {
		return 0, 0, errorNotEnoughMoney(toUserID)
	}

	if fromUserID < toUserID {
		err := b.db.Change(fromUserID, usr1_amount-amount)
		if err != nil {
			return 0, 0, err
		}

		err = b.db.Change(toUserID, usr2_amount+amount)
		if err != nil {
			return 0, 0, err
		}
	} else if fromUserID > toUserID {
		err := b.db.Change(toUserID, usr2_amount+amount)
		if err != nil {
			return 0, 0, err
		}

		err = b.db.Change(fromUserID, usr1_amount-amount)
		if err != nil {
			return 0, 0, err
		}
	}

	return usr1_amount - amount, usr2_amount + amount, nil
}

// NewBalances конструктор для Balances
func NewBalances(db Database) *Balances {
	return &Balances{
		db: db,
	}
}
