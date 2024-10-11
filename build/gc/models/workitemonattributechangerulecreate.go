package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemonattributechangerulecreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemonattributechangerulecreateDud struct { 
    


    

}

// Workitemonattributechangerulecreate
type Workitemonattributechangerulecreate struct { 
    // Name - The name of the rule.
    Name string `json:"name"`


    // Condition - The rules condition. If the condition criteria is met the rules action will be executed.
    Condition Workitemonattributechangecondition `json:"condition"`

}

// String returns a JSON representation of the model
func (o *Workitemonattributechangerulecreate) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemonattributechangerulecreate) MarshalJSON() ([]byte, error) {
    type Alias Workitemonattributechangerulecreate

    if WorkitemonattributechangerulecreateMarshalled {
        return []byte("{}"), nil
    }
    WorkitemonattributechangerulecreateMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Condition Workitemonattributechangecondition `json:"condition"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

