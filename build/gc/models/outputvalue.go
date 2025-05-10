package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OutputvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OutputvalueDud struct { 
    


    

}

// Outputvalue
type Outputvalue struct { 
    // SchemaPropertyKey - The contract schema property key that describes the output value of this column.
    SchemaPropertyKey string `json:"schemaPropertyKey"`


    // Properties - The nested output properties of this column that will be provided by each row, if any.
    Properties []Outputvalue `json:"properties"`

}

// String returns a JSON representation of the model
func (o *Outputvalue) String() string {
    
     o.Properties = []Outputvalue{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Outputvalue) MarshalJSON() ([]byte, error) {
    type Alias Outputvalue

    if OutputvalueMarshalled {
        return []byte("{}"), nil
    }
    OutputvalueMarshalled = true

    return json.Marshal(&struct {
        
        SchemaPropertyKey string `json:"schemaPropertyKey"`
        
        Properties []Outputvalue `json:"properties"`
        *Alias
    }{

        


        
        Properties: []Outputvalue{{}},
        

        Alias: (*Alias)(u),
    })
}

