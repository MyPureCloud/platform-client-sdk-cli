package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemschangeworkitemdeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemschangeworkitemdeltaDud struct { 
    


    


    

}

// Workitemschangeworkitemdelta
type Workitemschangeworkitemdelta struct { 
    // Version - Version
    Version int `json:"version"`


    // ModifiedBy - modifiedBy
    ModifiedBy Userreference `json:"modifiedBy"`


    // Delta - The changes that originated this version
    Delta Workitemdelta `json:"delta"`

}

// String returns a JSON representation of the model
func (o *Workitemschangeworkitemdelta) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemschangeworkitemdelta) MarshalJSON() ([]byte, error) {
    type Alias Workitemschangeworkitemdelta

    if WorkitemschangeworkitemdeltaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemschangeworkitemdeltaMarshalled = true

    return json.Marshal(&struct {
        
        Version int `json:"version"`
        
        ModifiedBy Userreference `json:"modifiedBy"`
        
        Delta Workitemdelta `json:"delta"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

