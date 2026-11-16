package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthenticationextensionslargebloboutputsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthenticationextensionslargebloboutputsDud struct { 
    


    

}

// Authenticationextensionslargebloboutputs
type Authenticationextensionslargebloboutputs struct { 
    // Supported - Whether large blob support was negotiated for this credential.
    Supported bool `json:"supported"`


    // Blob - Whether a large blob was read or written during the ceremony.
    Blob bool `json:"blob"`

}

// String returns a JSON representation of the model
func (o *Authenticationextensionslargebloboutputs) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authenticationextensionslargebloboutputs) MarshalJSON() ([]byte, error) {
    type Alias Authenticationextensionslargebloboutputs

    if AuthenticationextensionslargebloboutputsMarshalled {
        return []byte("{}"), nil
    }
    AuthenticationextensionslargebloboutputsMarshalled = true

    return json.Marshal(&struct {
        
        Supported bool `json:"supported"`
        
        Blob bool `json:"blob"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

