package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangelistMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangelistDud struct { 
    


    

}

// Workitemsattributechangelist
type Workitemsattributechangelist struct { 
    // NewValue - New property value
    NewValue []interface{} `json:"newValue"`


    // OldValue - Old property value
    OldValue []interface{} `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangelist) String() string {
     o.NewValue = []interface{}{} 
     o.OldValue = []interface{}{} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangelist) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangelist

    if WorkitemsattributechangelistMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangelistMarshalled = true

    return json.Marshal(&struct {
        
        NewValue []interface{} `json:"newValue"`
        
        OldValue []interface{} `json:"oldValue"`
        *Alias
    }{

        
        NewValue: []interface{}{},
        


        
        OldValue: []interface{}{},
        

        Alias: (*Alias)(u),
    })
}

