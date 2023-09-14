package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowhealthversioninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowhealthversioninfoDud struct { 
    FlowVersion Addressableentityref `json:"flowVersion"`


    NluDomain Addressableentityref `json:"nluDomain"`

}

// Flowhealthversioninfo
type Flowhealthversioninfo struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Flowhealthversioninfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowhealthversioninfo) MarshalJSON() ([]byte, error) {
    type Alias Flowhealthversioninfo

    if FlowhealthversioninfoMarshalled {
        return []byte("{}"), nil
    }
    FlowhealthversioninfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

