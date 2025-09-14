package cache

// Options is a type for possible options to initialize the cache with.
type Options func(*Cache)

// OptionFullMessagebus sets the cache to use all messages from the messagebus.
//
// The default is, to use only messages with keys, that are already in the cache.
func OptionFullMessagebus(c *Cache) {
	c.fullMessagebus = true
}
