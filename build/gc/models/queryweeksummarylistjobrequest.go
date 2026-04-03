package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryweeksummarylistjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryweeksummarylistjobrequestDud struct { 
    


    

}

// Queryweeksummarylistjobrequest
type Queryweeksummarylistjobrequest struct { 
    // ManagementUnitIds - The IDs of management units from which to query shift trades
    ManagementUnitIds []string `json:"managementUnitIds"`


    // WeekDates - The start week dates in which to query shift trades in the business unit time zone (yyyy-MM-dd format)
    WeekDates []time.Time `json:"weekDates"`

}

// String returns a JSON representation of the model
func (o *Queryweeksummarylistjobrequest) String() string {
     o.ManagementUnitIds = []string{""} 
     o.WeekDates = []time.Time{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryweeksummarylistjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryweeksummarylistjobrequest

    if QueryweeksummarylistjobrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryweeksummarylistjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        WeekDates []time.Time `json:"weekDates"`
        *Alias
    }{

        
        ManagementUnitIds: []string{""},
        


        
        WeekDates: []time.Time{{}},
        

        Alias: (*Alias)(u),
    })
}

