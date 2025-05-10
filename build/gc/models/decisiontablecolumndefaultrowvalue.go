package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablecolumndefaultrowvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablecolumndefaultrowvalueDud struct { 
    


    

}

// Decisiontablecolumndefaultrowvalue - Must provide a valid value for exactly one of the fields in this class.
type Decisiontablecolumndefaultrowvalue struct { 
    // Value - A default string value for this column, will be cast to appropriate type according to the relevant contract schema property.
    Value string `json:"value"`


    // Special - A default special value enum for this column.
    Special string `json:"special"`

}

// String returns a JSON representation of the model
func (o *Decisiontablecolumndefaultrowvalue) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablecolumndefaultrowvalue) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablecolumndefaultrowvalue

    if DecisiontablecolumndefaultrowvalueMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablecolumndefaultrowvalueMarshalled = true

    return json.Marshal(&struct {
        
        Value string `json:"value"`
        
        Special string `json:"special"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

