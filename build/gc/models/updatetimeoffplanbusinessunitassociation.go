package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatetimeoffplanbusinessunitassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatetimeoffplanbusinessunitassociationDud struct { 
    


    

}

// Updatetimeoffplanbusinessunitassociation
type Updatetimeoffplanbusinessunitassociation struct { 
    // ManagementUnitIds - The IDs of management units to which this time-off plan applies. This must not be set if staffingGroupIds is populated
    ManagementUnitIds Setwrapperstring `json:"managementUnitIds"`


    // StaffingGroupIds - The IDs of staffing groups to which this time-off plan applies. This must not be set if managementUnitIds is populated
    StaffingGroupIds Setwrapperstring `json:"staffingGroupIds"`

}

// String returns a JSON representation of the model
func (o *Updatetimeoffplanbusinessunitassociation) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatetimeoffplanbusinessunitassociation) MarshalJSON() ([]byte, error) {
    type Alias Updatetimeoffplanbusinessunitassociation

    if UpdatetimeoffplanbusinessunitassociationMarshalled {
        return []byte("{}"), nil
    }
    UpdatetimeoffplanbusinessunitassociationMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitIds Setwrapperstring `json:"managementUnitIds"`
        
        StaffingGroupIds Setwrapperstring `json:"staffingGroupIds"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

