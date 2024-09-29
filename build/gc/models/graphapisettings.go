package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GraphapisettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GraphapisettingsDud struct { 
    


    Status string `json:"status"`


    ErrorInfo Emailerrorinfo `json:"errorInfo"`

}

// Graphapisettings
type Graphapisettings struct { 
    // Integration - The Graph API server integration to use for emails.
    Integration Domainentityref `json:"integration"`


    


    

}

// String returns a JSON representation of the model
func (o *Graphapisettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Graphapisettings) MarshalJSON() ([]byte, error) {
    type Alias Graphapisettings

    if GraphapisettingsMarshalled {
        return []byte("{}"), nil
    }
    GraphapisettingsMarshalled = true

    return json.Marshal(&struct {
        
        Integration Domainentityref `json:"integration"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

