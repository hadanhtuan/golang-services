package es

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// constant do not have memory address(cannot use &)
var (
	customAnalyzer   = "custom_analyzer"
	customNormalizer = "custom_normalizer"

	CacheSearchDocument = "search:document"
	CacheSearchTracking = "search:tracking"

	SearchAll      = "ALL"
	SearchProperty = "PROPERTY"
	SearchUser     = "USER"

	PropertyIndex         = "property_index"
	UserIndices           = "user_index"
	SearchTrackingIndices = "search_tracking_index"
)
var IndicesMap = map[string]*create.Request{
	PropertyIndex: PropertyIndexCnf,
}

var PropertyIndexCnf = &create.Request{
	Settings: &types.IndexSettings{
		Analysis: &types.IndexSettingsAnalysis{
			Analyzer: map[string]types.Analyzer{
				customAnalyzer: &types.CustomAnalyzer{
					Tokenizer: "standard",
					// CharFilter: []string{"amenity_filter"},
					Filter: []string{"lowercase", "classic"},
				},
			},
			Normalizer: map[string]types.Normalizer{
				customNormalizer: &types.LowercaseNormalizer{
					Type: "lowercase",
				},
			},
			// CharFilter: map[string]types.CharFilter{
			// 	"amenity_filter": &types.PatternReplaceCharFilter{
			// 		Type:        "pattern_replace",
			// 		Pattern:     "\\|",
			// 		Replacement: &emptyReplacement,
			// 	},
			// },
		},
	},
	Mappings: &types.TypeMapping{
		Properties: map[string]types.Property{
			"id": types.KeywordProperty{
				Type: "keyword",
			},
			"createdAt": types.DateProperty{
				Type: "date",
			},
			"amenities": types.NestedProperty{
				Type: "nested",
				Properties: map[string]types.Property{
					"id": types.KeywordProperty{
						Type: "keyword",
					},
					"name": types.KeywordProperty{
						Type:       "keyword",
						Normalizer: &customNormalizer,
					},
				},
			},
			"propertyType": types.KeywordProperty{
				Type: "keyword",
			},
			"overallRate": types.FloatNumberProperty{
				Type: "float",
			},
			"wardCode": types.KeywordProperty{
				Type: "keyword",
			},
			"nationCode": types.KeywordProperty{
				Type: "keyword",
			},
			"lat": types.KeywordProperty{
				Type: "keyword",
			},
			"long": types.KeywordProperty{
				Type: "keyword",
			},
			"numGuests": types.IntegerNumberProperty{
				Type: "integer",
			},
			"numBeds": types.IntegerNumberProperty{
				Type: "integer",
			},
			"numBedrooms": types.IntegerNumberProperty{
				Type: "integer",
			},
			"numBathrooms": types.IntegerNumberProperty{
				Type: "integer",
			},
			"price": types.FloatNumberProperty{
				Type: "float",
			},
			"IsGuestFavor": types.BooleanProperty{
				Type: "boolean",
			},
			"title": types.TextProperty{
				Type:     "text",
				Analyzer: &customAnalyzer,
			},
			"body": types.TextProperty{
				Type:     "text",
				Analyzer: &customAnalyzer,
			},
		},
	},
}
