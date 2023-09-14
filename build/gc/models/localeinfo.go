package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LocaleinfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LocaleinfoDud struct { 
    


    


    

}

// Localeinfo
type Localeinfo struct { 
    // Status - Status of health computation for this flow version.
    Status string `json:"status"`


    // ErrorInfo - Error details for the flow version, if any.
    ErrorInfo Flowhealtherrorinfo `json:"errorInfo"`


    // FlowVersionInfo - Info about given flow version.
    FlowVersionInfo Localeflowversioninfo `json:"flowVersionInfo"`

}

// String returns a JSON representation of the model
func (o *Localeinfo) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Localeinfo) MarshalJSON() ([]byte, error) {
    type Alias Localeinfo

    if LocaleinfoMarshalled {
        return []byte("{}"), nil
    }
    LocaleinfoMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        ErrorInfo Flowhealtherrorinfo `json:"errorInfo"`
        
        FlowVersionInfo Localeflowversioninfo `json:"flowVersionInfo"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

