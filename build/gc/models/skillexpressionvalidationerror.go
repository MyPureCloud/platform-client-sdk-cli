package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillexpressionvalidationerrorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillexpressionvalidationerrorDud struct { 
    


    


    

}

// Skillexpressionvalidationerror - A validation error found in the expression
type Skillexpressionvalidationerror struct { 
    // Code - Error code
    Code string `json:"code"`


    // Message - Human-readable error message
    Message string `json:"message"`


    // Position - Position in the expression where the error occurred (null if not applicable)
    Position int `json:"position"`

}

// String returns a JSON representation of the model
func (o *Skillexpressionvalidationerror) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillexpressionvalidationerror) MarshalJSON() ([]byte, error) {
    type Alias Skillexpressionvalidationerror

    if SkillexpressionvalidationerrorMarshalled {
        return []byte("{}"), nil
    }
    SkillexpressionvalidationerrorMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        
        Position int `json:"position"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

