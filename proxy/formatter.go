package proxy

import (
	// "fmt"
	"strings"

	// "github.com/fatih/structs"
	// "github.com/davecgh/go-spew/spew"
	"github.com/k0kubun/pp"
)

// EntityFormatter formats the response data
type EntityFormatter interface {
	Format(entity Response) Response
	Targets() []map[string]string
}

// EntityFormatterFunc holds the formatter function
type EntityFormatterFunc struct {
	Func func(Response) Response
}

// Format implements the EntityFormatter interface
func (e EntityFormatterFunc) Format(entity Response) Response {
	return e.Func(entity)
}

// Targets implements the EntityFormatter interface
func (e EntityFormatterFunc) Targets() []map[string]string {
	pp.Print(e)
	return nil
}

type propertyFilter func(entity *Response)

// type targetsFilter func(targets *[]map[string]string)

type entityFormatter struct {
	Target         string
	Prefix         string
	PropertyFilter propertyFilter
	Mapping        map[string]string
	Paths          []map[string]string
}

// NewEntityFormatter creates an entity formatter with the received params
func NewEntityFormatter(target string, whitelist, blacklist []string, group string, mappings map[string]string, paths []map[string]string) EntityFormatter {
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
		Paths:          paths,
	}
}

func (e entityFormatter) Targets() []map[string]string {
	return e.Paths
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
		// fmt.Printf("keys=%s , length: %d, len(wl[keys[0]]): %d, len(keys[1:]): %d\n", keys, len(keys), len(wl[keys[0]]), len(keys[1:]))
		if len(keys) > 1 {
			if _, ok := wl[keys[0]]; ok {
				for _, k := range keys[1:] {
					wl[keys[0]][k] = nil
				}
			} else {
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
			if sub, ok := wl[k]; ok {
				if len(sub) > 0 {
					if tmp := whitelistFilterSub(v, sub); len(tmp) > 0 {
						accumulator[k] = tmp
					}
				} else {
					accumulator[k] = v
				}
			}
			// fmt.Printf("key=%s, value=%s, length: %d, len(wl[k]): %d\n", k, v, len(k), len(wl[k]))
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

// encoding.To("xml", response)
