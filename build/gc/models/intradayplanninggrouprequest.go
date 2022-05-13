package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntradayplanninggrouprequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntradayplanninggrouprequestDud struct { 
    


    


    


    

}

// Intradayplanninggrouprequest
type Intradayplanninggrouprequest struct { 
    // BusinessUnitDate - Requested date in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    BusinessUnitDate time.Time `json:"businessUnitDate"`


    // Categories - The metric categories
    Categories []string `json:"categories"`


    // PlanningGroupIds - The IDs of the planning groups for which to fetch data.  Omitting or passing an empty list will return all available planning groups
    PlanningGroupIds []string `json:"planningGroupIds"`


    // IntervalLengthMinutes - The period/interval in minutes for which to aggregate the data. Required, defaults to 15
    IntervalLengthMinutes int `json:"intervalLengthMinutes"`

}

// String returns a JSON representation of the model
func (o *Intradayplanninggrouprequest) String() string {
    
     o.Categories = []string{""} 
     o.PlanningGroupIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intradayplanninggrouprequest) MarshalJSON() ([]byte, error) {
    type Alias Intradayplanninggrouprequest

    if IntradayplanninggrouprequestMarshalled {
        return []byte("{}"), nil
    }
    IntradayplanninggrouprequestMarshalled = true

    return json.Marshal(&struct {
        
        BusinessUnitDate time.Time `json:"businessUnitDate"`
        
        Categories []string `json:"categories"`
        
        PlanningGroupIds []string `json:"planningGroupIds"`
        
        IntervalLengthMinutes int `json:"intervalLengthMinutes"`
        *Alias
    }{

        


        
        Categories: []string{""},
        


        
        PlanningGroupIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

