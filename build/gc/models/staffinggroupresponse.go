package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffinggroupresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffinggroupresponseDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Staffinggroupresponse
type Staffinggroupresponse struct { 
    


    // Name - The name of the staffing group
    Name string `json:"name"`


    // Users - The list of users that belong to the staffing group
    Users []Userreference `json:"users"`


    // ManagementUnit - The ID of the management unit to which the staffing group users belong. If undefined the staffing group can include users from the entire business unit
    ManagementUnit Managementunitreference `json:"managementUnit"`


    // PlanningGroups - The list of planning groups that are associated with the staffing group
    PlanningGroups []Planninggroupreference `json:"planningGroups"`


    // Metadata - Version metadata for the staffing group
    Metadata Wfmversionedentitymetadata `json:"metadata"`


    

}

// String returns a JSON representation of the model
func (o *Staffinggroupresponse) String() string {
    
     o.Users = []Userreference{{}} 
    
     o.PlanningGroups = []Planninggroupreference{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffinggroupresponse) MarshalJSON() ([]byte, error) {
    type Alias Staffinggroupresponse

    if StaffinggroupresponseMarshalled {
        return []byte("{}"), nil
    }
    StaffinggroupresponseMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Users []Userreference `json:"users"`
        
        ManagementUnit Managementunitreference `json:"managementUnit"`
        
        PlanningGroups []Planninggroupreference `json:"planningGroups"`
        
        Metadata Wfmversionedentitymetadata `json:"metadata"`
        *Alias
    }{

        


        


        
        Users: []Userreference{{}},
        


        


        
        PlanningGroups: []Planninggroupreference{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

