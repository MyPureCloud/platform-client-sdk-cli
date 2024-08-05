package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkplanbidgroupsummarylistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkplanbidgroupsummarylistDud struct { 
    

}

// Workplanbidgroupsummarylist
type Workplanbidgroupsummarylist struct { 
    // WorkPlanBidGroupSummaryList - List of work plan bid group summary
    WorkPlanBidGroupSummaryList []Workplanbidgroupsummary `json:"workPlanBidGroupSummaryList"`

}

// String returns a JSON representation of the model
func (o *Workplanbidgroupsummarylist) String() string {
     o.WorkPlanBidGroupSummaryList = []Workplanbidgroupsummary{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workplanbidgroupsummarylist) MarshalJSON() ([]byte, error) {
    type Alias Workplanbidgroupsummarylist

    if WorkplanbidgroupsummarylistMarshalled {
        return []byte("{}"), nil
    }
    WorkplanbidgroupsummarylistMarshalled = true

    return json.Marshal(&struct {
        
        WorkPlanBidGroupSummaryList []Workplanbidgroupsummary `json:"workPlanBidGroupSummaryList"`
        *Alias
    }{

        
        WorkPlanBidGroupSummaryList: []Workplanbidgroupsummary{{}},
        

        Alias: (*Alias)(u),
    })
}

