package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DashboardconfigurationqueryrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DashboardconfigurationqueryrequestDud struct { 
    


    


    


    

}

// Dashboardconfigurationqueryrequest
type Dashboardconfigurationqueryrequest struct { 
    // DashboardConfigurationIds - The user supplied dashboard configuration ids
    DashboardConfigurationIds []string `json:"dashboardConfigurationIds"`


    // PageNumber - The page number of the queried response
    PageNumber int `json:"pageNumber"`


    // PageSize - The number of entities to return of the queried response. The max is 25
    PageSize int `json:"pageSize"`


    // SortBy - The order in which response will be sorted
    SortBy string `json:"sortBy"`

}

// String returns a JSON representation of the model
func (o *Dashboardconfigurationqueryrequest) String() string {
     o.DashboardConfigurationIds = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dashboardconfigurationqueryrequest) MarshalJSON() ([]byte, error) {
    type Alias Dashboardconfigurationqueryrequest

    if DashboardconfigurationqueryrequestMarshalled {
        return []byte("{}"), nil
    }
    DashboardconfigurationqueryrequestMarshalled = true

    return json.Marshal(&struct {
        
        DashboardConfigurationIds []string `json:"dashboardConfigurationIds"`
        
        PageNumber int `json:"pageNumber"`
        
        PageSize int `json:"pageSize"`
        
        SortBy string `json:"sortBy"`
        *Alias
    }{

        
        DashboardConfigurationIds: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

