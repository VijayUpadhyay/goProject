package main

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

// ============================================================================
// Example 1: regexChk.go
// ============================================================================

func example1_regexchk() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	var reSlash = regexp.MustCompile(`surveys\/sections\/([a-zA-Z_]+)\/fields`)
	match = reSlash.MatchString("surveys/sections/cogs/fields")
	fmt.Println(match)
	//match, _ = regexp.MatchString("/surveys/sections/[a-zA-Z]/fields", "/surveys/sections/balancesheet/fields")

	reSlash = regexp.MustCompile(`users\/([a-zA-Z0-9-]+)\/roles`)
	match = reSlash.MatchString("users/fbd3036f-0f1c-4e98-b71c-d4cd61213f90/roles")
	fmt.Println(match)
}

// ============================================================================
// Example 2: syncMap.go
// ============================================================================

type SyncMap struct {
	lock sync.RWMutex
	KV   map[string]interface{}
}

func (m *SyncMap) Store(k string, v interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.KV[k] = v
}

func (m *SyncMap) Load(k string) interface{} {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return m.KV[k]
}

func (m *SyncMap) Remove(k string) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.KV, k)
}

func (m *SyncMap) Size() int {
	m.lock.Lock()
	defer m.lock.Unlock()
	return len(m.KV)
}

// ============================================================================
// Example 3: utils.go
// ============================================================================

// ============================================================================
// Main Function - Run All Examples
// ============================================================================

func main() {
	fmt.Println("Running 1 example from BIUtils")
	fmt.Println(strings.Repeat("=", 60))

	fmt.Println("\n--- example1_regexchk ---")
	example1_regexchk()

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("All examples completed!")

	// Note: syncMap.go and utils.go contain utility types/functions without main()
	// SyncMap type is available for use in other code
}
