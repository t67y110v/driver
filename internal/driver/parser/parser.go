package parser

import (
	"encoding/binary"
	"reflect"

	"github.com/t67y110v/driver/internal/driver/model"
)

func MakeResponse(raw model.RawResponse, st reflect.Value) {
	raw.Offset = 6

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if !field.CanInterface() || field.Kind() == reflect.Struct {
			continue
		}

		var tmp reflect.Value
	SWITCH:

		switch field.Kind() {
		case reflect.Uint32:
			tmp = reflect.ValueOf(binary.LittleEndian.Uint32(raw.Get(4)))
			break SWITCH
		case reflect.Uint16:
			tmp = reflect.ValueOf(binary.LittleEndian.Uint16(raw.Get(2)))
			break SWITCH
		case reflect.Uint8:
			t := raw.Get(1)[0]
			tmp = reflect.ValueOf(t)
			break SWITCH
		case reflect.Bool:
			t := raw.Get(1)[0] == 1
			tmp = reflect.ValueOf(t)
			break SWITCH
		case reflect.String:
			pair := raw.Get(2)
			var slice []byte
			for !reflect.DeepEqual(pair, []byte{0x0D, 0x0A}) {
				slice = append(slice, pair[0])
				pair[0] = pair[1]
				pair[1] = raw.Get(1)[0]
			}
			tmp = reflect.ValueOf(string(slice))
			break SWITCH
		case reflect.Array:
			slice := raw.Get(field.Cap())
			tmp = reflect.New(field.Type()).Elem()
			for i := 0; i < field.Cap(); i++ {
				tmp.Index(i).Set(reflect.ValueOf(slice[i]))
			}
			break SWITCH
		}

		field.Set(tmp)
	}
}
