package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SchedulebidgroupsummarylistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SchedulebidgroupsummarylistDud struct { 
    

}

// Schedulebidgroupsummarylist
type Schedulebidgroupsummarylist struct { 
    // ScheduleBidGroupSummaries - Schedule bid group summary
    ScheduleBidGroupSummaries []Schedulebidgroupsummary `json:"scheduleBidGroupSummaries"`

}

// String returns a JSON representation of the model
func (o *Schedulebidgroupsummarylist) String() string {
     o.ScheduleBidGroupSummaries = []Schedulebidgroupsummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Schedulebidgroupsummarylist) MarshalJSON() ([]byte, error) {
    type Alias Schedulebidgroupsummarylist

    if SchedulebidgroupsummarylistMarshalled {
        return []byte("{}"), nil
    }
    SchedulebidgroupsummarylistMarshalled = true

    return json.Marshal(&struct {
        
        ScheduleBidGroupSummaries []Schedulebidgroupsummary `json:"scheduleBidGroupSummaries"`
        *Alias
    }{

        
        ScheduleBidGroupSummaries: []Schedulebidgroupsummary{{}},
        

        Alias: (*Alias)(u),
    })
}

