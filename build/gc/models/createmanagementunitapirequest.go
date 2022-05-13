package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatemanagementunitapirequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatemanagementunitapirequestDud struct { 
    


    


    


    


    


    

}

// Createmanagementunitapirequest - Create Management Unit
type Createmanagementunitapirequest struct { 
    // Name - The name of the management unit
    Name string `json:"name"`


    // TimeZone - The default time zone to use for this management unit.  Moving to Business Unit
    TimeZone string `json:"timeZone"`


    // StartDayOfWeek - The configured first day of the week for scheduling and forecasting purposes. Moving to Business Unit
    StartDayOfWeek string `json:"startDayOfWeek"`


    // Settings - The configuration for the management unit.  If omitted, reasonable defaults will be assigned
    Settings Createmanagementunitsettingsrequest `json:"settings"`


    // DivisionId - The id of the division to which this management unit belongs.  Defaults to home division ID
    DivisionId string `json:"divisionId"`


    // BusinessUnitId - The id of the business unit to which this management unit belongs
    BusinessUnitId string `json:"businessUnitId"`

}

// String returns a JSON representation of the model
func (o *Createmanagementunitapirequest) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createmanagementunitapirequest) MarshalJSON() ([]byte, error) {
    type Alias Createmanagementunitapirequest

    if CreatemanagementunitapirequestMarshalled {
        return []byte("{}"), nil
    }
    CreatemanagementunitapirequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        TimeZone string `json:"timeZone"`
        
        StartDayOfWeek string `json:"startDayOfWeek"`
        
        Settings Createmanagementunitsettingsrequest `json:"settings"`
        
        DivisionId string `json:"divisionId"`
        
        BusinessUnitId string `json:"businessUnitId"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

