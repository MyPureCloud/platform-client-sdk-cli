package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ExportscriptresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ExportscriptresponseDud struct { 
    

}

// Exportscriptresponse
type Exportscriptresponse struct { 
    // Url
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Exportscriptresponse) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Exportscriptresponse) MarshalJSON() ([]byte, error) {
    type Alias Exportscriptresponse

    if ExportscriptresponseMarshalled {
        return []byte("{}"), nil
    }
    ExportscriptresponseMarshalled = true

    return json.Marshal(&struct { 
        Url string `json:"url"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

