package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PagingspecMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PagingspecDud struct { 
    


    

}

// Pagingspec
type Pagingspec struct { 
    // PageSize - How many results per page
    PageSize int `json:"pageSize"`


    // PageNumber - How many pages in
    PageNumber int `json:"pageNumber"`

}

// String returns a JSON representation of the model
func (o *Pagingspec) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pagingspec) MarshalJSON() ([]byte, error) {
    type Alias Pagingspec

    if PagingspecMarshalled {
        return []byte("{}"), nil
    }
    PagingspecMarshalled = true

    return json.Marshal(&struct { 
        PageSize int `json:"pageSize"`
        
        PageNumber int `json:"pageNumber"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

