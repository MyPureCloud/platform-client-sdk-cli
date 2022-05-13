package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FailedobjectMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FailedobjectDud struct { 
    


    


    


    

}

// Failedobject
type Failedobject struct { 
    // Id
    Id string `json:"id"`


    // Version
    Version string `json:"version"`


    // Name
    Name string `json:"name"`


    // ErrorCode
    ErrorCode string `json:"errorCode"`

}

// String returns a JSON representation of the model
func (o *Failedobject) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Failedobject) MarshalJSON() ([]byte, error) {
    type Alias Failedobject

    if FailedobjectMarshalled {
        return []byte("{}"), nil
    }
    FailedobjectMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Version string `json:"version"`
        
        Name string `json:"name"`
        
        ErrorCode string `json:"errorCode"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

