package proxy

import (
	"fmt"
	"strings"
	// map_path "github.com/firewut/go-json-map"
	// "github.com/bloglovin/obpath"
)

/*
	Refs:
	- https://github.com/bloglovin/obpath
		".store",
		".store.books",
		".store.*",
		"..Author",
		".store.counts[*]",
		".store.counts[3]",
		".store.counts[1:2]",
		".store.counts[-2:]",
		".store.counts[:1]",
		".store.counts[:1].Price",
		"..books[*](has(@.ISBN))",
		".store.books[*](!empty(@.ISBN))",
		".store.books[*](eq(@.Price, 8.99))",
		".store.books[0:4](eq(@.Author, \"Louis L'Amour\"))",
		"..books.*(between(@.Price, 8, 10)).Title",
		"..books[*](gt(@.Price, 9))",
		"..books[*](has(@.Metadata))",
		"..books[*](contains(@.Title, 'R')).Title",
		"..books[*](cicontains(@.Title, 'R')).Title",
		".store.*[*](gt(@.Price, 18))"
		- cat testdata/search.json  | obp --path=".items.books"
*/

// "github.com/roscopecoltran/krakend/logging"

// EntityFormatter formats the response data
type EntityFormatter interface {
	Format(entity Response) Response
}

// EntityFormatterFunc holds the formatter function
type EntityFormatterFunc struct {
	Func func(Response) Response
}

// Format implements the EntityFormatter interface
func (e EntityFormatterFunc) Format(entity Response) Response {
	return e.Func(entity)
}

type propertyFilter func(entity *Response)

type entityFormatter struct {
	Target         string
	Prefix         string
	PropertyFilter propertyFilter
	Mapping        map[string]string
}

// NewEntityFormatter creates an entity formatter with the received params
func NewEntityFormatter(target string, whitelist, blacklist []string, group string, mappings map[string]string) EntityFormatter {
	var propertyFilter propertyFilter
	if len(whitelist) > 0 {
		propertyFilter = newWhitelistingFilter(whitelist)
	} else {
		propertyFilter = newBlacklistingFilter(blacklist)
	}
	sanitizedMappings := make(map[string]string, len(mappings))
	for i, m := range mappings {
		v := strings.Split(m, ".")
		sanitizedMappings[i] = v[0]
	}
	return entityFormatter{
		Target:         target,
		Prefix:         group,
		PropertyFilter: propertyFilter,
		Mapping:        sanitizedMappings,
	}
}

// Format implements the EntityFormatter interface
func (e entityFormatter) Format(entity Response) Response {
	if e.Target != "" {
		extractTarget(e.Target, &entity)
	}
	if len(entity.Data) > 0 {
		e.PropertyFilter(&entity)
	}
	if len(entity.Data) > 0 {
		for formerKey, newKey := range e.Mapping {
			if v, ok := entity.Data[formerKey]; ok {
				entity.Data[newKey] = v
				delete(entity.Data, formerKey)
			}
		}
	}
	if e.Prefix != "" {
		entity.Data = map[string]interface{}{e.Prefix: entity.Data}
	}
	return entity
}

func extractTarget(target string, entity *Response) {
	if tmp, ok := entity.Data[target]; ok {
		entity.Data, ok = tmp.(map[string]interface{})
		if !ok {
			entity.Data = map[string]interface{}{}
		}
	} else {
		entity.Data = map[string]interface{}{}
	}
}

func newWhitelistingFilter(whitelist []string) propertyFilter {
	wl := make(map[string]map[string]interface{}, len(whitelist))
	for _, k := range whitelist {
		keys := strings.Split(k, ".")
		tmp := make(map[string]interface{}, len(keys)-1)
		fmt.Printf("keys=%s \n", keys)
		if len(keys) > 1 {
			fmt.Printf("wl[keys[0]]=%s \n", wl[keys[0]])
			if _, ok := wl[keys[0]]; ok {
				fmt.Printf("keys[1:]=%s \n", keys[1:])
				for _, k := range keys[1:] {
					wl[keys[0]][k] = nil //
				}
			} else {
				fmt.Printf("keys[1:]=%s \n", keys[1:])
				for _, k := range keys[1:] {
					tmp[k] = nil
				}
				wl[keys[0]] = tmp
			}
		} else {
			wl[keys[0]] = tmp
		}
	}

	return func(entity *Response) {
		accumulator := make(map[string]interface{}, len(whitelist))
		for k, v := range entity.Data {
			// property, err = map_path.GetProperty(document, "one.two.three", ".")
			// fmt.Println(property)
			if sub, ok := wl[k]; ok {
				if len(sub) > 0 {
					if tmp := whitelistFilterSub(v, sub); len(tmp) > 0 {
						accumulator[k] = tmp
					}
				} else {
					accumulator[k] = v
				}
			}
		}
		*entity = Response{Data: accumulator, IsComplete: entity.IsComplete}
	}
}

func whitelistFilterSub(v interface{}, whitelist map[string]interface{}) map[string]interface{} {
	entity, ok := v.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}
	tmp := make(map[string]interface{}, len(whitelist))
	for k, v := range entity {
		if _, ok := whitelist[k]; ok {
			tmp[k] = v
		}
	}
	return tmp
}

func newBlacklistingFilter(blacklist []string) propertyFilter {
	bl := make(map[string][]string, len(blacklist))
	for _, key := range blacklist {
		keys := strings.Split(key, ".")
		if len(keys) > 1 {
			if sub, ok := bl[keys[0]]; ok {
				bl[keys[0]] = append(sub, keys[1])
			} else {
				bl[keys[0]] = []string{keys[1]}
			}
		} else {
			bl[keys[0]] = []string{}
		}
	}

	return func(entity *Response) {
		for k, sub := range bl {
			if len(sub) == 0 {
				delete(entity.Data, k)
			} else {
				if tmp := blacklistFilterSub(entity.Data[k], sub); len(tmp) > 0 {
					entity.Data[k] = tmp
				}
			}
		}
	}
}

func blacklistFilterSub(v interface{}, blacklist []string) map[string]interface{} {
	tmp, ok := v.(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}
	for _, key := range blacklist {
		delete(tmp, key)
	}
	return tmp
}
