package regex

import (
	"bytes"
	"encoding/json"
	"sync"
	"testing"
)

type PublicPage struct {
	ID       int
	Name     string
	Url      string
	OwnerID  int
	ImageUrl string
	Tags     []string
	Desc     string
	Rules    []int
}

var (
	coolPage = PublicPage{
		ID:       1,
		Name:     "some page",
		Url:      "http://some-page.org",
		OwnerID:  123456,
		ImageUrl: "http://some-page.org/omg/123",
		Tags:     []string{"tag1", "tag2"},
		Desc:     "page to test sunc.Pool",
		Rules:    []int{1, 7, 11},
	}

	pages = []PublicPage{
		coolPage,
		coolPage,
		coolPage,
	}
)

func BenchmarkAllocNew(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := bytes.NewBuffer(make([]byte, 0, 64))
			_ = json.NewEncoder(data).Encode(pages)
		}
	})
}

var pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 64))
	},
}

func BenchmarkAllocPool(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			data := pool.Get().(*bytes.Buffer)
			_ = json.NewEncoder(data).Encode(pages)
			data.Reset()
			pool.Put(data)
		}
	})
}
