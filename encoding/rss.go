package encoding

import (
	"fmt"
	"io"

	"github.com/k0kubun/pp"
	"github.com/mmcdole/gofeed"
	//	"github.com/roscopecoltran/krakend/logging"
)

// RSS is the key for the rss encoding
const RSS = "rss"

// NewRSSDecoder returns the RSS decoder
func NewRSSDecoder() Decoder {
	fmt.Printf("NewRSSDecoder(..)")
	fp := gofeed.NewParser()
	return func(r io.Reader, v *map[string]interface{}, paths []map[string]string) error {
		fmt.Printf("NewRSSDecoder(..), paths:")
		pp.Print(paths)
		feed, err := fp.Parse(r)
		if err != nil {
			fmt.Println("error while init fp.Parse(...),  err:\n", err)
			return err
		}
		*(v) = map[string]interface{}{
			"items":       feed.Items,
			"author":      feed.Author,
			"categories":  feed.Categories,
			"custom":      feed.Custom,
			"copyright":   feed.Copyright,
			"description": feed.Description,
			"type":        feed.FeedType,
			"language":    feed.Language,
			"title":       feed.Title,
			"published":   feed.Published,
			"updated":     feed.Updated,
		}
		if feed.Image != nil {
			(*v)["img_url"] = feed.Image.URL
		}
		return nil
	}
}
