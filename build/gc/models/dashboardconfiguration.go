package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DashboardconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DashboardconfigurationDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    Restricted bool `json:"restricted"`


    


    


    


    CreatedBy Addressableentityref `json:"createdBy"`


    


    


    SelfUri string `json:"selfUri"`

}

// Dashboardconfiguration
type Dashboardconfiguration struct { 
    


    // Name - The name of dashboard configuration.
    Name string `json:"name"`


    // Rows - The count of rows for the specific dashboard configuration.
    Rows int `json:"rows"`


    // Columns - The count of columns for the specific dashboard.
    Columns int `json:"columns"`


    // Widgets - List of widgets for dashboard configuration.
    Widgets []Widget `json:"widgets"`


    // Favorite - The flag indicates if the dashboard is favorited by the user
    Favorite bool `json:"favorite"`


    // PublicDashboard - The flag to indicate if the dashboard is published by an user
    PublicDashboard bool `json:"publicDashboard"`


    


    // LayoutType - The layout type of the dashboard
    LayoutType string `json:"layoutType"`


    // DateCreated - The created date of the dashboard. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // DateModified - The last modified date of the dashboard. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateModified time.Time `json:"dateModified"`


    


    // Shared - The flag to indicate if the dashboard is shared
    Shared bool `json:"shared"`


    // DashboardsSharedWith - The list of users and teams the dashboard is shared with
    DashboardsSharedWith Dashboardssharedwith `json:"dashboardsSharedWith"`


    

}

// String returns a JSON representation of the model
func (o *Dashboardconfiguration) String() string {
    
    
    
     o.Widgets = []Widget{{}} 
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dashboardconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Dashboardconfiguration

    if DashboardconfigurationMarshalled {
        return []byte("{}"), nil
    }
    DashboardconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Rows int `json:"rows"`
        
        Columns int `json:"columns"`
        
        Widgets []Widget `json:"widgets"`
        
        Favorite bool `json:"favorite"`
        
        PublicDashboard bool `json:"publicDashboard"`
        
        LayoutType string `json:"layoutType"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        DateModified time.Time `json:"dateModified"`
        
        Shared bool `json:"shared"`
        
        DashboardsSharedWith Dashboardssharedwith `json:"dashboardsSharedWith"`
        *Alias
    }{

        


        


        


        


        
        Widgets: []Widget{{}},
        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

