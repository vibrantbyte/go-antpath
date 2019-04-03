/**
 * Created by GoLand.
 * Brief: matcher fields read/write
 * User: vibrant
 * Date: 2019/04/03
 * Time: 11:20
 */
package extend

import "sync"

//ClearSyncMap
func ClearSyncMap(m *sync.Map) {
	if m != nil {
		m.Range(func(key, value interface{}) bool {
			m.Delete(key)
			return true
		})
	}
}
