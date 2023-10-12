package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangestringMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangestringDud struct { 
    


    

}

// Workitemsattributechangestring
type Workitemsattributechangestring struct { 
    // NewValue - New property value
    NewValue string `json:"newValue"`


    // OldValue - Old property value
    OldValue string `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangestring) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangestring) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangestring

    if WorkitemsattributechangestringMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangestringMarshalled = true

    return json.Marshal(&struct {
        
        NewValue string `json:"newValue"`
        
        OldValue string `json:"oldValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

