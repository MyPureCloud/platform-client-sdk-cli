package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuqueryagentschedulesrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuqueryagentschedulesrequestDud struct { 
    


    


    

}

// Buqueryagentschedulesrequest
type Buqueryagentschedulesrequest struct { 
    // ManagementUnitId - The ID of the management unit to query
    ManagementUnitId string `json:"managementUnitId"`


    // UserIds - The IDs of the users to query.  Omit to query all user schedules in the management unit. Note: If teamIds is also specified, only schedules for users in the requested teams will be returned
    UserIds []string `json:"userIds"`


    // TeamIds - The teamIds to request. If null or not set, results will be queried for requested users if applicable or otherwise all users in the management unit
    TeamIds []string `json:"teamIds"`

}

// String returns a JSON representation of the model
func (o *Buqueryagentschedulesrequest) String() string {
    
     o.UserIds = []string{""} 
     o.TeamIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buqueryagentschedulesrequest) MarshalJSON() ([]byte, error) {
    type Alias Buqueryagentschedulesrequest

    if BuqueryagentschedulesrequestMarshalled {
        return []byte("{}"), nil
    }
    BuqueryagentschedulesrequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitId string `json:"managementUnitId"`
        
        UserIds []string `json:"userIds"`
        
        TeamIds []string `json:"teamIds"`
        *Alias
    }{

        


        
        UserIds: []string{""},
        


        
        TeamIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

