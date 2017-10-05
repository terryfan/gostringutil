package stringutil

import (
	"fmt"
	"os"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abcdefghijklmnopqrstuvwxyz")
	}
}

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func subTest1(t *testing.T) {

}

func subTest2(t *testing.T) {

}

func subTest3(t *testing.T) {

}

func TestGroupTest(t *testing.T) {
	t.Run("group", func(t *testing.T) {
		t.Run("Test1", subTest1)
		t.Run("Test2", subTest2)
		t.Run("Test3", subTest3)
	})
}

func TestMain(m *testing.M) {
	fmt.Println("start")
	defer fmt.Println("end...............")
	os.Exit(m.Run())
}
