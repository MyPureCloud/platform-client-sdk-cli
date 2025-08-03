package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    StaffinggroupallocationsresponsetemplateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type StaffinggroupallocationsresponsetemplateDud struct { 
    

}

// Staffinggroupallocationsresponsetemplate
type Staffinggroupallocationsresponsetemplate struct { 
    // StaffingGroupAllocations - List of staffing group allocations
    StaffingGroupAllocations []Staffinggroupallocation `json:"staffingGroupAllocations"`

}

// String returns a JSON representation of the model
func (o *Staffinggroupallocationsresponsetemplate) String() string {
     o.StaffingGroupAllocations = []Staffinggroupallocation{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Staffinggroupallocationsresponsetemplate) MarshalJSON() ([]byte, error) {
    type Alias Staffinggroupallocationsresponsetemplate

    if StaffinggroupallocationsresponsetemplateMarshalled {
        return []byte("{}"), nil
    }
    StaffinggroupallocationsresponsetemplateMarshalled = true

    return json.Marshal(&struct {
        
        StaffingGroupAllocations []Staffinggroupallocation `json:"staffingGroupAllocations"`
        *Alias
    }{

        
        StaffingGroupAllocations: []Staffinggroupallocation{{}},
        

        Alias: (*Alias)(u),
    })
}

