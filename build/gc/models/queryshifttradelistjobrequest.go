package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryshifttradelistjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryshifttradelistjobrequestDud struct { 
    


    


    

}

// Queryshifttradelistjobrequest
type Queryshifttradelistjobrequest struct { 
    // ManagementUnitIds - The IDs of management units from which to query shift trades
    ManagementUnitIds []string `json:"managementUnitIds"`


    // WeekDates - The start week dates in which to query shift trades in the business unit time zone (yyyy-MM-dd format)
    WeekDates []time.Time `json:"weekDates"`


    // UserIds - The IDs of the users for whom to query trades
    UserIds []string `json:"userIds"`

}

// String returns a JSON representation of the model
func (o *Queryshifttradelistjobrequest) String() string {
     o.ManagementUnitIds = []string{""} 
     o.WeekDates = []time.Time{{}} 
     o.UserIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryshifttradelistjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryshifttradelistjobrequest

    if QueryshifttradelistjobrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryshifttradelistjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        WeekDates []time.Time `json:"weekDates"`
        
        UserIds []string `json:"userIds"`
        *Alias
    }{

        
        ManagementUnitIds: []string{""},
        


        
        WeekDates: []time.Time{{}},
        


        
        UserIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

