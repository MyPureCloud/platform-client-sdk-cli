package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Documentationv2searchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Documentationv2searchrequestDud struct { 
    


    


    


    


    


    


    


    

}

// Documentationv2searchrequest
type Documentationv2searchrequest struct { 
    // SortOrder - The sort order for results
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the resource that you want to sort the results by
    SortBy string `json:"sortBy"`


    // PageSize - The number of results per page
    PageSize int `json:"pageSize"`


    // PageNumber - The page of resources you want to retrieve
    PageNumber int `json:"pageNumber"`


    // Sort - Multi-value sort order, list of multiple sort values
    Sort []Searchsort `json:"sort"`


    // Types - Resource domain type to search
    Types []string `json:"types"`


    // Query - The search criteria
    Query []Documentationv2searchcriteria `json:"query"`


    // Aggregations - Aggregation criteria
    Aggregations []Documentationv2searchaggregation `json:"aggregations"`

}

// String returns a JSON representation of the model
func (o *Documentationv2searchrequest) String() string {
    
    
    
    
     o.Sort = []Searchsort{{}} 
     o.Types = []string{""} 
     o.Query = []Documentationv2searchcriteria{{}} 
     o.Aggregations = []Documentationv2searchaggregation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentationv2searchrequest) MarshalJSON() ([]byte, error) {
    type Alias Documentationv2searchrequest

    if Documentationv2searchrequestMarshalled {
        return []byte("{}"), nil
    }
    Documentationv2searchrequestMarshalled = true

    return json.Marshal(&struct {
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        Types []string `json:"types"`
        
        Query []Documentationv2searchcriteria `json:"query"`
        
        Aggregations []Documentationv2searchaggregation `json:"aggregations"`
        *Alias
    }{

        


        


        


        


        
        Sort: []Searchsort{{}},
        


        
        Types: []string{""},
        


        
        Query: []Documentationv2searchcriteria{{}},
        


        
        Aggregations: []Documentationv2searchaggregation{{}},
        

        Alias: (*Alias)(u),
    })
}

