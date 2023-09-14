package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FlowhealthintentversioninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FlowhealthintentversioninfoDud struct { 
    NluVersion Addressableentityref `json:"nluVersion"`


    FlowVersion Addressableentityref `json:"flowVersion"`


    NluDomain Addressableentityref `json:"nluDomain"`

}

// Flowhealthintentversioninfo
type Flowhealthintentversioninfo struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Flowhealthintentversioninfo) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Flowhealthintentversioninfo) MarshalJSON() ([]byte, error) {
    type Alias Flowhealthintentversioninfo

    if FlowhealthintentversioninfoMarshalled {
        return []byte("{}"), nil
    }
    FlowhealthintentversioninfoMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

