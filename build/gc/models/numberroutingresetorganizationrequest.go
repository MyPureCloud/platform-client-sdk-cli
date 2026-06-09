package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NumberroutingresetorganizationrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NumberroutingresetorganizationrequestDud struct { 
    

}

// Numberroutingresetorganizationrequest - Number Routing reset routing request body
type Numberroutingresetorganizationrequest struct { 
    // ResetOrganizationId - Organization Id where all rerouted numbers will be reset to
    ResetOrganizationId string `json:"resetOrganizationId"`

}

// String returns a JSON representation of the model
func (o *Numberroutingresetorganizationrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Numberroutingresetorganizationrequest) MarshalJSON() ([]byte, error) {
    type Alias Numberroutingresetorganizationrequest

    if NumberroutingresetorganizationrequestMarshalled {
        return []byte("{}"), nil
    }
    NumberroutingresetorganizationrequestMarshalled = true

    return json.Marshal(&struct {
        
        ResetOrganizationId string `json:"resetOrganizationId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

