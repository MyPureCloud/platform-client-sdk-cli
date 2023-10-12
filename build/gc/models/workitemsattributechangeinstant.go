package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemsattributechangeinstantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemsattributechangeinstantDud struct { 
    


    

}

// Workitemsattributechangeinstant
type Workitemsattributechangeinstant struct { 
    // NewValue - New property value. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    NewValue time.Time `json:"newValue"`


    // OldValue - Old property value. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    OldValue time.Time `json:"oldValue"`

}

// String returns a JSON representation of the model
func (o *Workitemsattributechangeinstant) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemsattributechangeinstant) MarshalJSON() ([]byte, error) {
    type Alias Workitemsattributechangeinstant

    if WorkitemsattributechangeinstantMarshalled {
        return []byte("{}"), nil
    }
    WorkitemsattributechangeinstantMarshalled = true

    return json.Marshal(&struct {
        
        NewValue time.Time `json:"newValue"`
        
        OldValue time.Time `json:"oldValue"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

