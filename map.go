package collection

import "sync"

type SafeMap struct {
	sync.Map
}

func (m *SafeMap) LoadOrNew(key any, f func() any) any {
	if v, ok := m.Load(key); !ok {
		i := f()
		m.Store(key, i)
		return i
	} else {
		return v
	}
}
