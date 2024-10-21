package utils

type ResponseStructure struct {
	Data   interface{} `json:"data,omitempty"`
	Error  *ErrorInfo  `json:"error,omitempty"`
	Status bool        `json:"status"`
}

type ErrorInfo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// ColumnValue struct to hold column names and their corresponding values
type ColumnValue struct {
	ColName string
	Value   interface{}
}
