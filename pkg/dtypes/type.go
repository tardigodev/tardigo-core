package dtypes

type DataType interface {
	GetName() string
	Validate() error
	Convert(any) (any, error)
}

type Column struct {
	Name string
	Type DataType
}

type Schema []Column
