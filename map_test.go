package collection

import (
	"log"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestMap_LoadOrNew(t *testing.T) {

	m := SafeMap{}

	f := func() any {
		return "xxx"
	}
	want := "xxx"

	for i := 0; i < 100; i++ {
		if got := m.LoadOrNew(i, f); !reflect.DeepEqual(got, want) {
			t.Errorf("LoadOrNew() = %v, want %v", got, want)
		}
	}

}

func TestSafeMap_Concurrent(t *testing.T) {
	m := SafeMap{}
	o := sync.Once{}

	slow := func() any {
		o.Do(func() {
			log.Println("-------")
			time.Sleep(1 * time.Second)
		})
		return "xxx"
	}

	want := 100
	wg := sync.WaitGroup{}

	for i := 0; i < want; i++ {
		wg.Add(1)
		go func() {
			v := m.LoadOrNew("127.0.0.1", slow)
			log.Println(">>>>>", v)
			wg.Done()
		}()
	}

	wg.Wait()
}
