package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidateaddressresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidateaddressresponseDud struct { 
    


    

}

// Validateaddressresponse
type Validateaddressresponse struct { 
    // Valid - Was the passed in address valid
    Valid bool `json:"valid"`


    // Response - Subscriber schema
    Response Subscriberresponse `json:"response"`

}

// String returns a JSON representation of the model
func (o *Validateaddressresponse) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validateaddressresponse) MarshalJSON() ([]byte, error) {
    type Alias Validateaddressresponse

    if ValidateaddressresponseMarshalled {
        return []byte("{}"), nil
    }
    ValidateaddressresponseMarshalled = true

    return json.Marshal(&struct { 
        Valid bool `json:"valid"`
        
        Response Subscriberresponse `json:"response"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

