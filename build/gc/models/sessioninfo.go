package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessioninfoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessioninfoDud struct { 
    


    


    


    


    

}

// Sessioninfo
type Sessioninfo struct { 
    // Version - Version of the continuous forecast session
    Version int `json:"version"`


    // SessionId - Session ID of the continuous forecast session
    SessionId string `json:"sessionId"`


    // BusinessUnitId - Business unit ID of the continuous forecast session
    BusinessUnitId string `json:"businessUnitId"`


    // PlanningGroupsVersion - Version of the planning groups
    PlanningGroupsVersion int `json:"planningGroupsVersion"`


    // DateOfSession - Date and Time of the Session. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateOfSession time.Time `json:"dateOfSession"`

}

// String returns a JSON representation of the model
func (o *Sessioninfo) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessioninfo) MarshalJSON() ([]byte, error) {
    type Alias Sessioninfo

    if SessioninfoMarshalled {
        return []byte("{}"), nil
    }
    SessioninfoMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        SessionId string `json:"sessionId"`
        
        BusinessUnitId string `json:"businessUnitId"`
        
        PlanningGroupsVersion int `json:"planningGroupsVersion"`
        
        DateOfSession time.Time `json:"dateOfSession"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

