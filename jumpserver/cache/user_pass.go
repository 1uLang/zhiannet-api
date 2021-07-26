package cache

import "sync"

var userPassMaps sync.Map

func Set(k, v interface{}) {
	userPassMaps.Store(k, v)
}
func Get(k interface{}) (interface{}, bool) {
	return userPassMaps.Load(k)
}
func Delete(k interface{}) {
	userPassMaps.Delete(k)
}
