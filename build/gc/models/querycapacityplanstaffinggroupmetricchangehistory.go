package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerycapacityplanstaffinggroupmetricchangehistoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerycapacityplanstaffinggroupmetricchangehistoryDud struct { 
    


    

}

// Querycapacityplanstaffinggroupmetricchangehistory
type Querycapacityplanstaffinggroupmetricchangehistory struct { 
    // StaffingGroupIds - The IDs of the staffing groups for which to fetch the metric change history
    StaffingGroupIds []string `json:"staffingGroupIds"`


    // QueryEndDate - Fetch the history of metric changes on or before this date, in IS0-8601 format. If not specified, sets to current date
    QueryEndDate time.Time `json:"queryEndDate"`

}

// String returns a JSON representation of the model
func (o *Querycapacityplanstaffinggroupmetricchangehistory) String() string {
     o.StaffingGroupIds = []string{""} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querycapacityplanstaffinggroupmetricchangehistory) MarshalJSON() ([]byte, error) {
    type Alias Querycapacityplanstaffinggroupmetricchangehistory

    if QuerycapacityplanstaffinggroupmetricchangehistoryMarshalled {
        return []byte("{}"), nil
    }
    QuerycapacityplanstaffinggroupmetricchangehistoryMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroupIds []string `json:"staffingGroupIds"`
        
        QueryEndDate time.Time `json:"queryEndDate"`
        *Alias
    }{

        
        StaffingGroupIds: []string{""},
        


        

        Alias: (*Alias)(u),
    })
}

