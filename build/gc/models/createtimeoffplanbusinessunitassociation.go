package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatetimeoffplanbusinessunitassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatetimeoffplanbusinessunitassociationDud struct { 
    


    

}

// Createtimeoffplanbusinessunitassociation
type Createtimeoffplanbusinessunitassociation struct { 
    // ManagementUnitIds - The IDs of management units to which this time-off plan applies. This must not be set if staffingGroupIds is populated
    ManagementUnitIds []string `json:"managementUnitIds"`


    // StaffingGroupIds - The IDs of staffing groups to which this time-off plan applies. This must not be set if managementUnitIds is populated
    StaffingGroupIds []string `json:"staffingGroupIds"`

}

// String returns a JSON representation of the model
func (o *Createtimeoffplanbusinessunitassociation) String() string {
     o.ManagementUnitIds = []string{""} 
     o.StaffingGroupIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createtimeoffplanbusinessunitassociation) MarshalJSON() ([]byte, error) {
    type Alias Createtimeoffplanbusinessunitassociation

    if CreatetimeoffplanbusinessunitassociationMarshalled {
        return []byte("{}"), nil
    }
    CreatetimeoffplanbusinessunitassociationMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        StaffingGroupIds []string `json:"staffingGroupIds"`
        *Alias
    }{

        
        ManagementUnitIds: []string{""},
        


        
        StaffingGroupIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

