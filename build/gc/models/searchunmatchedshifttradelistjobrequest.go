package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SearchunmatchedshifttradelistjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SearchunmatchedshifttradelistjobrequestDud struct { 
    


    


    


    

}

// Searchunmatchedshifttradelistjobrequest
type Searchunmatchedshifttradelistjobrequest struct { 
    // ManagementUnitIds - The IDs of management units from which to query shift trades
    ManagementUnitIds []string `json:"managementUnitIds"`


    // WeekDates - The start week dates in which to query shift trades in the business unit time zone (yyyy-MM-dd format)
    WeekDates []time.Time `json:"weekDates"`


    // ReceivingSchedule - Associated schedule information for the receiving user
    ReceivingSchedule Receivingschedulelookup `json:"receivingSchedule"`


    // ReceivingShiftIds - The IDs of shifts that the receiving user would potentially be willing to trade. If empty, only returns one-sided trades
    ReceivingShiftIds []string `json:"receivingShiftIds"`

}

// String returns a JSON representation of the model
func (o *Searchunmatchedshifttradelistjobrequest) String() string {
     o.ManagementUnitIds = []string{""} 
     o.WeekDates = []time.Time{{}} 
    
     o.ReceivingShiftIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Searchunmatchedshifttradelistjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Searchunmatchedshifttradelistjobrequest

    if SearchunmatchedshifttradelistjobrequestMarshalled {
        return []byte("{}"), nil
    }
    SearchunmatchedshifttradelistjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        ManagementUnitIds []string `json:"managementUnitIds"`
        
        WeekDates []time.Time `json:"weekDates"`
        
        ReceivingSchedule Receivingschedulelookup `json:"receivingSchedule"`
        
        ReceivingShiftIds []string `json:"receivingShiftIds"`
        *Alias
    }{

        
        ManagementUnitIds: []string{""},
        


        
        WeekDates: []time.Time{{}},
        


        


        
        ReceivingShiftIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

