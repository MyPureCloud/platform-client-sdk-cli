package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulkcallbackdisconnectrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulkcallbackdisconnectrequestDud struct { 
    

}

// Bulkcallbackdisconnectrequest
type Bulkcallbackdisconnectrequest struct { 
    // CallbackDisconnectIdentifiers - The list of requests to disconnect callbacks in bulk
    CallbackDisconnectIdentifiers []Callbackdisconnectidentifier `json:"callbackDisconnectIdentifiers"`

}

// String returns a JSON representation of the model
func (o *Bulkcallbackdisconnectrequest) String() string {
    
    
     o.CallbackDisconnectIdentifiers = []Callbackdisconnectidentifier{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulkcallbackdisconnectrequest) MarshalJSON() ([]byte, error) {
    type Alias Bulkcallbackdisconnectrequest

    if BulkcallbackdisconnectrequestMarshalled {
        return []byte("{}"), nil
    }
    BulkcallbackdisconnectrequestMarshalled = true

    return json.Marshal(&struct { 
        CallbackDisconnectIdentifiers []Callbackdisconnectidentifier `json:"callbackDisconnectIdentifiers"`
        
        *Alias
    }{
        

        
        CallbackDisconnectIdentifiers: []Callbackdisconnectidentifier{{}},
        

        
        Alias: (*Alias)(u),
    })
}

