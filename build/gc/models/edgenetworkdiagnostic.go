package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgenetworkdiagnosticMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgenetworkdiagnosticDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Edgenetworkdiagnostic
type Edgenetworkdiagnostic struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Edgenetworkdiagnostic) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgenetworkdiagnostic) MarshalJSON() ([]byte, error) {
    type Alias Edgenetworkdiagnostic

    if EdgenetworkdiagnosticMarshalled {
        return []byte("{}"), nil
    }
    EdgenetworkdiagnosticMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

