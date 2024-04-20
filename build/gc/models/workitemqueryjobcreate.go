package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemqueryjobcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemqueryjobcreateDud struct { 
    


    


    


    


    


    

}

// Workitemqueryjobcreate
type Workitemqueryjobcreate struct { 
    // PageSize - The total page size requested. Default 25
    PageSize int `json:"pageSize"`


    // PageNumber - The page number requested
    PageNumber int `json:"pageNumber"`


    // Filters - List of filter objects to be used in the search.
    Filters []Workitemqueryjobfilter `json:"filters"`


    // Expands - List of entity attributes to be expanded in the result.
    Expands []string `json:"expands"`


    // Attributes - List of entity attributes to be retrieved in the result.
    Attributes []string `json:"attributes"`


    // Sort - Sort
    Sort Workitemqueryjobsort `json:"sort"`

}

// String returns a JSON representation of the model
func (o *Workitemqueryjobcreate) String() string {
    
    
     o.Filters = []Workitemqueryjobfilter{{}} 
     o.Expands = []string{""} 
     o.Attributes = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemqueryjobcreate) MarshalJSON() ([]byte, error) {
    type Alias Workitemqueryjobcreate

    if WorkitemqueryjobcreateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemqueryjobcreateMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Filters []Workitemqueryjobfilter `json:"filters"`
        
        Expands []string `json:"expands"`
        
        Attributes []string `json:"attributes"`
        
        Sort Workitemqueryjobsort `json:"sort"`
        *Alias
    }{

        


        


        
        Filters: []Workitemqueryjobfilter{{}},
        


        
        Expands: []string{""},
        


        
        Attributes: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

