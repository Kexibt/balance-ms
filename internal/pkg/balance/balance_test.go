package balance

// import "testing"

// func TestGet(t *testing.T) {
// 	b := NewBalances()
// 	id := "user1_test"

// 	res := b.GetByUID(id)
// 	if res != 0 {
// 		t.Errorf("Expected 0, but got %v", res)
// 	}

// 	expected := float64(200)
// 	b.data[id] = expected

// 	res = b.GetByUID(id)
// 	if res != expected {
// 		t.Errorf("Expected: %v, but got: %v", expected, res)
// 	}
// }

// func TestNegativeChange(t *testing.T) {
// 	b := NewBalances()
// 	id := "user1_test"

// 	res, err := b.Change(id, -100)
// 	if err == nil {
// 		t.Error("Expected error, but got nil")
// 	}
// 	if res != 0 {
// 		t.Errorf("Expected unchanged amount, but got %v", res)
// 	}
// 	if res = b.GetByUID(id); res != 0 {
// 		t.Errorf("Expected unchanged amount, but got %v", res)
// 	}

// 	expected := errorNotEnoughMoney(id)
// 	if err.Error() != expected.Error() {
// 		t.Errorf("Expected %v, but got %v", expected, err)
// 	}
// }

// func TestPositiveChange(t *testing.T) {
// 	b := NewBalances()
// 	id := "user1_test"
// 	expected := float64(100)

// 	res, err := b.Change(id, expected)
// 	if err != nil {
// 		t.Errorf("Expected ok status, but got an error: %v", err)
// 	}
// 	if res != expected {
// 		t.Errorf("Expected: %v, but got: %v", expected, res)
// 	}
// }

// func TestPositiveNegativeChange(t *testing.T) {
// 	b := NewBalances()
// 	id := "user1_test"
// 	add := float64(100)
// 	sub := float64(-50)
// 	expected := float64(50)

// 	res, err := b.Change(id, add)
// 	if err != nil {
// 		t.Errorf("Expected ok status, but got an error: %v", err)
// 	}
// 	if res != add {
// 		t.Errorf("Expected: %v, but got: %v", add, res)
// 	}

// 	res, err = b.Change(id, sub)
// 	if err != nil {
// 		t.Errorf("Expected ok status, but got an error: %v", err)
// 	}
// 	if res != expected {
// 		t.Errorf("Expected: %v, but got: %v", expected, res)
// 	}
// }

// func TestPositiveTransfer(t *testing.T) {
// 	b := NewBalances()
// 	id1 := "user1_test"
// 	id2 := "user2_test"
// 	balance1 := float64(1000)
// 	balance2 := float64(1500)
// 	transfer := float64(1000)
// 	expected1 := float64(0)
// 	expected2 := float64(2500)

// 	b.data[id1] = balance1
// 	b.data[id2] = balance2

// 	res1, res2, err := b.Transfer(id1, id2, transfer)
// 	if err != nil {
// 		t.Errorf("Expected ok status, but got an error: %s", err)
// 	}
// 	if res1 != expected1 || res2 != expected2 {
// 		t.Errorf("Expected:\nres1: %v, res2: %v, but got:\nres1: %v, res2: %v", expected1, expected2, res1, res2)
// 	}
// }

// func TestNegativeTransfer(t *testing.T) {
// 	b := NewBalances()
// 	id1 := "user1_test"
// 	id2 := "user2_test"
// 	balance1 := float64(900)
// 	balance2 := float64(1500)
// 	transfer := float64(1000)
// 	expected1 := float64(900)
// 	expected2 := float64(1500)

// 	b.data[id1] = balance1
// 	b.data[id2] = balance2

// 	_, _, err := b.Transfer(id1, id2, transfer)
// 	if err == nil {
// 		t.Error("Expected err, but got nil")
// 	}
// 	if b.data[id1] != balance1 || b.data[id2] != balance2 {
// 		t.Errorf("Expected:\nres1: %v, res2: %v, but got:\nres1: %v, res2: %v", expected1, expected2, b.data[id1], b.data[id2])
// 	}
// }
