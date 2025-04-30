package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangeworkitemrulesettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangeworkitemrulesettingsDud struct { 
    


    

}

// Workitemsattributechangeworkitemrulesettings
type Workitemsattributechangeworkitemrulesettings struct { 
    // NewValue - New property value
    NewValue Workitemrulesettings `json:"newValue"`


    // OldValue - Old property value
    OldValue Workitemrulesettings `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangeworkitemrulesettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangeworkitemrulesettings) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangeworkitemrulesettings

    if WorkitemsattributechangeworkitemrulesettingsMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangeworkitemrulesettingsMarshalled = true

    return json.Marshal(&struct {
        
        NewValue Workitemrulesettings `json:"newValue"`
        
        OldValue Workitemrulesettings `json:"oldValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

