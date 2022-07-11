package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanreferenceDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Workplanreference
type Workplanreference struct { 
    


    // ManagementUnit - The management unit to which this work plan belongs.  Nullable in some routes
    ManagementUnit Managementunitreference `json:"managementUnit"`


    

}

// String returns a JSON representation of the model
func (o *Workplanreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanreference) MarshalJSON() ([]byte, error) {
    type Alias Workplanreference

    if WorkplanreferenceMarshalled {
        return []byte("{}"), nil
    }
    WorkplanreferenceMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

