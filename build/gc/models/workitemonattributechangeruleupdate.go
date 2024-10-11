package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemonattributechangeruleupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemonattributechangeruleupdateDud struct { 
    


    

}

// Workitemonattributechangeruleupdate
type Workitemonattributechangeruleupdate struct { 
    // Name - The name of the rule.
    Name string `json:"name"`


    // Condition - The rules condition. If the condition criteria is met the rules action will be executed.
    Condition Workitemonattributechangeconditionupdate `json:"condition"`

}

// String returns a JSON representation of the model
func (o *Workitemonattributechangeruleupdate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemonattributechangeruleupdate) MarshalJSON() ([]byte, error) {
    type Alias Workitemonattributechangeruleupdate

    if WorkitemonattributechangeruleupdateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemonattributechangeruleupdateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Condition Workitemonattributechangeconditionupdate `json:"condition"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

