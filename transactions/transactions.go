package transactions

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	// loop through every transactions if the name matches
	// will add or subtract depending on the from or to value
	var balance float64
	for _, t := range transactions {

		if t.From == name {
			balance -= t.Sum
		}

		if t.To == name {
			balance += t.Sum
		}
	}

	return balance
}
