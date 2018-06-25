package coordinator_test

import (
	"testing"

	. "github.com/evanisnor/coordinator"
)

func TestRunOne(t *testing.T) {
	executed := false
	Run(func() {
		executed = true
	})

	Wait()
	if !executed {
		t.Fail()
	}
}
func TestRunTwo(t *testing.T) {
	executed1 := false
	executed2 := false
	Run(func() {
		executed1 = true
	})
	Run(func() {
		executed2 = true
	})

	Wait()
	if !executed1 || !executed2 {
		t.Fail()
	}
}
