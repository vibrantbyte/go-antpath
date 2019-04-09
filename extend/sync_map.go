/**
 * Created by GoLand.
 * Brief: matcher fields read/write
 * User: vibrant
 * Date: 2019/04/09
 * Time: 17:20
 */
package extend

import "sync"

type SyncMap struct {
	//data
	sync.Map
	sync.RWMutex
	//len
	len int64
}

//MyStore
func (m *SyncMap) MyStore(key, value interface{}) {
	m.Store(key,value)
	m.Lock()
	defer m.Unlock()
	m.len ++
}

//MyLoad
func (m *SyncMap) MyLoad(key interface{}) (value interface{}, ok bool){
	return m.Load(key)
}

//MyLoadOrStore
func (m *SyncMap) MyLoadOrStore(key, value interface{}) (actual interface{}, loaded bool){
	return m.LoadOrStore(key,value)
}

//MyDelete
func (m *SyncMap) MyDelete(key interface{}){
	m.Delete(key)
	m.Lock()
	defer m.Unlock()
	m.len --
}

//MyRange
func (m *SyncMap) MyRange(f func(key, value interface{}) bool){
	m.Range(f)
}

//MyLen
func (m *SyncMap) MyLen() int64 {
	m.RLock()
	defer m.RUnlock()
	return m.len
}