package terraform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

// ResourceSlice has no documentation
type ResourceSlice []interface{}

// Marshal has no documentation
func Marshal(v interface{}, resource string, name string) ([]byte, error) {
	mapped := toMap("", v, "")
	m := OrderedMap{}
	switch typed := mapped.(type) {
	case *interface{}:
		m = (*typed).(OrderedMap)
	case OrderedMap:
		m = typed
	}
	m.Delete("id")
	var buf bytes.Buffer
	fmt.Fprintln(&buf, fmt.Sprintf(`resource %s %s {`, jenc(terraformat(resource)), jenc(terraformat(name))))
	for _, key := range m.OrderedKeys() {
		tfPrint(&buf, key.Value, m[key], indentStr)
	}
	fmt.Fprintln(&buf, "}")

	return buf.Bytes(), nil
}

// MarshalJSON has no documentation
func MarshalJSON(v interface{}, resource string, name string) ([]byte, error) {
	mapped := toMap("", v, "")
	m := OrderedMap{}
	switch typed := mapped.(type) {
	case *interface{}:
		m = (*typed).(OrderedMap)
	case OrderedMap:
		m = typed
	}
	m.Delete("id")
	outer := OrderedMap{
		OrderedMapKey{Value: "resource", Order: 0}: OrderedMap{
			OrderedMapKey{Value: terraformat(resource), Order: 0}: OrderedMap{
				OrderedMapKey{Value: terraformat(name), Order: 0}: m,
			},
		},
	}
	return json.MarshalIndent(outer, "", "\t")
}

func setm(m OrderedMap, key string, v interface{}, optional bool) {
	if v != nil {
		shouldSet := true
		if optional {
			switch reflect.TypeOf(v).Kind() {
			case reflect.String:
				shouldSet = len(fmt.Sprintf("%v", v)) > 0
			case reflect.Bool:
				shouldSet = v.(bool)
			case reflect.Map, reflect.Slice:
				shouldSet = (reflect.ValueOf(v).Len() > 0)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16:
				shouldSet = (fmt.Sprintf("%v", v) != "0")
			}
		}
		if shouldSet {
			m[m.NextKey(key)] = v
		}

	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func unOrderSlice(elems []interface{}) []interface{} {
	if (elems == nil) || (len(elems) == 0) {
		return elems
	}
	for i := 0; i < len(elems); i++ {
		switch typedElem := elems[i].(type) {
		case OrderedMap:
			elems[i] = unOrderMap(typedElem)
		case []interface{}:
			elems[i] = unOrderSlice(typedElem)
		default:
		}
	}
	return elems
}

func unOrderMap(m OrderedMap) map[string]interface{} {
	result := map[string]interface{}{}

	for k, v := range m {
		switch typedValue := v.(type) {
		case OrderedMap:
			result[k.Value] = unOrderMap(typedValue)
		case []interface{}:
			result[k.Value] = unOrderSlice(typedValue)
		default:
			result[k.Value] = v
		}
	}
	return result
}

func toMap(hint string, v interface{}, indent string) interface{} {
	if v == nil {
		return nil
	}
	switch reflect.TypeOf(v).Kind() {
	case reflect.String, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		return v
	case reflect.Ptr:
		rv := reflect.ValueOf(v)
		if rv.IsNil() {
			return nil
		}
		result := toMap(hint, rv.Elem().Interface(), indent)
		return &result
	case reflect.Slice:
		rv := reflect.ValueOf(v)

		result := reflect.ValueOf([]interface{}{})
		if !isPrimitiveType(rv.Type().Elem()) {
			result = reflect.ValueOf(ResourceSlice{})
		}

		for i := 0; i < rv.Len(); i++ {
			// fmt.Println(fmt.Sprintf("%s[%d]", hint, i))
			v := reflect.ValueOf(toMap(fmt.Sprintf("%s[%d]", hint, i), rv.Index(i).Interface(), indent+"  "))
			result = reflect.Append(result, v)
		}
		return result.Interface()
	case reflect.Struct:
		result := OrderedMap{}
		rv := reflect.ValueOf(v)
		for i := 0; i < rv.NumField(); i++ {
			field := rv.Field(i)
			// fmt.Println(indent + reflect.TypeOf(v).Field(i).Name)
			if rv.Type().Field(i).Anonymous {
				anonMember := toMap(hint+"."+rv.Type().Field(i).Name, field.Interface(), indent+"  ")
				if reflect.TypeOf(anonMember).Kind() != reflect.Map {
					panic(fmt.Errorf("anonymous field was expected to produce a map. actual: %T", anonMember))
				}
				for k, v := range anonMember.(OrderedMap) {
					result[k] = v
				}
			} else {
				typeField := rv.Type().Field(i)
				propertyName := unCamel(typeField.Name)
				isOptional := false
				if jsonTag, ok := typeField.Tag.Lookup("json"); ok {
					jsonTag = strings.TrimSpace(jsonTag)
					if len(jsonTag) > 0 {
						if strings.Contains(jsonTag, ",") {
							jsonPropertyName := strings.TrimSpace(strings.Split(jsonTag, ",")[0])
							if len(jsonPropertyName) > 0 {
								propertyName = jsonPropertyName
							}
						} else {
							propertyName = jsonTag
						}
						if strings.Contains(jsonTag, "omitempty") {
							isOptional = true
						}
					}
				}
				if propertyName != "-" {
					finalPropertyName := propertyName
					if reflect.TypeOf(v).Field(i).Type.Kind() == reflect.Interface {
						finalPropertyName = unref(reflect.TypeOf(field.Interface())).Name()
					}
					v := toMap(hint+"."+typeField.Name, field.Interface(), indent+"  ")
					setm(result, unCamel(finalPropertyName), v, isOptional)
				}
			}
		}
		return result
	case reflect.Map:
		switch reflect.TypeOf(v).Elem().Kind() {
		case reflect.String, reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
			values := OrderedMap{}
			rv := reflect.ValueOf(v)
			for _, key := range rv.MapKeys() {
				value := rv.MapIndex(key)
				storageKey := values.NextKey(key.Interface().(string))
				values[storageKey] = toMap(fmt.Sprintf("%s[%s]", hint, storageKey.Value), value.Interface(), indent+"  ")
			}
			result := ResourceSlice{}
			for k, v := range values {
				result = append(result, OrderedMap{
					OrderedMapKey{Value: "key", Order: 0}:    k.Value,
					OrderedMapKey{Value: "values", Order: 1}: v,
				})
			}
			return result
		case reflect.Struct:
			results := OrderedMap{}
			rv := reflect.ValueOf(v)
			for _, key := range rv.MapKeys() {
				omk := results.NextKey(key.Interface().(string))
				results[omk] = toMap(fmt.Sprintf("%s[%s]", hint, omk.Value), rv.MapIndex(key).Interface(), indent+"  ")
			}
			sliceResults := ResourceSlice{}
			for k, v := range results {
				ov := v.(OrderedMap)
				ov[ov.PrefixKey("key")] = k.Value
				sliceResults = append(sliceResults, ov)
			}
			return sliceResults
		case reflect.Ptr, reflect.Map, reflect.Slice:
			results := OrderedMap{}
			rv := reflect.ValueOf(v)
			for _, key := range rv.MapKeys() {
				omk := results.NextKey(key.Interface().(string))
				results[omk] = toMap(fmt.Sprintf("%s[%s]", hint, omk.Value), rv.MapIndex(key).Interface(), indent+"  ")
			}
			if len(results) == 0 {
				return ResourceSlice{}
			}
			elemType := reflect.TypeOf("")
			for _, v := range results {
				elemType = reflect.TypeOf(v)
			}
			switch elemType.Kind() {
			case reflect.Map:
				sliceResults := ResourceSlice{}
				for k, v := range results {
					ov := v.(OrderedMap)
					ov[ov.PrefixKey("key")] = k
					sliceResults = append(sliceResults, ov)
				}
				return sliceResults
			case reflect.Slice:
				sliceResults := ResourceSlice{}
				for k, v := range results {
					sliceResults = append(sliceResults, OrderedMap{
						OrderedMapKey{Value: "key", Order: 0}:    k.Value,
						OrderedMapKey{Value: "values", Order: 1}: v,
					})
				}
				return sliceResults
			default:
				return ResourceSlice{}
			}
		default:
			panic(fmt.Errorf("unsupported type %T", v))
		}
	default:
		panic(fmt.Errorf("unsupported type %T", v))
	}
}
