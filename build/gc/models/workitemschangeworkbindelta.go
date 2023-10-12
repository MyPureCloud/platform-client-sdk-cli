package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemschangeworkbindeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemschangeworkbindeltaDud struct { 
    


    


    

}

// Workitemschangeworkbindelta
type Workitemschangeworkbindelta struct { 
    // Version - Version
    Version int `json:"version"`


    // ModifiedBy - modifiedBy
    ModifiedBy Userreference `json:"modifiedBy"`


    // Delta - The changes that originated this version
    Delta Workbindelta `json:"delta"`

}

// String returns a JSON representation of the model
func (o *Workitemschangeworkbindelta) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemschangeworkbindelta) MarshalJSON() ([]byte, error) {
    type Alias Workitemschangeworkbindelta

    if WorkitemschangeworkbindeltaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemschangeworkbindeltaMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        Delta Workbindelta `json:"delta"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

