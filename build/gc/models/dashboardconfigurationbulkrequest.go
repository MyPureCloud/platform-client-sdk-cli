package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DashboardconfigurationbulkrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DashboardconfigurationbulkrequestDud struct { 
    

}

// Dashboardconfigurationbulkrequest
type Dashboardconfigurationbulkrequest struct { 
    // DashboardConfigurationIds - The user supplied dashboard configuration ids
    DashboardConfigurationIds []string `json:"dashboardConfigurationIds"`

}

// String returns a JSON representation of the model
func (o *Dashboardconfigurationbulkrequest) String() string {
     o.DashboardConfigurationIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dashboardconfigurationbulkrequest) MarshalJSON() ([]byte, error) {
    type Alias Dashboardconfigurationbulkrequest

    if DashboardconfigurationbulkrequestMarshalled {
        return []byte("{}"), nil
    }
    DashboardconfigurationbulkrequestMarshalled = true

    return json.Marshal(&struct {
        
        DashboardConfigurationIds []string `json:"dashboardConfigurationIds"`
        *Alias
    }{

        
        DashboardConfigurationIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

