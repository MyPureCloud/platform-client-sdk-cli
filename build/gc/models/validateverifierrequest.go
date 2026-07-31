package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidateverifierrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidateverifierrequestDud struct { 
    


    

}

// Validateverifierrequest
type Validateverifierrequest struct { 
    // Enable - Whether to enable the verifier upon successful validation.
    Enable bool `json:"enable"`


    // Token - The verification token to validate against the verifier.
    Token string `json:"token"`

}

// String returns a JSON representation of the model
func (o *Validateverifierrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validateverifierrequest) MarshalJSON() ([]byte, error) {
    type Alias Validateverifierrequest

    if ValidateverifierrequestMarshalled {
        return []byte("{}"), nil
    }
    ValidateverifierrequestMarshalled = true

    return json.Marshal(&struct {
        
        Enable bool `json:"enable"`
        
        Token string `json:"token"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

