package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TagqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TagqueryrequestDud struct { 
    


    


    

}

// Tagqueryrequest
type Tagqueryrequest struct { 
    // Query
    Query string `json:"query"`


    // PageNumber
    PageNumber int `json:"pageNumber"`


    // PageSize
    PageSize int `json:"pageSize"`

}

// String returns a JSON representation of the model
func (o *Tagqueryrequest) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Tagqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Tagqueryrequest

    if TagqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    TagqueryrequestMarshalled = true

    return json.Marshal(&struct { 
        Query string `json:"query"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

