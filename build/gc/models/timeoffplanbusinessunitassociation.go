package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffplanbusinessunitassociationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffplanbusinessunitassociationDud struct { 
    


    

}

// Timeoffplanbusinessunitassociation
type Timeoffplanbusinessunitassociation struct { 
    // ManagementUnits - Management units to which this time-off plan applies. This must not be set if staffingGroups is populated
    ManagementUnits []Managementunitreference `json:"managementUnits"`


    // StaffingGroups - Staffing groups to which this time-off plan applies. This must not be set if managementUnits is populated
    StaffingGroups []Staffinggroupreference `json:"staffingGroups"`

}

// String returns a JSON representation of the model
func (o *Timeoffplanbusinessunitassociation) String() string {
     o.ManagementUnits = []Managementunitreference{{}} 
     o.StaffingGroups = []Staffinggroupreference{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffplanbusinessunitassociation) MarshalJSON() ([]byte, error) {
    type Alias Timeoffplanbusinessunitassociation

    if TimeoffplanbusinessunitassociationMarshalled {
        return []byte("{}"), nil
    }
    TimeoffplanbusinessunitassociationMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnits []Managementunitreference `json:"managementUnits"`
        
        StaffingGroups []Staffinggroupreference `json:"staffingGroups"`
        *Alias
    }{

        
        ManagementUnits: []Managementunitreference{{}},
        


        
        StaffingGroups: []Staffinggroupreference{{}},
        

        Alias: (*Alias)(u),
    })
}

