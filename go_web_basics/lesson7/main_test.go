package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	A   string
	B   string
	Res string
}

func TestStringConcat(t *testing.T) {
	testCases := []TestCase{
		{A: "HELLO", B: "WORLD", Res: "HELLO WORLD"},
		{A: "1", B: "2", Res: "1 2"},
		{A: "Big", B: "Data", Res: "Big Data"},
	}
	for _, tcase := range testCases {
		res := StringConcat(tcase.A, tcase.B)
		if res != tcase.Res {
			t.Errorf("got %s, excepter %s", res, tcase.Res)
		}
	}
}

func BenchmarkStringConcat(b *testing.B) {
	arr := make([]int, 0, 8)
	for i := 0; i < b.N; i++ {
		// _ = StringConcat("A", "B")
		arr = append(arr, 32)
	}
}

func TestHandleEcho(t *testing.T) {
	testCases := []string{
		"Hello world",
		"123",
	}

	for _, tcase := range testCases {

		reader := bytes.NewReader([]byte(tcase))
		req, _ := http.NewRequest("POST", "/echo", reader)
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HandleEcho)
		handler.ServeHTTP(rr, req)
		if resp := rr.Body.String(); resp != tcase {
			t.Errorf("got %s, excpected %s", tcase, resp)
		}
	}
}
