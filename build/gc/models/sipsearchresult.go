package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SipsearchresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SipsearchresultDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Sipsearchresult
type Sipsearchresult struct { 
    


    // Status - Status of the search request
    Status int `json:"status"`


    // Sid - Session id associated to the search request
    Sid string `json:"sid"`


    // Auth - Auth token used for this search request
    Auth string `json:"auth"`


    // Message - Any messages returned from homer as part of the response
    Message string `json:"message"`


    // Data - Homer search data that is returned
    Data []Homerrecord `json:"data"`


    // Count - Number of records returned
    Count int `json:"count"`


    

}

// String returns a JSON representation of the model
func (o *Sipsearchresult) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Data = []Homerrecord{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sipsearchresult) MarshalJSON() ([]byte, error) {
    type Alias Sipsearchresult

    if SipsearchresultMarshalled {
        return []byte("{}"), nil
    }
    SipsearchresultMarshalled = true

    return json.Marshal(&struct { 
        
        
        Status int `json:"status"`
        
        Sid string `json:"sid"`
        
        Auth string `json:"auth"`
        
        Message string `json:"message"`
        
        Data []Homerrecord `json:"data"`
        
        Count int `json:"count"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        
        Data: []Homerrecord{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

