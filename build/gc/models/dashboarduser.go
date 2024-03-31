package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DashboarduserMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DashboarduserDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dashboarduser
type Dashboarduser struct { 
    


    // Name
    Name string `json:"name"`


    // DashboardCount - The count of dashboards for the user
    DashboardCount int `json:"dashboardCount"`


    // PublicDashboardCount - The count of public dashboards for the user
    PublicDashboardCount int `json:"publicDashboardCount"`


    // State - The state of the user
    State string `json:"state"`


    

}

// String returns a JSON representation of the model
func (o *Dashboarduser) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dashboarduser) MarshalJSON() ([]byte, error) {
    type Alias Dashboarduser

    if DashboarduserMarshalled {
        return []byte("{}"), nil
    }
    DashboarduserMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        DashboardCount int `json:"dashboardCount"`
        
        PublicDashboardCount int `json:"publicDashboardCount"`
        
        State string `json:"state"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

