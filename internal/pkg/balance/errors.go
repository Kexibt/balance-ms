package balance

import (
	"fmt"
)

func errorNotEnoughMoney(userID string) error {
	return fmt.Errorf("у пользователя %s недостаточно средств", userID)
}

func errorNegativeAmount(typeTransaction string) error {
	return fmt.Errorf("сумма %s не может быть отрицательной", typeTransaction)
}
