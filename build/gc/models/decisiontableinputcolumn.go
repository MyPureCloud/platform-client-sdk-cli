package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableinputcolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableinputcolumnDud struct { 
    Id string `json:"id"`


    


    

}

// Decisiontableinputcolumn
type Decisiontableinputcolumn struct { 
    


    // DefaultsTo - The default row value for this column that will complete the condition expression where no value is provided by a row.
    DefaultsTo Decisiontablecolumndefaultrowvalue `json:"defaultsTo"`


    // Expression - The input column condition expression, comprising the left side and comparator of a logical condition in the form of left|comparator|right, where each row of the decision table will provide the right side to form a complete condition
    Expression Decisiontableinputcolumnexpression `json:"expression"`

}

// String returns a JSON representation of the model
func (o *Decisiontableinputcolumn) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableinputcolumn) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableinputcolumn

    if DecisiontableinputcolumnMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableinputcolumnMarshalled = true

    return json.Marshal(&struct {
        
        DefaultsTo Decisiontablecolumndefaultrowvalue `json:"defaultsTo"`
        
        Expression Decisiontableinputcolumnexpression `json:"expression"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

