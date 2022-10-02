package di

import "testing"

// clientImpl{}を模倣するもの
type testClientImpl struct{}

func (c testClientImpl) Get() int {
	return 2
}

func TestPrintSample2(t *testing.T) {
	sample := &Sample2{&testClientImpl{}}
	got := sample.doGet()
	want := 2
	if got != want {
		t.Errorf("エラー")
	}
}

func TestPrintSample3(t *testing.T) {
	// モック化
	client := testClientImpl{}
	got := client.Get()
	want := 2
	if got != want {
		t.Errorf("エラー")
	}
}
