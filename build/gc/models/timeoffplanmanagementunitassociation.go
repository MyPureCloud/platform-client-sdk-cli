package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffplanmanagementunitassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffplanmanagementunitassociationDud struct { 
    


    

}

// Timeoffplanmanagementunitassociation
type Timeoffplanmanagementunitassociation struct { 
    // ManagementUnit - Management unit to which this time-off plan belongs
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // StaffingGroups - Staffing groups to which this time-off plan applies. If not defined, the plan applies to the management unit
    StaffingGroups []Staffinggroupreference `json:"staffingGroups"`

}

// String returns a JSON representation of the model
func (o *Timeoffplanmanagementunitassociation) String() string {
    
     o.StaffingGroups = []Staffinggroupreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffplanmanagementunitassociation) MarshalJSON() ([]byte, error) {
    type Alias Timeoffplanmanagementunitassociation

    if TimeoffplanmanagementunitassociationMarshalled {
        return []byte("{}"), nil
    }
    TimeoffplanmanagementunitassociationMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        StaffingGroups []Staffinggroupreference `json:"staffingGroups"`
        *Alias
    }{

        


        
        StaffingGroups: []Staffinggroupreference{{}},
        

        Alias: (*Alias)(u),
    })
}

