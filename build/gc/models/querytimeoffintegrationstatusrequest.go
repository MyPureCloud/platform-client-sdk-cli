package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerytimeoffintegrationstatusrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerytimeoffintegrationstatusrequestDud struct { 
    

}

// Querytimeoffintegrationstatusrequest
type Querytimeoffintegrationstatusrequest struct { 
    // TimeOffRequestLookups - A list of time off request lookups
    TimeOffRequestLookups []Timeoffrequestlookup `json:"timeOffRequestLookups"`

}

// String returns a JSON representation of the model
func (o *Querytimeoffintegrationstatusrequest) String() string {
     o.TimeOffRequestLookups = []Timeoffrequestlookup{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querytimeoffintegrationstatusrequest) MarshalJSON() ([]byte, error) {
    type Alias Querytimeoffintegrationstatusrequest

    if QuerytimeoffintegrationstatusrequestMarshalled {
        return []byte("{}"), nil
    }
    QuerytimeoffintegrationstatusrequestMarshalled = true

    return json.Marshal(&struct {
        
        TimeOffRequestLookups []Timeoffrequestlookup `json:"timeOffRequestLookups"`
        *Alias
    }{

        
        TimeOffRequestLookups: []Timeoffrequestlookup{{}},
        

        Alias: (*Alias)(u),
    })
}

