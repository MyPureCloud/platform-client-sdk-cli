package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemdatebasedruleupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemdatebasedruleupdateDud struct { 
    


    

}

// Workitemdatebasedruleupdate
type Workitemdatebasedruleupdate struct { 
    // Name - The name of the rule.
    Name string `json:"name"`


    // Condition - The rules condition. If the condition criteria is met the rules action will be executed.
    Condition Workitemdatebasedconditionupdate `json:"condition"`

}

// String returns a JSON representation of the model
func (o *Workitemdatebasedruleupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemdatebasedruleupdate) MarshalJSON() ([]byte, error) {
    type Alias Workitemdatebasedruleupdate

    if WorkitemdatebasedruleupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemdatebasedruleupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Condition Workitemdatebasedconditionupdate `json:"condition"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

