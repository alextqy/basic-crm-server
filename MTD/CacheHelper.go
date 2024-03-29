package mtd

import (
	"sync"
	"time"
)

type CacheHelper struct {
	mu            sync.RWMutex
	items         map[string]interface{}
	itemFunc      map[string]func(key string, value interface{})
	itemFieldFunc map[string]map[string]func(key string, field string, value interface{})
}

// return
func (c *CacheHelper) init() {
	c.items = make(map[string]interface{})
	c.itemFunc = make(map[string]func(key string, value interface{}))
	c.itemFieldFunc = make(map[string]map[string]func(key string, field string, value interface{}))
	Start(c.runJanitor)
}

// Reset the expiration
func (c *CacheHelper) Expire(key string, expiration time.Duration) (success bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.items[key]; !ok {
		return false
	}
	if expiration > 0 {
		AddTimer(key+":::"+"", expiration, map[string]string{
			"key":   key,
			"field": "",
		})
	}
	return true
}

// Delete the item
// remove the item or item field resource
func (c *CacheHelper) Del(keys ...string) (success int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var count int64
	for _, key := range keys {
		if _, ok := c.items[key]; ok {
			// success counter
			count++

			// remove item field resource
			{
				if fieldMap, ok := c.items[key].(map[string]interface{}); ok {
					for field := range fieldMap {
						c.removeJanitor(key, field)
					}
				}
			}

			//  remove item resource
			{
				delete(c.items, key)
				c.removeJanitor(key, "")
			}
		}
	}
	return count
}

// Create a goroutine to clean the key resource when timer expired
func (c *CacheHelper) runJanitor(data interface{}) {
	dataMap := data.(map[string]string)
	key := dataMap["key"]
	field := dataMap["field"]
	if field == "" {
		var callback func(key string, value interface{})
		// Lock
		c.mu.Lock()
		callbackValue := c.items[key]
		// if exits expiration callback , exec it then delete
		// hashmap key has no expiration callback
		if _, ok := c.itemFunc[key]; ok && c.itemFunc[key] != nil {
			callback = c.itemFunc[key]
		}
		// delete item callback
		delete(c.itemFunc, key)

		// remove item
		delete(c.items, key)

		// release Lock
		c.mu.Unlock()

		// exec expiration callback at last,avoid deadlock
		if callback != nil {
			go callback(key, callbackValue)
		}
	} else {
		var callback func(key string, field string, value interface{})
		c.mu.Lock()
		// get item value
		val := c.items[key]
		fieldMap, _ := val.(map[string]interface{})
		callbackValue := fieldMap[field]
		// exec the expiration callback and then delete the callback
		// if field func's length equals zero, delete the whole field's funcs
		if c.itemFieldFunc[key][field] != nil {
			callback = c.itemFieldFunc[key][field]
		}
		delete(c.itemFieldFunc[key], field)
		if len(c.itemFieldFunc[key]) == 0 {
			delete(c.itemFieldFunc, key)
		}
		// delete the item field
		delete(fieldMap, field)
		c.items[key] = fieldMap
		// if field length equals zero , delete whole item
		if len(fieldMap) == 0 {
			delete(c.items, key)
		}
		c.mu.Unlock()
		// exec expiration callback at last,avoid deadlock
		if callback != nil {
			go callback(key, field, callbackValue)
		}
	}
}

// removeJanitor
// remove the janitor of the item or item field
func (c *CacheHelper) removeJanitor(key string, field string) {
	StopTimer(key + ":::" + field)
	if field == "" {
		// remove the item func
		delete(c.itemFunc, key)
	} else {
		// remove item field expiration callback function
		if _, ok := c.itemFieldFunc[key]; ok {
			delete(c.itemFieldFunc[key], field)
			if len(c.itemFieldFunc[key]) == 0 {
				delete(c.itemFieldFunc, key)
			}
		}
	}
}

// HMSet Batch Set the hash field
func (c *CacheHelper) HMSet(key string, values ...interface{}) (success bool) {
	if len(values) == 0 || len(values)%2 != 0 {
		return false
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	kv := make(map[string]interface{})
	for len(values) > 0 {
		if field, ok := values[0].(string); !ok {
			return false
		} else {
			value := values[1]
			kv[field] = value
		}
		values = values[2:]
	}
	for field, value := range kv {
		if c.items[key] != nil {
			if fieldMap, ok := c.items[key].(map[string]interface{}); !ok {
				return false
			} else {
				fieldMap[field] = value
				c.items[key] = fieldMap
			}
		} else {
			c.items[key] = make(map[string]interface{})
			fieldMap := make(map[string]interface{})
			fieldMap[field] = value
			c.items[key] = fieldMap
		}
	}
	return true
}

// HSet Set the hash field with expiration and expiration callback function
func (c *CacheHelper) HSet(
	key string,
	field string,
	value interface{},
	expiration time.Duration,
	expirationFunc func(key string, field string, value interface{}),
) (success bool) {
	if key == "" || field == "" {
		return false
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	// init item field
	// if not a hashmap, return false
	if c.items[key] != nil {
		if fieldMap, ok := c.items[key].(map[string]interface{}); !ok {
			return false
		} else {
			fieldMap[field] = value
			c.items[key] = fieldMap
		}
	} else {
		c.items[key] = make(map[string]interface{})
		fieldMap := make(map[string]interface{})
		fieldMap[field] = value
		c.items[key] = fieldMap
	}

	// init item field expiration callback function
	if c.itemFieldFunc[key] == nil {
		c.itemFieldFunc[key] = make(map[string]func(key string, field string, value interface{}))
	}

	if expiration > 0 {
		c.itemFieldFunc[key][field] = expirationFunc
		AddTimer(key+":::"+field, expiration, map[string]string{
			"key":   key,
			"field": field,
		})
	}
	return true
}

// Get All Hash key value
func (c *CacheHelper) HGetAll(key string) (value map[string]interface{}, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// if not a map, return false
	if fieldMap, ok := c.items[key].(map[string]interface{}); ok {
		return fieldMap, true
	} else {
		return nil, false
	}
}

// Get hash field value, if found return true
func (c *CacheHelper) HGet(key string, field string) (value interface{}, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	// if not a map, return false
	if fieldMap, ok := c.items[key].(map[string]interface{}); ok {
		if fieldMap == nil {
			return nil, false
		}
		if v, ok := fieldMap[field]; ok {
			return v, true
		}
	}
	return nil, false
}

// Delete Hash field, if success return true
func (c *CacheHelper) HDel(key string, fields ...string) (count int64, success bool) {
	if key == "" {
		return 0, false
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	// if not a map, return false
	if fieldMap, ok := c.items[key].(map[string]interface{}); ok {
		for _, field := range fields {
			if field == "" {
				continue
			}
			if _, ok := fieldMap[field]; ok {
				count++
				delete(fieldMap, field)
				c.removeJanitor(key, field)
			}
		}
		// reset the item value
		if len(fieldMap) > 0 {
			c.items[key] = fieldMap
		} else {
			delete(c.items, key)
		}
		// return the success number
		if count > 0 {
			return count, true
		} else {
			return 0, false
		}
	} else {
		return 0, false
	}
}

// Reset the expiration of the hash field
func (c *CacheHelper) HExpire(key string, field string, expiration time.Duration) (success bool) {
	if key == "" || field == "" {
		return false
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	// if not a hash , return false
	if fieldMap, ok := c.items[key].(map[string]interface{}); ok {
		if _, ok := fieldMap[field]; !ok {
			return false
		}
	} else {
		return false
	}
	if expiration > 0 {
		AddTimer(key+":::"+field, expiration, map[string]string{
			"key":   key,
			"field": field,
		})
	}
	return true
}

// Set Key value with expiration and expiration callback function
func (c *CacheHelper) CacheSet(key string, value interface{}, expiration time.Duration, expirationFunc func(key string, value interface{})) (success bool) {
	if key == "" {
		return false
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
	if expiration > 0 {
		c.itemFunc[key] = expirationFunc
		AddTimer(key+":::"+"", expiration, map[string]string{
			"key":   key,
			"field": "",
		})
	}
	return true
}

// Get the vlaue of given key , if exist return true, or return false
func (c *CacheHelper) CacheGet(key string) (value interface{}, found bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if _, ok := c.items[key]; ok {
		return c.items[key], true
	}
	return nil, false
}
