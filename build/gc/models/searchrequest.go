package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchrequestDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Searchrequest
type Searchrequest struct { 
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


    // ReturnFields - A List of strings.  Possible values are any field in the resource you are searching on.  The other option is to use ALL_FIELDS, when this is provided all fields in the resource will be returned in the search results.
    ReturnFields []string `json:"returnFields"`


    // Expand - Provides more details about a specified resource
    Expand []string `json:"expand"`


    // Types - Resource domain type to search
    Types []string `json:"types"`


    // Query - The search criteria
    Query []Searchcriteria `json:"query"`


    // Aggregations - Aggregation criteria
    Aggregations []Searchaggregation `json:"aggregations"`

}

// String returns a JSON representation of the model
func (o *Searchrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Sort = []Searchsort{{}} 
    
    
    
     o.ReturnFields = []string{""} 
    
    
    
     o.Expand = []string{""} 
    
    
    
     o.Types = []string{""} 
    
    
    
     o.Query = []Searchcriteria{{}} 
    
    
    
     o.Aggregations = []Searchaggregation{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchrequest) MarshalJSON() ([]byte, error) {
    type Alias Searchrequest

    if SearchrequestMarshalled {
        return []byte("{}"), nil
    }
    SearchrequestMarshalled = true

    return json.Marshal(&struct { 
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        ReturnFields []string `json:"returnFields"`
        
        Expand []string `json:"expand"`
        
        Types []string `json:"types"`
        
        Query []Searchcriteria `json:"query"`
        
        Aggregations []Searchaggregation `json:"aggregations"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Sort: []Searchsort{{}},
        

        

        
        ReturnFields: []string{""},
        

        

        
        Expand: []string{""},
        

        

        
        Types: []string{""},
        

        

        
        Query: []Searchcriteria{{}},
        

        

        
        Aggregations: []Searchaggregation{{}},
        

        
        Alias: (*Alias)(u),
    })
}

