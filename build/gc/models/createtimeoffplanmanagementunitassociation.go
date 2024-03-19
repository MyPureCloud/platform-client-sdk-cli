package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatetimeoffplanmanagementunitassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatetimeoffplanmanagementunitassociationDud struct { 
    


    

}

// Createtimeoffplanmanagementunitassociation
type Createtimeoffplanmanagementunitassociation struct { 
    // ManagementUnitId - The ID of the management unit to which this time-off plan belongs
    ManagementUnitId string `json:"managementUnitId"`


    // StaffingGroupIds - A IDs of staffing groups to which this time-off plan applies. If not defined, the plan is applied to the management unit
    StaffingGroupIds []string `json:"staffingGroupIds"`

}

// String returns a JSON representation of the model
func (o *Createtimeoffplanmanagementunitassociation) String() string {
    
     o.StaffingGroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createtimeoffplanmanagementunitassociation) MarshalJSON() ([]byte, error) {
    type Alias Createtimeoffplanmanagementunitassociation

    if CreatetimeoffplanmanagementunitassociationMarshalled {
        return []byte("{}"), nil
    }
    CreatetimeoffplanmanagementunitassociationMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitId string `json:"managementUnitId"`
        
        StaffingGroupIds []string `json:"staffingGroupIds"`
        *Alias
    }{

        


        
        StaffingGroupIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

