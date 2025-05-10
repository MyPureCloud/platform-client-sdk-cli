package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableinputcolumnexpressionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableinputcolumnexpressionDud struct { 
    


    

}

// Decisiontableinputcolumnexpression
type Decisiontableinputcolumnexpression struct { 
    // Contractual - A value that is defined by a contract schema and used to form the left side of a logical condition.
    Contractual *Contractual `json:"contractual"`


    // Comparator - A comparator used to join the left and right sides of a logical condition.
    Comparator string `json:"comparator"`

}

// String returns a JSON representation of the model
func (o *Decisiontableinputcolumnexpression) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableinputcolumnexpression) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableinputcolumnexpression

    if DecisiontableinputcolumnexpressionMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableinputcolumnexpressionMarshalled = true

    return json.Marshal(&struct {
        
        Contractual *Contractual `json:"contractual"`
        
        Comparator string `json:"comparator"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

