package cache

// DataHandler : all data stored in lru cache must implement DataHandler interface.
type DataHandler interface {
	// Free : function called when data is removed from the cache
	Free()
	// Get : should return data cached
	GetData() interface{}
}
