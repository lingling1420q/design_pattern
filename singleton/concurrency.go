package main

//
//import "sync"
//
//var m *monitor
//
//var mu sync.Mutex
//
//func GetMonitor(name string) *monitor{
//
//	if m.name != "" {
//		return m
//	}
//
//	mu.Lock()
//	if m.name != "" {
//		return m
//	}
//
//	m = &monitor{
//		name : name,
//	}
//	mu.Unlock()
//	return m
//}
//
//var once sync.Once
//
//func GetMonitor(name string) *monitor{
//	once.Do(func() {
//		m = &monitor{
//			name : name,
//		}
//	})
//	return m
//}

//
