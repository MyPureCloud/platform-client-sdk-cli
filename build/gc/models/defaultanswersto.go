package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DefaultanswerstoMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DefaultanswerstoDud struct { 
    


    


    


    

}

// Defaultanswersto
type Defaultanswersto struct { 
    // HighestScore - True, when answer should default to highest score
    HighestScore bool `json:"highestScore"`


    // NotApplicable - True, when answer should default to N/A
    NotApplicable bool `json:"notApplicable"`


    // LowestScore - True, when answer should default to lowest score
    LowestScore bool `json:"lowestScore"`


    // UserDefined - True, when answer should default to user defined answer
    UserDefined bool `json:"userDefined"`

}

// String returns a JSON representation of the model
func (o *Defaultanswersto) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Defaultanswersto) MarshalJSON() ([]byte, error) {
    type Alias Defaultanswersto

    if DefaultanswerstoMarshalled {
        return []byte("{}"), nil
    }
    DefaultanswerstoMarshalled = true

    return json.Marshal(&struct {
        
        HighestScore bool `json:"highestScore"`
        
        NotApplicable bool `json:"notApplicable"`
        
        LowestScore bool `json:"lowestScore"`
        
        UserDefined bool `json:"userDefined"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

