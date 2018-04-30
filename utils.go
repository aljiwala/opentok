package opentok

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
)

// Decode should decode the response data into provided interface value.
func Decode(r *http.Request, data interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		return err
	}
	return nil
}

// IsZero determines whether or not a variable/field/whatever is of it's type's
// zero value pass reflect.ValueOf(x).
// http://stackoverflow.com/a/23555352/3183170
func IsZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Func, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Array:
		z := true
		for i := 0; i < v.Len(); i++ {
			z = z && IsZero(v.Index(i))
		}
		return z
	case reflect.Struct:
		z := true
		for i := 0; i < v.NumField(); i++ {
			z = z && IsZero(v.Field(i))
		}
		return z
	}

	// Compare other types directly:
	z := reflect.Zero(v.Type())
	return v.Interface() == z.Interface()
}

// CheckResponse ...
func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; http.StatusOK <= c && c <= http.StatusIMUsed {
		return nil
	}

	exception := new(Exception)
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &exception)
	}

	return exception
}

func structToURLValues(i interface{}) url.Values {
	v := url.Values{}
	m := structToMapString(i)
	for k, s := range m {
		switch {
		case len(s) == 1:
			v.Set(k, s[0])
		case len(s) > 1:
			for i := range s {
				v.Add(k, s[i])
			}
		}
	}

	return v
}

// structToMapString converts struct as map string
func structToMapString(i interface{}) map[string][]string {
	ms := map[string][]string{}
	iv := reflect.ValueOf(i).Elem()
	tp := iv.Type()

	for i := 0; i < iv.NumField(); i++ {
		k := tp.Field(i).Name
		f := iv.Field(i)
		ms[k] = valueToString(f)
	}

	return ms
}

// valueToString converts supported type of f as slice string
func valueToString(f reflect.Value) []string {
	var v []string

	switch reflect.TypeOf(f.Interface()).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v = []string{strconv.FormatInt(f.Int(), 10)}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v = []string{strconv.FormatUint(f.Uint(), 10)}
	case reflect.Float32:
		v = []string{strconv.FormatFloat(f.Float(), 'f', 4, 32)}
	case reflect.Float64:
		v = []string{strconv.FormatFloat(f.Float(), 'f', 4, 64)}
	case reflect.Bool:
		v = []string{strconv.FormatBool(f.Bool())}
	case reflect.Slice:
		for i := 0; i < f.Len(); i++ {
			if s := valueToString(f.Index(i)); len(s) == 1 {
				v = append(v, s[0])
			}
		}
	case reflect.String:
		v = []string{f.String()}
	}

	return v
}
