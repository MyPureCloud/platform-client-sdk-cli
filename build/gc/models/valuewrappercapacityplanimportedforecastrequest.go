package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ValuewrappercapacityplanimportedforecastrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ValuewrappercapacityplanimportedforecastrequestDud struct { 
    

}

// Valuewrappercapacityplanimportedforecastrequest
type Valuewrappercapacityplanimportedforecastrequest struct { 
    // Value - The value for the associated field
    Value Capacityplanimportedforecastrequest `json:"value"`

}

// String returns a JSON representation of the model
func (o *Valuewrappercapacityplanimportedforecastrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Valuewrappercapacityplanimportedforecastrequest) MarshalJSON() ([]byte, error) {
    type Alias Valuewrappercapacityplanimportedforecastrequest

    if ValuewrappercapacityplanimportedforecastrequestMarshalled {
        return []byte("{}"), nil
    }
    ValuewrappercapacityplanimportedforecastrequestMarshalled = true

    return json.Marshal(&struct {
        
        Value Capacityplanimportedforecastrequest `json:"value"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

