package configuration

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func InitConfig(s *any) {
	configType := reflect.TypeOf(s)
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		tag := field.Tag.Get("env")

		if tag != "" {
			configDefinition := strings.TrimPrefix(tag, "env:")
			parts := strings.Split(configDefinition, ":")
			var envValue string
			var ok bool
			if len(parts) > 0 {
				envValue, ok = os.LookupEnv(parts[0])
				if ok && envValue != "" {
					setValue(s, field.Name, envValue)
					continue
				}
			}

			if len(parts) == 2 {
				setValue(s, field.Name, parts[1])
			}
		}
	}
}

func setValue(s *any, fieldName string, value string) {
	v := reflect.ValueOf(s).Elem()
	field := v.FieldByName(fieldName)
	if field.IsValid() && field.CanSet() {
		switch field.Kind() {
		case reflect.Int:
			var intValue int
			fmt.Sscanf(value, "%d", &intValue)
			field.SetInt(int64(intValue))
		case reflect.String:
			field.SetString(value)
		case reflect.Bool:
			var boolValue bool
			fmt.Sscanf(value, "%t", &boolValue)
			field.SetBool(boolValue)
		}
	}
}
