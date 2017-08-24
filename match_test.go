package gohelpconvey

import (
	"testing"

	"github.com/Cappta/gofixture"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMatch(t *testing.T) {
	Convey("Given an string value", t, func() {
		value := "Go help convey, birl!"
		Convey("Given an matching regex string", func() {
			regex := ".+"
			Convey("When checking if value match regex", func() {
				result := ShouldMatch(value, regex)
				Convey("Then result should equal an empty string", func() {
					So(result, ShouldEqual, "")
				})
			})
			Convey("When checking if value does not match regex", func() {
				result := ShouldNotMatch(value, regex)
				Convey("Then result should not equal an empty string", func() {
					So(result, ShouldNotEqual, "")
				})
			})
		})
		Convey("Given an unmatching regex string", func() {
			regex := "\\d+"
			Convey("When checking if value match regex", func() {
				result := ShouldMatch(value, regex)
				Convey("Then result should not equal an empty string", func() {
					So(result, ShouldNotEqual, "")
				})
			})
			Convey("When checking if value does not match regex", func() {
				result := ShouldNotMatch(value, regex)
				Convey("Then result should equal an empty string", func() {
					So(result, ShouldEqual, "")
				})
			})
		})
		Convey("Given an invalid regex string", func() {
			regex := string(gofixture.Bytes(256))
			Convey("When checking if value match regex", func() {
				result := ShouldMatch(value, regex)
				Convey("Then result should not equal an empty string", func() {
					So(result, ShouldNotEqual, "")
				})
			})
			Convey("When checking if value does not match regex", func() {
				result := ShouldNotMatch(value, regex)
				Convey("Then result should not equal an empty string", func() {
					So(result, ShouldNotEqual, "")
				})
			})
		})
		Convey("Given an integer instead of a regex", func() {
			regex := 1337
			Convey("When checking if value match regex", func() {
				result := ShouldMatch(value, regex)
				Convey("Then result should not equal an empty string", func() {
					So(result, ShouldNotEqual, "")
				})
			})
			Convey("When checking if value does not match regex", func() {
				result := ShouldNotMatch(value, regex)
				Convey("Then result should not equal an empty string", func() {
					So(result, ShouldNotEqual, "")
				})
			})
		})
		Convey("When checking if value match without a regex", func() {
			result := ShouldMatch(value)
			Convey("Then result should not equal an empty string", func() {
				So(result, ShouldNotEqual, "")
			})
		})
		Convey("When checking if value does not match without a regex", func() {
			result := ShouldNotMatch(value)
			Convey("Then result should not equal an empty string", func() {
				So(result, ShouldNotEqual, "")
			})
		})
	})
	Convey("Given an integer value", t, func() {
		value := 1337
		Convey("Given an matching regex string", func() {
			regex := ".+"
			Convey("When checking if value match regex", func() {
				result := ShouldMatch(value, regex)
				Convey("Then result should not equal an empty string", func() {
					So(result, ShouldNotEqual, "")
				})
			})
			Convey("When checking if value does not match regex", func() {
				result := ShouldNotMatch(value, regex)
				Convey("Then result should not equal an empty string", func() {
					So(result, ShouldNotEqual, "")
				})
			})
		})
	})
}
