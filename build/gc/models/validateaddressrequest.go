package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValidateaddressrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValidateaddressrequestDud struct { 
    

}

// Validateaddressrequest
type Validateaddressrequest struct { 
    // Address - Address schema
    Address Streetaddress `json:"address"`

}

// String returns a JSON representation of the model
func (o *Validateaddressrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Validateaddressrequest) MarshalJSON() ([]byte, error) {
    type Alias Validateaddressrequest

    if ValidateaddressrequestMarshalled {
        return []byte("{}"), nil
    }
    ValidateaddressrequestMarshalled = true

    return json.Marshal(&struct {
        
        Address Streetaddress `json:"address"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

