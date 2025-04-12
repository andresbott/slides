package goleak

import (
	"fmt"
	"go.uber.org/goleak"
	"os"
	"testing"
)

// BEGIN GOLEAKTEST OMIT
func TestA(t *testing.T) {
	defer goleak.VerifyNone(t) // HL
	// test logic here.
}

// END GOLEAKTEST OMIT

// BEGIN GOLEAKTESTADV OMIT
func TestMain(m *testing.M) {
	code := m.Run()

	//check for routine leaks
	opts := []goleak.Option{
		// TODO verify
		goleak.IgnoreAnyFunction("database/sql.(*DB).connectionOpener"),
	}
	if err := goleak.Find(opts...); err != nil {
		fmt.Printf("found routine leak: %v\n", err)
		os.Exit(1)
	}
	os.Exit(code)
}

// END GOLEAKTESTADV OMIT
