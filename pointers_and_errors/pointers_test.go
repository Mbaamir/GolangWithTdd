package pointers_and_errors

import (
	"testing"
)

func checkBitcoinEquality(t testing.TB, got Bitcoin, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, errGot, errWant error) {
	t.Helper()
	if errGot == nil {
		t.Fatal("expected an error but didn't get one")
	}
	if errGot != errWant {
		t.Errorf("got %q, want %q", errGot, errWant)
	}
}

func TestWallet(t *testing.T) {

	t.Run(("deposit"), func(t *testing.T) {
		wallet := Wallet{}
		bitcoinToDespoit := Bitcoin(10)

		wallet.Deposit(bitcoinToDespoit)

		got := wallet.Balance()
		want := Bitcoin(10)

		checkBitcoinEquality(t, got, want)
	})

	t.Run(("withdraw"), func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(10)

		assertNoError(t, err)

		got := wallet.Balance()
		want := Bitcoin(10)
		checkBitcoinEquality(t, got, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, insuffucientBalanceErr)
	})

}
