// Package builderpool is a small wrapper of sync.Pool which is
// specified one to strings.Builder.
// Go 1.10 or later is required to use this package because the
// strings.Buidler is introduced with Go 1.10.
package builderpool

import (
	"strings"
	"sync"
)

var global *BuilderPool

func init() {
	global = New()
}

// Get returns a strings.Builder from the global pool.
func Get() *strings.Builder {
	return global.Get()
}

// Release puts the given strings.Builder back into the global pool.
func Release(builder *strings.Builder) {
	global.Release(builder)
}

// BuilderPool is wrapper struct of sync.Pool for strings.Builder objects.
type BuilderPool struct {
	pool sync.Pool
}

// New returns a new BuilderPool instance.
func New() *BuilderPool {
	bp := BuilderPool{}
	bp.pool.New = allocBuilder
	return &bp
}

func allocBuilder() interface{} {
	return &strings.Builder{}
}

// Get returns a strings.Builder from the pool.
func (bp *BuilderPool) Get() *strings.Builder {
	return bp.pool.Get().(*strings.Builder)
}

// Release puts the given strings.Builder back into the pool
// after resetting the builder.
func (bp *BuilderPool) Release(builder *strings.Builder) {
	builder.Reset()
	bp.pool.Put(builder)
}
