package pointer_error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: 0}
		//存10
		wallet.Deposit(Bitcoin(10))

		//查询余额
		got := wallet.Balance()
		want := Bitcoin(10)

		assert.Equal(t, want, got)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		//取5
		err:= wallet.Withdraw(Bitcoin(5))
		assert.Empty(t,err)

		//查询余额
		got := wallet.Balance()
		want := Bitcoin(5)

		assert.Equal(t, want, got)
	})

	t.Run("Overspend", func(t *testing.T) {
		wallet := Wallet{balance: 5}
		//取100
		err := wallet.Withdraw(Bitcoin(100))

		assert.EqualError(t, err, InsufficientFundsError.Error())

		//查询余额
		got := wallet.Balance()
		want := Bitcoin(5)

		assert.Equal(t, want, got)
	})

}
