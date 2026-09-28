package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuqueryadherenceadjustmentsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuqueryadherenceadjustmentsrequestDud struct { 
    


    


    


    


    


    

}

// Buqueryadherenceadjustmentsrequest
type Buqueryadherenceadjustmentsrequest struct { 
    // StartDate - The start timestamp of the range to query in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // EndDate - The end timestamp of the range to query in ISO-8601 format
    EndDate time.Time `json:"endDate"`


    // ReasonCodeIds - A filter for the reason codes to include. Leave empty or omit entirely for all reason codes
    ReasonCodeIds []string `json:"reasonCodeIds"`


    // Statuses - A filter for which adherence adjustment statuses to include. Leave empty or omit entirely for all statuses
    Statuses []string `json:"statuses"`


    // UserIds - A filter for which users within the business unit to query. Leave empty or omit entirely for all users
    UserIds []string `json:"userIds"`


    // ManagementUnitIds - A filter for which management units to query. Leave empty or omit entirely for all management units in the business unit
    ManagementUnitIds []string `json:"managementUnitIds"`

}

// String returns a JSON representation of the model
func (o *Buqueryadherenceadjustmentsrequest) String() string {
    
    
     o.ReasonCodeIds = []string{""} 
     o.Statuses = []string{""} 
     o.UserIds = []string{""} 
     o.ManagementUnitIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buqueryadherenceadjustmentsrequest) MarshalJSON() ([]byte, error) {
    type Alias Buqueryadherenceadjustmentsrequest

    if BuqueryadherenceadjustmentsrequestMarshalled {
        return []byte("{}"), nil
    }
    BuqueryadherenceadjustmentsrequestMarshalled = true

    return json.Marshal(&struct {
        
        StartDate time.Time `json:"startDate"`
        
        EndDate time.Time `json:"endDate"`
        
        ReasonCodeIds []string `json:"reasonCodeIds"`
        
        Statuses []string `json:"statuses"`
        
        UserIds []string `json:"userIds"`
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        *Alias
    }{

        


        


        
        ReasonCodeIds: []string{""},
        


        
        Statuses: []string{""},
        


        
        UserIds: []string{""},
        


        
        ManagementUnitIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

