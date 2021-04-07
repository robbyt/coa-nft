package v1

// WalletType is an Enum of various crypto wallets
type WalletType int

const (
	ETH WalletType = iota
)

// this must be updated when adding more Enum values above
func (s WalletType) String() string {
	return [...]string{
		"ETH",
	}[s]
}

/* Wallet wraps the WalletType Enum

itle - The name of this wallet
Value - The actual crypto data associated with this wallet
Type - An Enum for various supported wallet types
*/
type Wallet struct {
	ID    int
	Title string
	Value string
	Type  WalletType
}
