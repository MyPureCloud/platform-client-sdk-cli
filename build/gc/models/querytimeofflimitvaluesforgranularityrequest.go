package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerytimeofflimitvaluesforgranularityrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerytimeofflimitvaluesforgranularityrequestDud struct { 
    

}

// Querytimeofflimitvaluesforgranularityrequest
type Querytimeofflimitvaluesforgranularityrequest struct { 
    // DateRanges - The date range to return time off limit, allocated and waitlisted minutes. Maximum allowed number of days in range in 366
    DateRanges []Localdaterange `json:"dateRanges"`

}

// String returns a JSON representation of the model
func (o *Querytimeofflimitvaluesforgranularityrequest) String() string {
     o.DateRanges = []Localdaterange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querytimeofflimitvaluesforgranularityrequest) MarshalJSON() ([]byte, error) {
    type Alias Querytimeofflimitvaluesforgranularityrequest

    if QuerytimeofflimitvaluesforgranularityrequestMarshalled {
        return []byte("{}"), nil
    }
    QuerytimeofflimitvaluesforgranularityrequestMarshalled = true

    return json.Marshal(&struct {
        
        DateRanges []Localdaterange `json:"dateRanges"`
        *Alias
    }{

        
        DateRanges: []Localdaterange{{}},
        

        Alias: (*Alias)(u),
    })
}

