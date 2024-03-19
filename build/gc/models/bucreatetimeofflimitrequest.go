package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BucreatetimeofflimitrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BucreatetimeofflimitrequestDud struct { 
    


    

}

// Bucreatetimeofflimitrequest
type Bucreatetimeofflimitrequest struct { 
    // StaffingGroupId - The ID of the staffing group to which this time-off limit is associated. It can be either management unit or business unit level staffing group. One of managementUnitId or staffingGroupId must be set. This must not be set if managementUnitId has value
    StaffingGroupId string `json:"staffingGroupId"`


    // ManagementUnitId - The ID of the management unit to which this time-off limit is associated. One of managementUnitId or staffingGroupId must be set. This must not be set if staffingGroupId has value
    ManagementUnitId string `json:"managementUnitId"`

}

// String returns a JSON representation of the model
func (o *Bucreatetimeofflimitrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bucreatetimeofflimitrequest) MarshalJSON() ([]byte, error) {
    type Alias Bucreatetimeofflimitrequest

    if BucreatetimeofflimitrequestMarshalled {
        return []byte("{}"), nil
    }
    BucreatetimeofflimitrequestMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroupId string `json:"staffingGroupId"`
        
        ManagementUnitId string `json:"managementUnitId"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

