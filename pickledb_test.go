package pickledb

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDB(t *testing.T) {
	Convey("test filedb", t, func() {
		Convey("test not exist db", func() {
			db := New("./test_data/test.db", "")
			So(db, ShouldNotEqual, nil)

			Convey("test set key & get value", func() {
				db.Set("test1", "string")
				db.Set("test2", 4567)
				db.Set("test3", true)

				test1 := db.Get("test1")
				test2 := db.Get("test2")
				test3 := db.Get("test3")

				So(test1, ShouldEqual, "string")
				So(test2.(int), ShouldEqual, 4567)
				So(test3.(bool), ShouldEqual, true)
			})

		})

		Convey("test exist db", func() {
			db := New("./test_data/exist.db", "")
			So(db, ShouldNotEqual, nil)
			db.Load()

			Convey("test get value", func() {
				test1 := db.Get("test1")
				test2 := db.Get("test2")
				test3 := db.Get("test3")

				So(test1.(string), ShouldEqual, "string")
				So(test2.(float64), ShouldEqual, 4567)
				So(test3.(bool), ShouldEqual, true)

				Convey("test remove & set value", func() {
					db.Remove("test1")
					test1 := db.Get("test1")
					So(test1, ShouldEqual, nil)

					db.Set("test1", "string")
					test1 = db.Get("test1")
					So(test1.(string), ShouldEqual, "string")
				})
			})
		})
	})
}
