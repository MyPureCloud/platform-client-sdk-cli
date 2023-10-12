package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangeworkitemstatuscategoryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangeworkitemstatuscategoryDud struct { 
    


    

}

// Workitemsattributechangeworkitemstatuscategory
type Workitemsattributechangeworkitemstatuscategory struct { 
    // NewValue - New property value
    NewValue string `json:"newValue"`


    // OldValue - Old property value
    OldValue string `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangeworkitemstatuscategory) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangeworkitemstatuscategory) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangeworkitemstatuscategory

    if WorkitemsattributechangeworkitemstatuscategoryMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangeworkitemstatuscategoryMarshalled = true

    return json.Marshal(&struct {
        
        NewValue string `json:"newValue"`
        
        OldValue string `json:"oldValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

