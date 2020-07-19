package pointer_error

import (
	"errors"
	"fmt"
)

//如果有人想要重新定义这个错误，那么测试就会失败，这将是非常恼人的，而对于我们的测试来说，这里有太多的细节了。
// 我们并不关心具体的措辞是什么，只是在给定条件的情况下返回一些有意义的错误。
//在 Go 中，错误是值，因此我们可以将其重构为一个变量，并为其提供一个单一的事实来源。
var InsufficientFundsError = errors.New("insufficient balance")

//在 非常强调安全性的钱包 中，我们不想暴露自己的内部状态，而是通过方法来控制访问的权限。
type Wallet struct {
	balance Bitcoin
}

//我们曾说过我们正在制做一个比特币钱包，但到目前为止我们还没有提到它们。我们一直在使用 int，因为当用来计数时它是不错的类型！
//为此创建一个结构体似乎有点过头了。就 int 的表现来说已经很好了，但问题是它不具有描述性。
//Go 允许从现有的类型创建新的类型。
type Bitcoin int

//类型别名有一个有趣的特性，你还可以对它们声明 方法。当你希望在现有类型之上添加一些领域内特定的功能时，这将非常有用。
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//在 Go 中，当调用一个函数或方法时，参数会被复制,需要用指针类型，如果直接调用w Wallet的话 指向的是一个副本
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//取款
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//当你传值给函数或方法时，Go 会复制这些值。因此，如果你写的函数需要更改状态，你就需要用指针指向你想要更改的值
//Go 取值的副本在大多数时候是有效的，但是有时候你不希望你的系统只使用副本，在这种情况下你需要传递一个引用。
// 例如，非常庞大的数据或者你只想有一个实例（比如数据库连接池）