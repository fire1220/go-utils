package marshal

import (
	"encoding/json"
	"reflect"
	"time"
)

type DateTime struct {
	T   time.Time
	Tag reflect.StructTag `json:"-"`
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	t := d.T
	format := d.Tag.Get("datetime")
	if format == "" {
		format = time.DateTime
	}
	mapTime := map[string]string{
		time.DateTime: "0000-00-00 00:00:00",
		time.DateOnly: "0000-00-00",
		time.TimeOnly: "00:00:00",
	}
	if t.IsZero() {
		if v, ok := mapTime[format]; ok {
			return []byte(`"` + v + `"`), nil
		} else {
			return []byte(`""`), nil
		}
	}
	return []byte(`"` + t.Format(format) + `"`), nil
}

func Marshal(p any) ([]byte, error) {
	ref := reflect.ValueOf(p)
	typ := ref.Type()
	newField := make([]reflect.StructField, 0, ref.NumField())
	dateTimeReflectType := reflect.TypeOf(DateTime{})
	for i := 0; i < ref.NumField(); i++ {
		field := typ.Field(i)
		fieldType := field.Type
		if field.Type.String() == "time.Time" {
			fieldType = dateTimeReflectType
		}
		newField = append(newField, reflect.StructField{
			Name: field.Name,
			Type: fieldType,
			Tag:  field.Tag,
		})
	}
	newStruct := reflect.New(reflect.StructOf(newField)).Elem()
	for i := 0; i < newStruct.NumField(); i++ {
		oldField := ref.Field(i)
		oldFieldType := typ.Field(i)
		if oldField.Type().String() != "time.Time" {
			newStruct.Field(i).Set(oldField)
			continue
		}
		if v, ok := oldField.Interface().(time.Time); ok {
			newStruct.Field(i).Set(reflect.ValueOf(DateTime{T: v, Tag: oldFieldType.Tag}))
		}
	}
	return json.Marshal(newStruct.Interface())
}
