package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocationsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocationsearchrequestDud struct { 
    


    


    


    


    


    


    

}

// Locationsearchrequest
type Locationsearchrequest struct { 
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


    // Expand - Provides more details about a specified resource
    Expand []string `json:"expand"`


    // Query
    Query []Locationsearchcriteria `json:"query"`

}

// String returns a JSON representation of the model
func (o *Locationsearchrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Sort = []Searchsort{{}} 
    
    
    
     o.Expand = []string{""} 
    
    
    
     o.Query = []Locationsearchcriteria{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Locationsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Locationsearchrequest

    if LocationsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    LocationsearchrequestMarshalled = true

    return json.Marshal(&struct { 
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        Expand []string `json:"expand"`
        
        Query []Locationsearchcriteria `json:"query"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        Sort: []Searchsort{{}},
        

        

        
        Expand: []string{""},
        

        

        
        Query: []Locationsearchcriteria{{}},
        

        
        Alias: (*Alias)(u),
    })
}

