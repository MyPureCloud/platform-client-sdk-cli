package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcentercompactcategorymoduletemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcentercompactcategorymoduletemplateDud struct { 
    

}

// Supportcentercompactcategorymoduletemplate
type Supportcentercompactcategorymoduletemplate struct { 
    // Active - Whether this template is active or not
    Active bool `json:"active"`

}

// String returns a JSON representation of the model
func (o *Supportcentercompactcategorymoduletemplate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcentercompactcategorymoduletemplate) MarshalJSON() ([]byte, error) {
    type Alias Supportcentercompactcategorymoduletemplate

    if SupportcentercompactcategorymoduletemplateMarshalled {
        return []byte("{}"), nil
    }
    SupportcentercompactcategorymoduletemplateMarshalled = true

    return json.Marshal(&struct {
        
        Active bool `json:"active"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

