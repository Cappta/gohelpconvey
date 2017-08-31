package gohelpconvey

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

// ConveyWithSeed starts a goconvey's convey given a seed, testing.T and an action
func ConveyWithSeed(seed int64, t *testing.T, action func()) {
	convey.Convey(fmt.Sprintf("Given the random seed %d", seed), t, func() {
		rand.Seed(seed)
		action()
	})
}
