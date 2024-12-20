package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangewrapupdeltaMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangewrapupdeltaDud struct { 
    


    

}

// Workitemsattributechangewrapupdelta
type Workitemsattributechangewrapupdelta struct { 
    // NewValue - New property value
    NewValue Wrapupdelta `json:"newValue"`


    // OldValue - Old property value
    OldValue Wrapupdelta `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangewrapupdelta) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangewrapupdelta) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangewrapupdelta

    if WorkitemsattributechangewrapupdeltaMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangewrapupdeltaMarshalled = true

    return json.Marshal(&struct {
        
        NewValue Wrapupdelta `json:"newValue"`
        
        OldValue Wrapupdelta `json:"oldValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

