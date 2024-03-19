package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatetimeoffplanmanagementunitassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatetimeoffplanmanagementunitassociationDud struct { 
    

}

// Updatetimeoffplanmanagementunitassociation
type Updatetimeoffplanmanagementunitassociation struct { 
    // StaffingGroupIds - The IDs of staffing groups to which this time-off plan applies. If not defined, the plan is applied to the management unit
    StaffingGroupIds Setwrapperstring `json:"staffingGroupIds"`

}

// String returns a JSON representation of the model
func (o *Updatetimeoffplanmanagementunitassociation) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatetimeoffplanmanagementunitassociation) MarshalJSON() ([]byte, error) {
    type Alias Updatetimeoffplanmanagementunitassociation

    if UpdatetimeoffplanmanagementunitassociationMarshalled {
        return []byte("{}"), nil
    }
    UpdatetimeoffplanmanagementunitassociationMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroupIds Setwrapperstring `json:"staffingGroupIds"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

