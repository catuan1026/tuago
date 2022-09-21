package tuago

type KeyAble interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint | string
}

type NumberAble interface {
	int8 | uint8 | int16 | uint16 | int32 | uint32 | int64 | uint64 | int | uint | float32 | float64
}

type StringAble interface {
	ToString() string
}

type Result struct {
	ErrCode int    `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
	Data    any    `json:"data"`
}

type CleanAble interface {
	Clean() error
}

type TableInf interface {
	TableName() string
}

type InitAble interface {
	Init() error
}
