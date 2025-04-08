package rss

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*Feed, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	xmlFeed := &Feed{}
	err = xml.Unmarshal(bodyBytes, xmlFeed)
	if err != nil {
		return nil, err
	}

	xmlFeed.Unescape()

	return xmlFeed, nil
}
