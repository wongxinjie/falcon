package dbu

import (
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
	"strings"

	"github.com/fatih/structtag"
)

type Meta struct {
	info  map[uintptr]Key
	sf    map[uintptr]Key // 只装exported struct fields
	name  Key
	alias Key
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

		tag, err := tags.Get("gorm")
		if err != nil {
			continue
		}

		var tagName string
		for _, tagElem := range strings.Split(tag.Name, ";") {
			if !strings.HasPrefix(tagElem, "column:") {
				continue
			}
			tagName = tagElem[len("column:"):]
		}

		m.info[rv.Field(i).UnsafeAddr()] = Key(tagName)
	}
}

func (m *Meta) Init(i interface{}) {
	m.info = make(map[uintptr]Key)
	m.sf = make(map[uintptr]Key)

	if reflect.TypeOf(i).Elem().Kind() != reflect.Struct &&
		reflect.TypeOf(i).Kind() != reflect.Ptr {
		log.Fatal("NEED A PTR TO STRUCT")
	}

	name, ok := reflect.TypeOf(i).Elem().FieldByName("tableName")
	if !ok {
		log.Fatalln("NO TABLENAME FIELD")
	}

	tags, err := structtag.Parse(string(name.Tag))
	if err != nil {
		log.Fatalln("parse tag err: ", err)
	}

	var tag *structtag.Tag
	hasdbFlag := false
	for _, tempTag := range tags.Tags() {
		if tempTag.Key == "gorm" {
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

	for _, option := range tag.Options {
		if strings.Contains(option, "alias") {
			strs := strings.Split(option, ":")
			m.alias = Key(strs[1])
			break
		}
	}

	if len(m.alias) == 0 {
		log.Fatalln("NO ALIAS PROVIDED")
	}

	m.recursive(i)
}

func (m *Meta) TableName() string {
	return string(m.name)
}

func (m *Meta) AliasPk() string {
	return m.alias.V() + "." + "id"
}

func (m *Meta) Alias() string {
	return string(m.alias)
}

func (m *Meta) AliasCus(cus string) string {
	return fmt.Sprintf("%s_%s", string(m.alias), cus)
}

func (m *Meta) AliasAny() string {
	return fmt.Sprintf("%s.*", string(m.alias))
}

func (m *Meta) AliasTag(i interface{}) Key {
	k := m.Tag(i)
	return Key(fmt.Sprintf("%s.%s", m.alias, k))
}

func (m *Meta) AliasCusTag(cus string, i interface{}) Key {
	k := m.Tag(i)
	return Key(fmt.Sprintf("%s.%s", m.AliasCus(cus), k))
}

func (m *Meta) AliasTagEscape(i interface{}) Key {
	k := m.Tag(i)
	return Key(fmt.Sprintf(`%s."%s"`, m.alias, k))
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
