package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgeversioninformationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgeversioninformationDud struct { 
    

}

// Edgeversioninformation
type Edgeversioninformation struct { 
    // SoftwareVersion
    SoftwareVersion string `json:"softwareVersion"`

}

// String returns a JSON representation of the model
func (o *Edgeversioninformation) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgeversioninformation) MarshalJSON() ([]byte, error) {
    type Alias Edgeversioninformation

    if EdgeversioninformationMarshalled {
        return []byte("{}"), nil
    }
    EdgeversioninformationMarshalled = true

    return json.Marshal(&struct { 
        SoftwareVersion string `json:"softwareVersion"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

