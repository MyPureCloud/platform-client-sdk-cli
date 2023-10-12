package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemschangeworktypedeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemschangeworktypedeltaDud struct { 
    


    


    

}

// Workitemschangeworktypedelta
type Workitemschangeworktypedelta struct { 
    // Version - Version
    Version int `json:"version"`


    // ModifiedBy - modifiedBy
    ModifiedBy Userreference `json:"modifiedBy"`


    // Delta - The changes that originated this version
    Delta Worktypedelta `json:"delta"`

}

// String returns a JSON representation of the model
func (o *Workitemschangeworktypedelta) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemschangeworktypedelta) MarshalJSON() ([]byte, error) {
    type Alias Workitemschangeworktypedelta

    if WorkitemschangeworktypedeltaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemschangeworktypedeltaMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        Delta Worktypedelta `json:"delta"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

