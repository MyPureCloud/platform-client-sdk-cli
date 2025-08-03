package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapacityplandeleterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapacityplandeleterequestDud struct { 
    

}

// Capacityplandeleterequest
type Capacityplandeleterequest struct { 
    // CapacityPlanIds - The IDs of the capacity plans to be deleted
    CapacityPlanIds []string `json:"capacityPlanIds"`

}

// String returns a JSON representation of the model
func (o *Capacityplandeleterequest) String() string {
     o.CapacityPlanIds = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capacityplandeleterequest) MarshalJSON() ([]byte, error) {
    type Alias Capacityplandeleterequest

    if CapacityplandeleterequestMarshalled {
        return []byte("{}"), nil
    }
    CapacityplandeleterequestMarshalled = true

    return json.Marshal(&struct {
        
        CapacityPlanIds []string `json:"capacityPlanIds"`
        *Alias
    }{

        
        CapacityPlanIds: []string{""},
        

        Alias: (*Alias)(u),
    })
}

