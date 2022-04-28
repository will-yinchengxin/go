package failpointer

import "github.com/pingcap/failpoint"

func FailPointer() {
	// GO_FAILPOINTS="test/faiil=return(true)"
	failpoint.Inject("failpoint-name", func(val failpoint.Value) {
		failpoint.Return("unit-test", val)
	})
}
