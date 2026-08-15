package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CasequeryjoberrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CasequeryjoberrorDud struct { 
    


    

}

// Casequeryjoberror
type Casequeryjoberror struct { 
    // Code - System-defined error code for the error.
    Code string `json:"code"`


    // Message - Error message for the failed job.
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Casequeryjoberror) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Casequeryjoberror) MarshalJSON() ([]byte, error) {
    type Alias Casequeryjoberror

    if CasequeryjoberrorMarshalled {
        return []byte("{}"), nil
    }
    CasequeryjoberrorMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

