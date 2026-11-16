package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ClientextensionresultsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ClientextensionresultsDud struct { 
    


    


    


    


    

}

// Clientextensionresults
type Clientextensionresults struct { 
    // Appid - App ID extension output (legacy U2F migration).
    Appid bool `json:"appid"`


    // CredProps - Credential properties extension output.
    CredProps Credentialpropertiesoutput `json:"credProps"`


    // HmacCreateSecret - HMAC secret extension output.
    HmacCreateSecret bool `json:"hmacCreateSecret"`


    // LargeBlob - Large blob extension outputs.
    LargeBlob Authenticationextensionslargebloboutputs `json:"largeBlob"`


    // Prf - Pseudo-random function (PRF) extension outputs.
    Prf Authenticationextensionsprfoutputs `json:"prf"`

}

// String returns a JSON representation of the model
func (o *Clientextensionresults) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Clientextensionresults) MarshalJSON() ([]byte, error) {
    type Alias Clientextensionresults

    if ClientextensionresultsMarshalled {
        return []byte("{}"), nil
    }
    ClientextensionresultsMarshalled = true

    return json.Marshal(&struct {
        
        Appid bool `json:"appid"`
        
        CredProps Credentialpropertiesoutput `json:"credProps"`
        
        HmacCreateSecret bool `json:"hmacCreateSecret"`
        
        LargeBlob Authenticationextensionslargebloboutputs `json:"largeBlob"`
        
        Prf Authenticationextensionsprfoutputs `json:"prf"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

