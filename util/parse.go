package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

//绑定参数
func Bind(req *http.Request,obj interface{}) error  {
	contentType := req.Header.Get("Content-Type")
	if strings.Contains(strings.ToLower(contentType),"application/json"){
		return BindJson(req,obj)
	}
	if strings.Contains(strings.ToLower(contentType),"application/x-www-form-urlencoded"){
		return BindForm(req,obj)
	}
	return errors.New("系统无法识别的类型")
}
//绑定json
func BindJson(req *http.Request,obj interface{}) error  {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return err
	}
	if err = json.Unmarshal(body, obj); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return err
	}
	return nil
}
//绑定表单
func BindForm(req *http.Request,obj interface{}) error  {
	fmt.Println(req.Form.Encode())
	err := mapForm(obj,req)
	return err
}

func mapForm(ptr interface{},req *http.Request) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	// 创建字段映射表，键为有效名称
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// 对请求中的每个参数更新结构体中对应的字段
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // 忽略不能识别的 HTTP 参数
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return fmt.Errorf("%s: %v", name, err)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)

	case reflect.Int:
	case reflect.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)

	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)

	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}