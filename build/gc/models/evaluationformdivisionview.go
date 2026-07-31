package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationformdivisionviewMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationformdivisionviewDud struct { 
    Id string `json:"id"`


    


    


    ContextId string `json:"contextId"`


    ModifiedDate time.Time `json:"modifiedDate"`


    


    SelfUri string `json:"selfUri"`

}

// Evaluationformdivisionview
type Evaluationformdivisionview struct { 
    


    // Name
    Name string `json:"name"`


    // Division - The division to which this entity belongs.
    Division Writablestarrabledivision `json:"division"`


    


    


    // Dialect - The language dialect for this evaluation form. Supported dialects: ar, cs, da, de, en-US, es, fi, fr, fr-CA, he, hi, it, ja, ko, nl, no, pl, pt-BR, pt-PT, ru, sv, th, tr, uk, zh-CN, zh-TW
    Dialect string `json:"dialect"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationformdivisionview) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationformdivisionview) MarshalJSON() ([]byte, error) {
    type Alias Evaluationformdivisionview

    if EvaluationformdivisionviewMarshalled {
        return []byte("{}"), nil
    }
    EvaluationformdivisionviewMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Division Writablestarrabledivision `json:"division"`
        
        Dialect string `json:"dialect"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

