package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AutosearchconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AutosearchconfigDud struct { 
    

}

// Autosearchconfig
type Autosearchconfig struct { 
    // VarType - Auto search configuration type.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Autosearchconfig) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Autosearchconfig) MarshalJSON() ([]byte, error) {
    type Alias Autosearchconfig

    if AutosearchconfigMarshalled {
        return []byte("{}"), nil
    }
    AutosearchconfigMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

