package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableoutputcolumnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableoutputcolumnDud struct { 
    Id string `json:"id"`


    


    

}

// Decisiontableoutputcolumn
type Decisiontableoutputcolumn struct { 
    


    // DefaultsTo - The default row value for this column that will be used for an output value where no value  is provided by a row.
    DefaultsTo Decisiontablecolumndefaultrowvalue `json:"defaultsTo"`


    // Value - The output data of this column that will be provided by each row.
    Value Outputvalue `json:"value"`

}

// String returns a JSON representation of the model
func (o *Decisiontableoutputcolumn) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableoutputcolumn) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableoutputcolumn

    if DecisiontableoutputcolumnMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableoutputcolumnMarshalled = true

    return json.Marshal(&struct {
        
        DefaultsTo Decisiontablecolumndefaultrowvalue `json:"defaultsTo"`
        
        Value Outputvalue `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

