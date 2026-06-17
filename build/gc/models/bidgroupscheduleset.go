package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BidgroupschedulesetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BidgroupschedulesetDud struct { 
    

}

// Bidgroupscheduleset
type Bidgroupscheduleset struct { 
    // ShiftSets - The shift sets that will be used for schedule generation
    ShiftSets []Shiftset `json:"shiftSets"`

}

// String returns a JSON representation of the model
func (o *Bidgroupscheduleset) String() string {
     o.ShiftSets = []Shiftset{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bidgroupscheduleset) MarshalJSON() ([]byte, error) {
    type Alias Bidgroupscheduleset

    if BidgroupschedulesetMarshalled {
        return []byte("{}"), nil
    }
    BidgroupschedulesetMarshalled = true

    return json.Marshal(&struct {
        
        ShiftSets []Shiftset `json:"shiftSets"`
        *Alias
    }{

        
        ShiftSets: []Shiftset{{}},
        

        Alias: (*Alias)(u),
    })
}

