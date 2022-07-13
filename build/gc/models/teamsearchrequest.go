package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TeamsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TeamsearchrequestDud struct { 
    


    


    


    


    


    

}

// Teamsearchrequest
type Teamsearchrequest struct { 
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


    // Query - Team Search Criteria
    Query []Teamsearchcriteria `json:"query"`

}

// String returns a JSON representation of the model
func (o *Teamsearchrequest) String() string {
    
    
    
    
     o.Sort = []Searchsort{{}} 
     o.Query = []Teamsearchcriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Teamsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Teamsearchrequest

    if TeamsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    TeamsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        Query []Teamsearchcriteria `json:"query"`
        *Alias
    }{

        


        


        


        


        
        Sort: []Searchsort{{}},
        


        
        Query: []Teamsearchcriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

