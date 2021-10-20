package esu

import (
	"log"
	"reflect"
	"runtime/debug"

	"github.com/fatih/structtag"
)

type Meta struct {
	info  map[uintptr]Key
	sf    map[uintptr]Key // 只装exported struct fields
	name  Key
}

func (m *Meta) recursive(i interface{}) {
	rv := reflect.ValueOf(i).Elem()
	rt := reflect.TypeOf(i).Elem()

	for i := 0; i < rt.NumField(); i++ {
		fv := rv.Field(i)
		field := rt.Field(i)

		// exported field
		if field.PkgPath == "" {
			_, ok := m.sf[fv.UnsafeAddr()]
			// 这里必须这么做, 不然会被顶掉
			if !ok {
				m.sf[fv.UnsafeAddr()] = Key(field.Name)
			}
		}

		switch field.Type.Kind() {
		case reflect.Struct:
			if fv.Addr().CanInterface() {
				m.recursive(fv.Addr().Interface())
			}
		default:
			break
		}

		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			continue
		}

		tag, err := tags.Get("json")
		if err != nil {
			continue
		}

		m.info[rv.Field(i).UnsafeAddr()] = Key(tag.Name)
	}
}

func (m *Meta) Init(i interface{}) {
	m.info = make(map[uintptr]Key)
	m.sf = make(map[uintptr]Key)

	if reflect.TypeOf(i).Elem().Kind() != reflect.Struct &&
		reflect.TypeOf(i).Kind() != reflect.Ptr {
		log.Fatal("NEED A PTR TO STRUCT")
	}

	name, ok := reflect.TypeOf(i).Elem().FieldByName("indexName")
	if !ok {
		log.Fatalln("NO indexName FIELD")
	}

	tags, err := structtag.Parse(string(name.Tag))
	if err != nil {
		log.Fatalln("parse tag err: ", err)
	}

	var tag *structtag.Tag
	hasdbFlag := false
	// dbu&db混用的情况
	for _, tempTag := range tags.Tags() {
		if tempTag.Key == "json" {
			if tempTag.Name != "" {
				hasdbFlag = true
				tag = tempTag
				break
			}
		}
	}

	if !hasdbFlag || tag == nil {
		log.Fatalln("get tag err: can not find a dbu tag, name.Tag:", string(name.Tag))
		// 不加的话goland标黄
		return
	}

	m.name = Key(tag.Name)

	m.recursive(i)
}

func (m *Meta) IndexName() string {
	return string(m.name)
}

func (m *Meta) Tag(i interface{}) Key {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		log.Println("NEED PTR", string(debug.Stack()))
		return ""
	}
	addr := reflect.ValueOf(i).Elem().UnsafeAddr()
	if m.info[addr] == "" {
		log.Println("NO FIELD", string(debug.Stack()))
		return ""
	}
	return Key(m.info[addr])
}


func (m *Meta) Field(i interface{}) Key {
	if reflect.TypeOf(i).Kind() != reflect.Ptr {
		log.Println("NEED PTR", string(debug.Stack()))
		return ""
	}
	addr := reflect.ValueOf(i).Elem().UnsafeAddr()
	if m.sf[addr] == "" {
		log.Println("NO FIELD", string(debug.Stack()))
		return ""
	}
	return Key(m.sf[addr])
}