package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatemanagementunitrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatemanagementunitrequestDud struct { 
    


    


    

}

// Updatemanagementunitrequest
type Updatemanagementunitrequest struct { 
    // Name - The new name of the management unit
    Name string `json:"name"`


    // DivisionId - The new division id for the management unit
    DivisionId string `json:"divisionId"`


    // Settings - Updated settings for the management unit
    Settings Managementunitsettingsrequest `json:"settings"`

}

// String returns a JSON representation of the model
func (o *Updatemanagementunitrequest) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatemanagementunitrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatemanagementunitrequest

    if UpdatemanagementunitrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatemanagementunitrequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DivisionId string `json:"divisionId"`
        
        Settings Managementunitsettingsrequest `json:"settings"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

