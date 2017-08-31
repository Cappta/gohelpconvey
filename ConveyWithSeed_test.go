package gohelpconvey

import (
	"math/rand"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConveyWithSeed(t *testing.T) {
	ConveyWithSeed(1337, t, func() {
		Convey("Then next random Int63 should equal 5799089487994996006", func() {
			So(rand.Int63(), ShouldEqual, 5799089487994996006)
		})
		Convey("Then next random Int31 should equal 1350205738", func() {
			So(rand.Int31(), ShouldEqual, 1350205738)
		})
		Convey("Then next random Uint64 should equal 15022461524849771814", func() {
			So(rand.Uint64(), ShouldEqual, uint(15022461524849771814))
		})
		Convey("Then next random Uint32 should equal 2700411476", func() {
			So(rand.Uint32(), ShouldEqual, 2700411476)
		})
	})
}
