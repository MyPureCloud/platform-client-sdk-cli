package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentscheduleupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentscheduleupdateDud struct { 
    


    

}

// Buagentscheduleupdate
type Buagentscheduleupdate struct { 
    // VarType - The type of update
    VarType string `json:"type"`


    // ShiftStartDates - The start date for the affected shifts
    ShiftStartDates []time.Time `json:"shiftStartDates"`

}

// String returns a JSON representation of the model
func (o *Buagentscheduleupdate) String() string {
    
     o.ShiftStartDates = []time.Time{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentscheduleupdate) MarshalJSON() ([]byte, error) {
    type Alias Buagentscheduleupdate

    if BuagentscheduleupdateMarshalled {
        return []byte("{}"), nil
    }
    BuagentscheduleupdateMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        ShiftStartDates []time.Time `json:"shiftStartDates"`
        *Alias
    }{

        


        
        ShiftStartDates: []time.Time{{}},
        

        Alias: (*Alias)(u),
    })
}

