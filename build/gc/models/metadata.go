package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MetadataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MetadataDud struct { 
    


    


    

}

// Metadata
type Metadata struct { 
    // PairingToken
    PairingToken string `json:"pairing-token"`


    // PairingTrust
    PairingTrust []string `json:"pairing-trust"`


    // PairingUrl
    PairingUrl string `json:"pairing-url"`

}

// String returns a JSON representation of the model
func (o *Metadata) String() string {
    
     o.PairingTrust = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Metadata) MarshalJSON() ([]byte, error) {
    type Alias Metadata

    if MetadataMarshalled {
        return []byte("{}"), nil
    }
    MetadataMarshalled = true

    return json.Marshal(&struct {
        
        PairingToken string `json:"pairing-token"`
        
        PairingTrust []string `json:"pairing-trust"`
        
        PairingUrl string `json:"pairing-url"`
        *Alias
    }{

        


        
        PairingTrust: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

