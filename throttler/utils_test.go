package throttler

import (
	"reflect"
	"testing"
)

func TestConvertSliceToMap(t *testing.T) {
	testMethods := []string{"GET", "POST", "PUT", "DELETE"}
	testMapMethods := map[string]struct{}{"DELETE": {}, "GET": {}, "POST": {}, "PUT": {}}
	mapMethods := convertSliceToMap(testMethods)
	for k, v := range mapMethods {
		if tv, ok := testMapMethods[k]; ok {
			if tv != v {
				t.Error(
					"key", k,
					"value", v,
					"message", "values are not equal",
				)
			}
		} else {
			t.Error(
				"key", k,
				"value", v,
				"message", "not key",
			)
		}
	}
}

func TestReplace(t *testing.T) {
	testS := []string{"/servers/*/status"}
	tesR := []string{"/servers/.+/status"}
	replace(testS)
	if !reflect.DeepEqual(testS, tesR) {
		t.Error(
			"message", "the values are incorrect",
		)
	}
}

func TestPattern(t *testing.T) {
	url := "http://apidomain.com/network/routes"
	testGR := []string{"/servers/.+/status"}
	testBR := []string{"/apidomain.com/.+/routes"}
	if !checkPattern(url, testBR) {
		t.Error(
			"message", "TestPattern not working",
		)
	}
	if checkPattern(url, testGR) {
		t.Error(
			"message", "TestPattern not working",
		)
	}
}
