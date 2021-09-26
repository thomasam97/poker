package main

type TestResp struct {
	MSG string `json:"msg"`
}

func main() {
	srv := NewServer()
	srv.Run()
}
