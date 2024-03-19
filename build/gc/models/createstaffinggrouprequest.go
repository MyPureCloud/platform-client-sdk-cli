package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatestaffinggrouprequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatestaffinggrouprequestDud struct { 
    


    


    

}

// Createstaffinggrouprequest
type Createstaffinggrouprequest struct { 
    // Name - The name of the staffing group
    Name string `json:"name"`


    // UserIds - The set of user IDs to associate with the staffing group
    UserIds []string `json:"userIds"`


    // ManagementUnitId - The ID of the management unit to which the staffing group users belong. If undefined the staffing group can include users from the entire business unit
    ManagementUnitId string `json:"managementUnitId"`

}

// String returns a JSON representation of the model
func (o *Createstaffinggrouprequest) String() string {
    
     o.UserIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createstaffinggrouprequest) MarshalJSON() ([]byte, error) {
    type Alias Createstaffinggrouprequest

    if CreatestaffinggrouprequestMarshalled {
        return []byte("{}"), nil
    }
    CreatestaffinggrouprequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        UserIds []string `json:"userIds"`
        
        ManagementUnitId string `json:"managementUnitId"`
        *Alias
    }{

        


        
        UserIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

