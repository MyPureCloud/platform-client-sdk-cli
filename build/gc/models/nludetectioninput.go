package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludetectioninputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludetectioninputDud struct { 
    


    

}

// Nludetectioninput
type Nludetectioninput struct { 
    // Text - The text to perform NLU detection on.
    Text string `json:"text"`


    // Language - Language of the version for multilingual detection, e.g. `en-us`, `de-de`
    Language string `json:"language"`

}

// String returns a JSON representation of the model
func (o *Nludetectioninput) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludetectioninput) MarshalJSON() ([]byte, error) {
    type Alias Nludetectioninput

    if NludetectioninputMarshalled {
        return []byte("{}"), nil
    }
    NludetectioninputMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Language string `json:"language"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

