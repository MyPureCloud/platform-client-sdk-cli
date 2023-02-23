package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EvaluationsourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EvaluationsourceDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Evaluationsource
type Evaluationsource struct { 
    


    // Name
    Name string `json:"name"`


    // VarType - Type of the evaluation source.
    VarType string `json:"type"`


    

}

// String returns a JSON representation of the model
func (o *Evaluationsource) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Evaluationsource) MarshalJSON() ([]byte, error) {
    type Alias Evaluationsource

    if EvaluationsourceMarshalled {
        return []byte("{}"), nil
    }
    EvaluationsourceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

