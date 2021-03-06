package host

import (
	"fmt"
	"testing"
)

func TestReadFromArgs(t *testing.T) {
	f := func(domains []string, ip string) {
		t.Helper()
		data := ReadFromArgs(domains, ip)
		if len(data.profiles["default"]) != len(domains) {
			t.Fatalf("bad count of records; got %d; want %d", len(data.profiles["default"]), len(domains))
		}
		for k, d := range domains {
			got := data.profiles["default"][k]
			want := fmt.Sprintf("%s %s", ip, d)
			if got != want {
				t.Fatalf("unexpected record; got %s; want %s", got, want)
			}
		}
	}
	f([]string{}, "")
	f([]string{"dom1.local"}, "")
	f([]string{"dom1.local"}, "127.0.0.1")
	f([]string{"dom3.local", "dom4.local"}, "localhost")
}
