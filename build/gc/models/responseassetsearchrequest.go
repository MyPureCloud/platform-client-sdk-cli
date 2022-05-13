package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResponseassetsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResponseassetsearchrequestDud struct { 
    


    


    


    


    

}

// Responseassetsearchrequest
type Responseassetsearchrequest struct { 
    // PageSize - The number of results per page. Default: 25, Maximum: 100.
    PageSize int `json:"pageSize"`


    // PageNumber - The page of resources you want to retrieve
    PageNumber int `json:"pageNumber"`


    // SortOrder - The sort order for results
    SortOrder string `json:"sortOrder"`


    // SortBy - The field in the resource that you want to sort the results by
    SortBy string `json:"sortBy"`


    // Query - Filter the query results.
    Query []Responseassetfilter `json:"query"`

}

// String returns a JSON representation of the model
func (o *Responseassetsearchrequest) String() string {
    
    
    
    
     o.Query = []Responseassetfilter{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Responseassetsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Responseassetsearchrequest

    if ResponseassetsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    ResponseassetsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        Query []Responseassetfilter `json:"query"`
        *Alias
    }{

        


        


        


        


        
        Query: []Responseassetfilter{{}},
        

        Alias: (*Alias)(u),
    })
}

