package xx

import (
	"go/constant"
	"go/token"
	"reflect"
	"strconv"
)

// GeneralizeMatcher generalize a matcher
func GeneralizeMatcher[T any](m Matcher) F11[T, bool] {
	return GF11[T, bool](m)
}

// Matcher matcher accept 1 argument and returns 1 bool
type Matcher = F11[any, bool]

// Compare create a Matcher compares a value with a token
func Compare(v any, tok token.Token) Matcher {
	return func(v1 any) bool {
		return compare(v1, v, tok)
	}
}

// And returns a new matcher returns true if only all children returns true
func And(m ...Matcher) Matcher {
	return func(v1 any) bool {
		if len(m) == 0 {
			return false
		}
		for _, item := range m {
			if !item.Do(v1) {
				return false
			}
		}
		return true
	}
}

// Or returns a new matcher returns true if any of children returns true
func Or(m ...Matcher) Matcher {
	return func(v1 any) bool {
		for _, item := range m {
			if item.Do(v1) {
				return true
			}
		}
		return false
	}
}

// Not returns a new matcher returns true if none of children returns true
func Not(m ...Matcher) Matcher {
	return func(v1 any) bool {
		for _, item := range m {
			if item.Do(v1) {
				return false
			}
		}
		return true
	}
}

// Eq returns a matcher that returns true if incoming value is same
func Eq(v any) Matcher {
	return Compare(v, token.EQL)
}

// Neq returns a func that returns true if incoming value is not same
func Neq(v any) Matcher {
	return Not(Eq(v))
}

// Gt returns a func that returns true if incoming values is greater
func Gt(v any) Matcher {
	return Compare(v, token.GTR)
}

// Lt returns a func that returns true if incoming values is less
func Lt(v any) Matcher {
	return Compare(v, token.LSS)
}

// Gte returns a func that returns true if incoming values is greater or equal
func Gte(v any) Matcher {
	return Compare(v, token.GEQ)
}

// Lte returns a func that returns true if incoming values is less or equal
func Lte(v any) Matcher {
	return Compare(v, token.LEQ)
}

// compare compare two values with an operation token
// code copied from https://github.com/asdine/storm/blob/23213e9525985cdd3f0ca079c19b78db7e67b76f/q/compare.go
func compare(a, b interface{}, tok token.Token) bool {
	vala := reflect.ValueOf(a)
	valb := reflect.ValueOf(b)

	ak := vala.Kind()
	bk := valb.Kind()
	switch {
	// comparing nil values
	case (ak == reflect.Ptr || ak == reflect.Slice || ak == reflect.Interface || ak == reflect.Invalid) &&
		(bk == reflect.Ptr || ak == reflect.Slice || bk == reflect.Interface || bk == reflect.Invalid) &&
		(!vala.IsValid() || vala.IsNil()) && (!valb.IsValid() || valb.IsNil()):
		return true
	case ak >= reflect.Int && ak <= reflect.Int64:
		if bk >= reflect.Int && bk <= reflect.Int64 {
			return constant.Compare(constant.MakeInt64(vala.Int()), tok, constant.MakeInt64(valb.Int()))
		}

		if bk >= reflect.Uint && bk <= reflect.Uint64 {
			return constant.Compare(constant.MakeInt64(vala.Int()), tok, constant.MakeInt64(int64(valb.Uint())))
		}

		if bk == reflect.Float32 || bk == reflect.Float64 {
			return constant.Compare(constant.MakeFloat64(float64(vala.Int())), tok, constant.MakeFloat64(valb.Float()))
		}

		if bk == reflect.String {
			bla, err := strconv.ParseFloat(valb.String(), 64)
			if err != nil {
				return false
			}

			return constant.Compare(constant.MakeFloat64(float64(vala.Int())), tok, constant.MakeFloat64(bla))
		}
	case ak >= reflect.Uint && ak <= reflect.Uint64:
		if bk >= reflect.Uint && bk <= reflect.Uint64 {
			return constant.Compare(constant.MakeUint64(vala.Uint()), tok, constant.MakeUint64(valb.Uint()))
		}

		if bk >= reflect.Int && bk <= reflect.Int64 {
			return constant.Compare(constant.MakeUint64(vala.Uint()), tok, constant.MakeUint64(uint64(valb.Int())))
		}

		if bk == reflect.Float32 || bk == reflect.Float64 {
			return constant.Compare(constant.MakeFloat64(float64(vala.Uint())), tok, constant.MakeFloat64(valb.Float()))
		}

		if bk == reflect.String {
			bla, err := strconv.ParseFloat(valb.String(), 64)
			if err != nil {
				return false
			}

			return constant.Compare(constant.MakeFloat64(float64(vala.Uint())), tok, constant.MakeFloat64(bla))
		}
	case ak == reflect.Float32 || ak == reflect.Float64:
		if bk == reflect.Float32 || bk == reflect.Float64 {
			return constant.Compare(constant.MakeFloat64(vala.Float()), tok, constant.MakeFloat64(valb.Float()))
		}

		if bk >= reflect.Int && bk <= reflect.Int64 {
			return constant.Compare(constant.MakeFloat64(vala.Float()), tok, constant.MakeFloat64(float64(valb.Int())))
		}

		if bk >= reflect.Uint && bk <= reflect.Uint64 {
			return constant.Compare(constant.MakeFloat64(vala.Float()), tok, constant.MakeFloat64(float64(valb.Uint())))
		}

		if bk == reflect.String {
			bla, err := strconv.ParseFloat(valb.String(), 64)
			if err != nil {
				return false
			}

			return constant.Compare(constant.MakeFloat64(vala.Float()), tok, constant.MakeFloat64(bla))
		}
	case ak == reflect.String:
		if bk == reflect.String {
			return constant.Compare(constant.MakeString(vala.String()), tok, constant.MakeString(valb.String()))
		}
	}

	typea, typeb := reflect.TypeOf(a), reflect.TypeOf(b)

	if typea != nil && (typea.String() == "time.Time" || typea.String() == "*time.Time") &&
		typeb != nil && (typeb.String() == "time.Time" || typeb.String() == "*time.Time") {

		if typea.String() == "*time.Time" && vala.IsNil() {
			return true
		}

		if typeb.String() == "*time.Time" {
			if valb.IsNil() {
				return true
			}
			valb = valb.Elem()
		}

		var x, y int64
		x = 1
		if vala.MethodByName("Equal").Call([]reflect.Value{valb})[0].Bool() {
			y = 1
		} else if vala.MethodByName("Before").Call([]reflect.Value{valb})[0].Bool() {
			y = 2
		}
		return constant.Compare(constant.MakeInt64(x), tok, constant.MakeInt64(y))
	}

	if tok == token.EQL {
		return reflect.DeepEqual(a, b)
	}

	return false
}
