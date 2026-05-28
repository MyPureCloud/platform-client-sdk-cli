package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SkillexpressiondataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SkillexpressiondataDud struct { 
    

}

// Skillexpressiondata - Request data for skill expression validation
type Skillexpressiondata struct { 
    // Expression - The skill expression in raw format to validate
    Expression string `json:"expression"`

}

// String returns a JSON representation of the model
func (o *Skillexpressiondata) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Skillexpressiondata) MarshalJSON() ([]byte, error) {
    type Alias Skillexpressiondata

    if SkillexpressiondataMarshalled {
        return []byte("{}"), nil
    }
    SkillexpressiondataMarshalled = true

    return json.Marshal(&struct {
        
        Expression string `json:"expression"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

