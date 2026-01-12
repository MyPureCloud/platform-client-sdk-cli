package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RichlinkheaderMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RichlinkheaderDud struct { 
    


    


    

}

// Richlinkheader
type Richlinkheader struct { 
    // VarType - Describes the Rich Link header type.
    VarType string `json:"type"`


    // Value - Rich Link header text value.
    Value string `json:"value"`


    // Url - Rich Link header URL.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Richlinkheader) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Richlinkheader) MarshalJSON() ([]byte, error) {
    type Alias Richlinkheader

    if RichlinkheaderMarshalled {
        return []byte("{}"), nil
    }
    RichlinkheaderMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

