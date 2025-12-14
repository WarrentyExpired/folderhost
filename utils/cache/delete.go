package cache

func (c *Cache[KeyType, DataType]) Delete(key KeyType) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	delete(c.Items, key)
}

func (c *Cache[KeyType, DataType]) DeleteDirCacheItemsByPath(path string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	for key := range c.Items {
		if dirKey, ok := any(key).(DirectoryCacheKey); ok {
			if dirKey.Path == path {
				delete(c.Items, key)
			}
		}
	}
}

func (c *Cache[KeyType, DataType]) DeleteDirCacheItemsByPathWithoutThisScope(path string, scope string) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	for key := range c.Items {
		if dirKey, ok := any(key).(DirectoryCacheKey); ok {
			if dirKey.Path == path && dirKey.Scope != scope {
				delete(c.Items, key)
			}
		}
	}
}
