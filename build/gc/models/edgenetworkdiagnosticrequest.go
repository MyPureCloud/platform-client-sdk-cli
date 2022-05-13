package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EdgenetworkdiagnosticrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EdgenetworkdiagnosticrequestDud struct { 
    

}

// Edgenetworkdiagnosticrequest
type Edgenetworkdiagnosticrequest struct { 
    // Host - IPv4/6 address or host to be probed for connectivity. No port allowed.
    Host string `json:"host"`

}

// String returns a JSON representation of the model
func (o *Edgenetworkdiagnosticrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Edgenetworkdiagnosticrequest) MarshalJSON() ([]byte, error) {
    type Alias Edgenetworkdiagnosticrequest

    if EdgenetworkdiagnosticrequestMarshalled {
        return []byte("{}"), nil
    }
    EdgenetworkdiagnosticrequestMarshalled = true

    return json.Marshal(&struct {
        
        Host string `json:"host"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

