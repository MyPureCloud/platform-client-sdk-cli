package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DashboardssharedwithMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DashboardssharedwithDud struct { 
    


    

}

// Dashboardssharedwith
type Dashboardssharedwith struct { 
    // UserIds - The list of user ids to share the dashboard with
    UserIds []string `json:"userIds"`


    // TeamIds - The list of team ids to share the dashboard with
    TeamIds []string `json:"teamIds"`

}

// String returns a JSON representation of the model
func (o *Dashboardssharedwith) String() string {
     o.UserIds = []string{""} 
     o.TeamIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dashboardssharedwith) MarshalJSON() ([]byte, error) {
    type Alias Dashboardssharedwith

    if DashboardssharedwithMarshalled {
        return []byte("{}"), nil
    }
    DashboardssharedwithMarshalled = true

    return json.Marshal(&struct {
        
        UserIds []string `json:"userIds"`
        
        TeamIds []string `json:"teamIds"`
        *Alias
    }{

        
        UserIds: []string{""},
        


        
        TeamIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

