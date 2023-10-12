package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangeintegerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangeintegerDud struct { 
    


    

}

// Workitemsattributechangeinteger
type Workitemsattributechangeinteger struct { 
    // NewValue - New property value
    NewValue int `json:"newValue"`


    // OldValue - Old property value
    OldValue int `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangeinteger) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangeinteger) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangeinteger

    if WorkitemsattributechangeintegerMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangeintegerMarshalled = true

    return json.Marshal(&struct {
        
        NewValue int `json:"newValue"`
        
        OldValue int `json:"oldValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

