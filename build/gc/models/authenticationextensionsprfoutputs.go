package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthenticationextensionsprfoutputsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthenticationextensionsprfoutputsDud struct { 
    


    

}

// Authenticationextensionsprfoutputs
type Authenticationextensionsprfoutputs struct { 
    // Enabled - Whether the PRF extension was enabled for this ceremony.
    Enabled bool `json:"enabled"`


    // Results - PRF evaluation results keyed by input name (values are base64url-encoded octets).
    Results map[string]string `json:"results"`

}

// String returns a JSON representation of the model
func (o *Authenticationextensionsprfoutputs) String() string {
    
     o.Results = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authenticationextensionsprfoutputs) MarshalJSON() ([]byte, error) {
    type Alias Authenticationextensionsprfoutputs

    if AuthenticationextensionsprfoutputsMarshalled {
        return []byte("{}"), nil
    }
    AuthenticationextensionsprfoutputsMarshalled = true

    return json.Marshal(&struct {
        
        Enabled bool `json:"enabled"`
        
        Results map[string]string `json:"results"`
        *Alias
    }{

        


        
        Results: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

