package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasequeryjobcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasequeryjobcreateDud struct { 
    


    


    


    


    


    

}

// Casequeryjobcreate
type Casequeryjobcreate struct { 
    // PageSize - The total page size requested (default 25).
    PageSize int `json:"pageSize"`


    // PageNumber - The requested page number.
    PageNumber int `json:"pageNumber"`


    // Filters - List of filter objects to be used in the search. Use an empty list to run the query with no filters.
    Filters []Casequeryjobfilter `json:"filters"`


    // Sort - Sort order for results.
    Sort Casequeryjobsort `json:"sort"`


    // Attributes - List of entity attributes to be retrieved in the result.
    Attributes []string `json:"attributes"`


    // Expands - Attributes to expand on each case in the job results. Expands are stored on the job and enriched by PubAPI when results are fetched.
    Expands []string `json:"expands"`

}

// String returns a JSON representation of the model
func (o *Casequeryjobcreate) String() string {
    
    
     o.Filters = []Casequeryjobfilter{{}} 
    
     o.Attributes = []string{""} 
     o.Expands = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casequeryjobcreate) MarshalJSON() ([]byte, error) {
    type Alias Casequeryjobcreate

    if CasequeryjobcreateMarshalled {
        return []byte("{}"), nil
    }
    CasequeryjobcreateMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Filters []Casequeryjobfilter `json:"filters"`
        
        Sort Casequeryjobsort `json:"sort"`
        
        Attributes []string `json:"attributes"`
        
        Expands []string `json:"expands"`
        *Alias
    }{

        


        


        
        Filters: []Casequeryjobfilter{{}},
        


        


        
        Attributes: []string{""},
        


        
        Expands: []string{""},
        

        Alias: (*Alias)(u),
    })
}

