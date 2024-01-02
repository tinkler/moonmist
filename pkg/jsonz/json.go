package jsonz

type Json interface {
	Unmarshal(data []byte, v any) error
	Marshal(v any) ([]byte, error)
}
