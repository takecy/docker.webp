package img

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"bytes"
	"encoding/json"

	"github.com/gin-gonic/gin"
	. "github.com/smartystreets/goconvey/convey"
)

func engine() *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.TestMode)
	New(r)
	return r
}

func TestToWebp(t *testing.T) {
	r := engine()

	Convey("Given id", t, func() {
		ID := "id_001"

		Convey("When call GET", func() {
			req, _ := http.NewRequest("GET", fmt.Sprintf("/%s/webp", ID), nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			Convey("Then 404", func() {
				So(w.Code, ShouldEqual, 404)
			})
		})

		Convey("When call POST nil body", func() {
			req, _ := http.NewRequest("POST", fmt.Sprintf("/%s/webp", ID), nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			Convey("Then 400", func() {
				So(w.Code, ShouldEqual, 400)
			})
		})

		Convey("When call POST not enough body", func() {
			u := &Upload{}
			b, _ := json.Marshal(u)
			req, _ := http.NewRequest("POST", fmt.Sprintf("/%s/webp", ID), bytes.NewReader(b))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			Convey("Then 400", func() {
				So(w.Code, ShouldEqual, 400)
			})
		})
	})
}
