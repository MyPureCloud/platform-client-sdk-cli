package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InfixoperatorMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InfixoperatorDud struct { 
    


    

}

// Infixoperator
type Infixoperator struct { 
    // OperatorType - The logical operation that is applied on the operand against the following operand
    OperatorType string `json:"operatorType"`


    // OperatorPosition - Dictates when the following operand should occur relative to current operand
    OperatorPosition Operatorposition `json:"operatorPosition"`

}

// String returns a JSON representation of the model
func (o *Infixoperator) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Infixoperator) MarshalJSON() ([]byte, error) {
    type Alias Infixoperator

    if InfixoperatorMarshalled {
        return []byte("{}"), nil
    }
    InfixoperatorMarshalled = true

    return json.Marshal(&struct {
        
        OperatorType string `json:"operatorType"`
        
        OperatorPosition Operatorposition `json:"operatorPosition"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

