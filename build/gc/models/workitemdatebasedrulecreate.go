package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemdatebasedrulecreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemdatebasedrulecreateDud struct { 
    


    

}

// Workitemdatebasedrulecreate
type Workitemdatebasedrulecreate struct { 
    // Name - The name of the rule.
    Name string `json:"name"`


    // Condition - The rules condition. If the condition criteria is met the rules action will be executed.
    Condition Workitemdatebasedcondition `json:"condition"`

}

// String returns a JSON representation of the model
func (o *Workitemdatebasedrulecreate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemdatebasedrulecreate) MarshalJSON() ([]byte, error) {
    type Alias Workitemdatebasedrulecreate

    if WorkitemdatebasedrulecreateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemdatebasedrulecreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Condition Workitemdatebasedcondition `json:"condition"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

