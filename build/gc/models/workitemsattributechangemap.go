package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangemapMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangemapDud struct { 
    


    

}

// Workitemsattributechangemap
type Workitemsattributechangemap struct { 
    // NewValue - New property value
    NewValue map[string]interface{} `json:"newValue"`


    // OldValue - Old property value
    OldValue map[string]interface{} `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangemap) String() string {
     o.NewValue = map[string]interface{}{"": Interface{}} 
     o.OldValue = map[string]interface{}{"": Interface{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangemap) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangemap

    if WorkitemsattributechangemapMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangemapMarshalled = true

    return json.Marshal(&struct {
        
        NewValue map[string]interface{} `json:"newValue"`
        
        OldValue map[string]interface{} `json:"oldValue"`
        *Alias
    }{

        
        NewValue: map[string]interface{}{"": Interface{}},
        


        
        OldValue: map[string]interface{}{"": Interface{}},
        

        Alias: (*Alias)(u),
    })
}

