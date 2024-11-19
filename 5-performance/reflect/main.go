package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Config struct {
	Name    string `json:"server-name"` // CONFIG_SERVER_NAME
	IP      string `json:"server-ip"`   // CONFIG_SERVER_IP
	URL     string `json:"server-url"`  // CONFIG_SERVER_URL
	Timeout string `json:"timeout"`     // CONFIG_TIMEOUT
}

func readConfig() *Config {
	// read from xxx.json，省略
	config := Config{}
	typ := reflect.TypeOf(config)
	value := reflect.Indirect(reflect.ValueOf(&config))
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)
		fmt.Printf("%+v\n", f)
		if v, ok := f.Tag.Lookup("json"); ok {
			key := fmt.Sprintf("CONFIG_%s", strings.ReplaceAll(strings.ToUpper(v), "-", "_"))
			if env, exist := os.LookupEnv(key); exist {
				value.FieldByName(f.Name).Set(reflect.ValueOf(env))
			}
		}
	}
	return &config
}

func convert(i interface{}) int64 {
	typ := reflect.TypeOf(i)
	switch typ.Kind() {
	case reflect.Int:
		return int64(i.(int))
	case reflect.Int8:
		return int64(i.(int8))
	case reflect.Int16:
		return int64(i.(int16))
	case reflect.Int32:
		return int64(i.(int32))
	case reflect.Int64:
		return i.(int64)
	default:
		panic("not support")
	}
}

func add(a, b int) int {
	return a + b
}

func reflectAdd(a, b interface{}) int64 {
	return convert(a) + convert(b)
}

func main() {
	os.Setenv("CONFIG_SERVER_NAME", "global_server")
	os.Setenv("CONFIG_SERVER_IP", "10.0.0.1")
	os.Setenv("CONFIG_SERVER_URL", "geektutu.com")
	c := readConfig()
	fmt.Printf("%+v", c)
}
