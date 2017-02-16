package chance_test

import (
	"regexp"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/victorquinn/chancego"
)

func TestChar(t *testing.T) {
	Convey("Existence", t, func() {
		c := chance.Char()
		So(c, ShouldNotBeNil)
	})

	Convey("Generates random character", t, func() {
		charRegex := regexp.MustCompile(`[a-z]`)
		charMap := map[string]int{}
		for i := 0; i < 2600; i++ {
			c := chance.Char()
			So(charRegex.MatchString(c), ShouldBeTrue)
			charMap[c] = charMap[c] + 1
		}
		
		Convey("Generated some of each character", func() {
			So(len(charMap), ShouldEqual, 26)
		})

		Convey("And is relatively evenly distributed", func() {
			for _, count := range charMap {
				So(count, ShouldAlmostEqual, 100, 30)
			}
		})
	})
}

func TestCharFromPool(t *testing.T) {
	Convey("Existence", t, func() {
		c := chance.CharFromPool("a")
		So(c, ShouldNotBeNil)
	})

	Convey("Should throw error if called with empty pool", t, func() {
		So(func() {
			c := chance.CharFromPool("")
			// This next line should never get hit but need to use c
			// or Go complains
			So(len(c), ShouldBeZeroValue)
		}, ShouldPanic)
	})

	Convey("Generates random character", t, func() {
		charRegex := regexp.MustCompile(`[a-e]`)
		charMap := map[string]int{}
		for i := 0; i < 500; i++ {
			c := chance.CharFromPool("abcde")
			So(charRegex.MatchString(c), ShouldBeTrue)
			charMap[c] = charMap[c] + 1
		}

		Convey("Generated some of each character", func() {
			So(len(charMap), ShouldEqual, 5)
		})

		Convey("And is relatively evenly distributed", func() {
			for _, count := range charMap {
				So(count, ShouldAlmostEqual, 100, 30)
			}
		})
	})
}
