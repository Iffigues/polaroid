package main

type Give struct {
	Err error
	Bytes  []byte
}

type Data struct {
	Data map[string]Give
}

func NewData() (D *Data) {
	return &Data {
		Data: make(map[string]Give),
	}
}
