// Package cache provides thread safe LRU cache
//
// LRU least recently used cache provided faster access to recently used data, bu utilizing the order of data stored in the cache
// This implemeation of cache uses hasmap and queue implemented using doubly linked list and make sure all the operation
// are thread safe.
//
// Example
// In Write ahead log segment object is kept inside cache.
//
// segment should implement DataHandler interface
// type segment struct{
// 		lg *LogFile
// 		idx *Index
// }
//
// func (s *segment) Free(){
// 		s.lg.Close()
// 		s.idx.Close()
// }
//
// func (s *segment) GetData() interface{}{
// 		return s;
// }
package cache
