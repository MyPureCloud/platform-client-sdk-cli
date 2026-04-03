package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QueryagentshifttradelistjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QueryagentshifttradelistjobrequestDud struct { 
    

}

// Queryagentshifttradelistjobrequest
type Queryagentshifttradelistjobrequest struct { 
    // WeekDates - The start week dates in which to query shift trades in the business unit time zone (yyyy-MM-dd format)
    WeekDates []time.Time `json:"weekDates"`

}

// String returns a JSON representation of the model
func (o *Queryagentshifttradelistjobrequest) String() string {
     o.WeekDates = []time.Time{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Queryagentshifttradelistjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Queryagentshifttradelistjobrequest

    if QueryagentshifttradelistjobrequestMarshalled {
        return []byte("{}"), nil
    }
    QueryagentshifttradelistjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        WeekDates []time.Time `json:"weekDates"`
        *Alias
    }{

        
        WeekDates: []time.Time{{}},
        

        Alias: (*Alias)(u),
    })
}

