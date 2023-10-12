package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangebooleanMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangebooleanDud struct { 
    


    

}

// Workitemsattributechangeboolean
type Workitemsattributechangeboolean struct { 
    // NewValue - New property value
    NewValue bool `json:"newValue"`


    // OldValue - Old property value
    OldValue bool `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangeboolean) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangeboolean) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangeboolean

    if WorkitemsattributechangebooleanMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangebooleanMarshalled = true

    return json.Marshal(&struct {
        
        NewValue bool `json:"newValue"`
        
        OldValue bool `json:"oldValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

