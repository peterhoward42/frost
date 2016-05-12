package contract

type FrostType int

const (
	FrostInt FrostType = iota
	FrostFloat
	FrostBool
	FrostString
)

type HasFrostType interface {
	GetFrostType() FrostType
}
