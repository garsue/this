package this

import (
	"fmt"
)

func init() {
	fmt.Println(`The Zen of Go, by David Cheney

Each package fulfils a single purpose
Handle errors explicitly
Return early rather than nesting deeply
Leave concurrency to the caller
Before you launch a goroutine, know when it will stop
Avoid package level state
Simplicity matters
Write tests to lock in the behaviour of your package’s API
If you think it’s slow, first prove it with a benchmark
Moderation is a virtue
Maintainability counts`)
}
