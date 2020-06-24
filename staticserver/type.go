package main

type Give struct {
	Err error
	Bytes  []byte
}

type Data struct {
	Data map[string]Give
}

func existe(e interface{}) {
}

func NewData() (D *Data) {
	//existe(Asset)
	return &Data {
		Data: make(map[string]Give),
	}
}
