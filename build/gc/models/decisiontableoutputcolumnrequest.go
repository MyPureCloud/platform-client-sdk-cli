package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableoutputcolumnrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableoutputcolumnrequestDud struct { 
    


    

}

// Decisiontableoutputcolumnrequest
type Decisiontableoutputcolumnrequest struct { 
    // DefaultsTo - The default row value for this column that will be used for an output value where no value  is provided by a row.
    DefaultsTo Decisiontablecolumndefaultrowvalue `json:"defaultsTo"`


    // Value - The output data of this column that will be provided by each row.
    Value Outputvalue `json:"value"`

}

// String returns a JSON representation of the model
func (o *Decisiontableoutputcolumnrequest) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableoutputcolumnrequest) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableoutputcolumnrequest

    if DecisiontableoutputcolumnrequestMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableoutputcolumnrequestMarshalled = true

    return json.Marshal(&struct {
        
        DefaultsTo Decisiontablecolumndefaultrowvalue `json:"defaultsTo"`
        
        Value Outputvalue `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

