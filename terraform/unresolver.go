package terraform

import (
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ToTerraform has no documentation
func ToTerraform(v interface{}, resourceData *schema.ResourceData) error {
	if _, err := toTerraform(v); err != nil {
		return err
	}
	return nil
}

func toTerraform(v interface{}) (interface{}, error) {
	if v == nil {
		return nil, nil
	}
	var err error
	switch reflect.TypeOf(v).Kind() {
	case reflect.Ptr:
		rv := reflect.ValueOf(v)
		if rv.IsNil() {
			return nil, nil
		}
		rv = rv.Elem()
		if (rv.Type().Kind() == reflect.Ptr) && (rv.IsNil()) {
			return nil, nil
		}
		return toTerraform(rv.Interface())
	case reflect.Struct:
		result := map[string]interface{}{}
		rv := reflect.ValueOf(v)
		for i := 0; i < rv.NumField(); i++ {
			iField := rv.Field(i)
			field := rv.Type().Field(i)
			if field.Anonymous {
				var anonResults interface{}
				if anonResults, err = toTerraform(iField.Interface()); err != nil {
					if anonResults != nil {
						for k, v := range anonResults.(map[string]interface{}) {
							result[k] = v
						}
					}
				}
			} else {
				if propertyName := propName(field); len(propertyName) > 0 {
					var propertyResult interface{}
					if propertyResult, err = toTerraform(iField.Interface()); err != nil {
						return nil, err
					}
					if propertyResult != nil {
						switch typedPropertyResult := propertyResult.(type) {
						case map[string]interface{}:
							result[propertyName] = []interface{}{typedPropertyResult}
						default:
							result[propertyName] = typedPropertyResult
						}

					}
				}
			}
		}
		if len(result) == 0 {
			return nil, nil
		}
		return result, nil
	case reflect.Slice:
		result := []interface{}{}
		rv := reflect.ValueOf(v)
		for i := 0; i < rv.Len(); i++ {
			elem := rv.Index(i)
			var elemResult interface{}
			if elemResult, err = toTerraform(elem.Interface()); err != nil {
				return nil, err
			}
			result = append(result, elemResult)
		}
		return result, nil
	case reflect.String, reflect.Bool, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uint, reflect.Float32, reflect.Float64:
		return v, nil
	case reflect.Map:
		result := []interface{}{}
		rv := reflect.ValueOf(v)
		for _, k := range rv.MapKeys() {
			elem := rv.MapIndex(k).Interface()
			var elemResult interface{}
			if elemResult, err = toTerraform(elem); err != nil {
				return nil, err
			}
			switch typedElemResult := elemResult.(type) {
			case map[string]interface{}:
				typedElemResult["key"] = k
				result = append(result, typedElemResult)
			default:
				result = append(result, map[string]interface{}{
					"key":    k,
					"values": typedElemResult,
				})
			}
		}
		return result, nil
	default:
		return nil, fmt.Errorf("unsupported type %T (kind: %v)", v, reflect.TypeOf(v).Kind())
	}
}
