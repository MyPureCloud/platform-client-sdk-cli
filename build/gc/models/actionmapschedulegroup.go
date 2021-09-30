package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionmapschedulegroupMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionmapschedulegroupDud struct { 
    

}

// Actionmapschedulegroup
type Actionmapschedulegroup struct { 
    // Id - The ID of the action maps's associated schedule group.
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Actionmapschedulegroup) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionmapschedulegroup) MarshalJSON() ([]byte, error) {
    type Alias Actionmapschedulegroup

    if ActionmapschedulegroupMarshalled {
        return []byte("{}"), nil
    }
    ActionmapschedulegroupMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

