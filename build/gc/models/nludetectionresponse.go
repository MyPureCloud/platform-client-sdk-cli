package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludetectionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludetectionresponseDud struct { 
    Version Nludomainversion `json:"version"`


    


    

}

// Nludetectionresponse
type Nludetectionresponse struct { 
    


    // Output
    Output Nludetectionoutput `json:"output"`


    // Input
    Input Nludetectioninput `json:"input"`

}

// String returns a JSON representation of the model
func (o *Nludetectionresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludetectionresponse) MarshalJSON() ([]byte, error) {
    type Alias Nludetectionresponse

    if NludetectionresponseMarshalled {
        return []byte("{}"), nil
    }
    NludetectionresponseMarshalled = true

    return json.Marshal(&struct {
        
        Output Nludetectionoutput `json:"output"`
        
        Input Nludetectioninput `json:"input"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

