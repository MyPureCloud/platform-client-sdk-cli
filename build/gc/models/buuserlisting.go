package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuuserlistingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuuserlistingDud struct { 
    

}

// Buuserlisting
type Buuserlisting struct { 
    // ManagementUnits - Management units and their associated users
    ManagementUnits []Managementunituserlisting `json:"managementUnits"`

}

// String returns a JSON representation of the model
func (o *Buuserlisting) String() string {
     o.ManagementUnits = []Managementunituserlisting{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buuserlisting) MarshalJSON() ([]byte, error) {
    type Alias Buuserlisting

    if BuuserlistingMarshalled {
        return []byte("{}"), nil
    }
    BuuserlistingMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnits []Managementunituserlisting `json:"managementUnits"`
        *Alias
    }{

        
        ManagementUnits: []Managementunituserlisting{{}},
        

        Alias: (*Alias)(u),
    })
}

