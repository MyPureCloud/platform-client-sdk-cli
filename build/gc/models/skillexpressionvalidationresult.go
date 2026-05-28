package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillexpressionvalidationresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillexpressionvalidationresultDud struct { 
    


    


    


    


    

}

// Skillexpressionvalidationresult - Result of skill expression validation
type Skillexpressionvalidationresult struct { 
    // Valid - Whether the expression is valid
    Valid bool `json:"valid"`


    // Expression - Normalized SpEL expression (null if validation failed)
    Expression string `json:"expression"`


    // Skills - List of skill references extracted from the expression (empty if no skills found and/or invalid expression)
    Skills []Skillreference `json:"skills"`


    // Errors - List of validation errors (empty if valid)
    Errors []Skillexpressionvalidationerror `json:"errors"`


    // Hint - Optional hint message (e.g., if expression is non-optimal or system is near capacity)
    Hint string `json:"hint"`

}

// String returns a JSON representation of the model
func (o *Skillexpressionvalidationresult) String() string {
    
    
     o.Skills = []Skillreference{{}} 
     o.Errors = []Skillexpressionvalidationerror{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillexpressionvalidationresult) MarshalJSON() ([]byte, error) {
    type Alias Skillexpressionvalidationresult

    if SkillexpressionvalidationresultMarshalled {
        return []byte("{}"), nil
    }
    SkillexpressionvalidationresultMarshalled = true

    return json.Marshal(&struct {
        
        Valid bool `json:"valid"`
        
        Expression string `json:"expression"`
        
        Skills []Skillreference `json:"skills"`
        
        Errors []Skillexpressionvalidationerror `json:"errors"`
        
        Hint string `json:"hint"`
        *Alias
    }{

        


        


        
        Skills: []Skillreference{{}},
        


        
        Errors: []Skillexpressionvalidationerror{{}},
        


        

        Alias: (*Alias)(u),
    })
}

