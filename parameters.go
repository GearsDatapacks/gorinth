package gorinth

import (
	"bytes"
	"fmt"
	"net/url"
	"strings"
)

// The parameters for searching for projects
type SearchQuery struct {
	// The query to search for
	Query string
	// The filters for the search results
	Facets [][]Facet
	// The order in which to sort results
	Index SearchIndex
	// The offset into search results. Skips this many results
	Offset uint
	// The number of results to return. Defaults to 10, must be <= 100
	Limit uint8
}

type Facet struct {
	// The key of this facet
	Key string
	// The operation to compare values
	Operation FacetOperation
	// The value to check for
	Value string
}

// The operation to compare values in search facets
type FacetOperation string

const (
	FacetEqual        FacetOperation = ":"
	FacetNotEqual     FacetOperation = "!="
	FacetLess         FacetOperation = "<"
	FacetLessEqual    FacetOperation = "<="
	FacetGreater      FacetOperation = ">"
	FacetGreaterEqual FacetOperation = ">="
)

// The order in which to sort search result
type SearchIndex string

const (
	IndexRelevance SearchIndex = "relevance"
	IndexDownloads SearchIndex = "downloads"
	IndexFollows   SearchIndex = "follows"
	IndexNewest    SearchIndex = "newest"
	IndexUpdated   SearchIndex = "updated"
)

func (s *SearchQuery) toQueryString() string {
	parameters := []string{}

	parameters = append(parameters, fmt.Sprintf("query=%s", url.QueryEscape(s.Query)))
	if s.Facets != nil {
		var facetString bytes.Buffer
		facetString.WriteByte('[')
		for _, facetList := range s.Facets {
			facetString.WriteByte('[')
			for i, facet := range facetList {
				if i != 0 {
					facetString.WriteByte(',')
				}
				fmt.Fprintf(&facetString, "\"%s%s%s\"", facet.Key, facet.Operation, facet.Value)
			}
			facetString.WriteByte(']')
		}
		facetString.WriteByte(']')

		parameters = append(parameters, fmt.Sprintf("facets=%s", url.QueryEscape(facetString.String())))
	}

	if len(s.Index) != 0 {
		parameters = append(parameters, fmt.Sprintf("index=%s", url.QueryEscape(string(s.Index))))
	}

	if s.Offset != 0 {
		parameters = append(parameters, fmt.Sprintf("offset=%d", s.Offset))
	}

	if s.Limit != 0 {
		limit := s.Limit
		if limit > 100 {
			limit = 100
		}
		parameters = append(parameters, fmt.Sprintf("limit=%d", limit))
	}

	return strings.Join(parameters, "&")
}

// The response to a "Search Projects" query
type SearchResponse struct {
	// The list of results
	Hits []SearchResult `json:"hits"`
	// The number of results skipped by the query
	Offset int `json:"offset"`
	// The number of results returned by the query
	Limit int `json:"limit"`
	// The total number of results which matched the query
	TotalHits int `json:"total_hits"`
}
