package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Test deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Test Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		if err != nil {
			t.Fatal("Unexpected Error")
		}
	})

	t.Run("Test Withdraw Witb unsufficient funds", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(200)

		got := wallet.Balance()
		want := Bitcoin(20)

		if err == nil || err.Error() != ErrInsufficientFunds.Error() {
			t.Fatal("Wanted an error dint get one")
		}

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})
}
