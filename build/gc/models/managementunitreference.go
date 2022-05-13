package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ManagementunitreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ManagementunitreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Managementunitreference - Management unit reference object for Workforce Management (ID/selfUri only)
type Managementunitreference struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Managementunitreference) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Managementunitreference) MarshalJSON() ([]byte, error) {
    type Alias Managementunitreference

    if ManagementunitreferenceMarshalled {
        return []byte("{}"), nil
    }
    ManagementunitreferenceMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

