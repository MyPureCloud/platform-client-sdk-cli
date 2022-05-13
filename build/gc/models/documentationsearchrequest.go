package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentationsearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentationsearchrequestDud struct { 
    


    


    


    


    


    

}

// Documentationsearchrequest
type Documentationsearchrequest struct { 
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


    // Query
    Query []Documentationsearchcriteria `json:"query"`

}

// String returns a JSON representation of the model
func (o *Documentationsearchrequest) String() string {
    
    
    
    
     o.Sort = []Searchsort{{}} 
     o.Query = []Documentationsearchcriteria{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentationsearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Documentationsearchrequest

    if DocumentationsearchrequestMarshalled {
        return []byte("{}"), nil
    }
    DocumentationsearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        SortOrder string `json:"sortOrder"`
        
        SortBy string `json:"sortBy"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        Query []Documentationsearchcriteria `json:"query"`
        *Alias
    }{

        


        


        


        


        
        Sort: []Searchsort{{}},
        


        
        Query: []Documentationsearchcriteria{{}},
        

        Alias: (*Alias)(u),
    })
}

