package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeoffbalancerequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeoffbalancerequestDud struct { 
    


    

}

// Timeoffbalancerequest
type Timeoffbalancerequest struct { 
    // ActivityCodeIds - The set of activity code IDs for which to query available time off balances
    ActivityCodeIds []string `json:"activityCodeIds"`


    // DateRanges - The list of date ranges for which to query time off balance
    DateRanges []Localdaterange `json:"dateRanges"`

}

// String returns a JSON representation of the model
func (o *Timeoffbalancerequest) String() string {
     o.ActivityCodeIds = []string{""} 
     o.DateRanges = []Localdaterange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeoffbalancerequest) MarshalJSON() ([]byte, error) {
    type Alias Timeoffbalancerequest

    if TimeoffbalancerequestMarshalled {
        return []byte("{}"), nil
    }
    TimeoffbalancerequestMarshalled = true

    return json.Marshal(&struct {
        
        ActivityCodeIds []string `json:"activityCodeIds"`
        
        DateRanges []Localdaterange `json:"dateRanges"`
        *Alias
    }{

        
        ActivityCodeIds: []string{""},
        


        
        DateRanges: []Localdaterange{{}},
        

        Alias: (*Alias)(u),
    })
}

