package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WorkitemonattributechangeruleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WorkitemonattributechangeruleDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Workitemonattributechangerule
type Workitemonattributechangerule struct { 
    


    // Name
    Name string `json:"name"`


    // VarType - The type of the rule.
    VarType string `json:"type"`


    // Action - The rules action. If the condition criteria is met this action will be executed.
    Action Workitemruleaction `json:"action"`


    // Worktype - The Worktype containing the rule.
    Worktype Worktypereference `json:"worktype"`


    // Condition - The rules condition. If the condition criteria is met the rules action will be executed.
    Condition Workitemonattributechangecondition `json:"condition"`


    

}

// String returns a JSON representation of the model
func (o *Workitemonattributechangerule) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Workitemonattributechangerule) MarshalJSON() ([]byte, error) {
    type Alias Workitemonattributechangerule

    if WorkitemonattributechangeruleMarshalled {
        return []byte("{}"), nil
    }
    WorkitemonattributechangeruleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Action Workitemruleaction `json:"action"`
        
        Worktype Worktypereference `json:"worktype"`
        
        Condition Workitemonattributechangecondition `json:"condition"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

