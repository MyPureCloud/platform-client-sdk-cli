package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationcustomattributessearchrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationcustomattributessearchrequestDud struct { 
    


    


    


    


    


    

}

// Conversationcustomattributessearchrequest
type Conversationcustomattributessearchrequest struct { 
    // Expand - Expand your search with bulk lookups
    Expand []string `json:"expand"`


    // PageSize - The number of results per page
    PageSize int `json:"pageSize"`


    // PageNumber - The page of resources you want to retrieve
    PageNumber int `json:"pageNumber"`


    // Sort - Multi-value sort order, list of multiple sort values
    Sort []Searchsort `json:"sort"`


    // SortBy - The field in the resource that you want to sort the results by
    SortBy string `json:"sortBy"`


    // SortOrder - The sort order for results
    SortOrder string `json:"sortOrder"`

}

// String returns a JSON representation of the model
func (o *Conversationcustomattributessearchrequest) String() string {
     o.Expand = []string{""} 
    
    
     o.Sort = []Searchsort{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationcustomattributessearchrequest) MarshalJSON() ([]byte, error) {
    type Alias Conversationcustomattributessearchrequest

    if ConversationcustomattributessearchrequestMarshalled {
        return []byte("{}"), nil
    }
    ConversationcustomattributessearchrequestMarshalled = true

    return json.Marshal(&struct {
        
        Expand []string `json:"expand"`
        
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        Sort []Searchsort `json:"sort"`
        
        SortBy string `json:"sortBy"`
        
        SortOrder string `json:"sortOrder"`
        *Alias
    }{

        
        Expand: []string{""},
        


        


        


        
        Sort: []Searchsort{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

