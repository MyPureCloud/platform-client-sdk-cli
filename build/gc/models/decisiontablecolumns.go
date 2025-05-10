package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablecolumnsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablecolumnsDud struct { 
    


    

}

// Decisiontablecolumns
type Decisiontablecolumns struct { 
    // Inputs - The input columns of the decision table.
    Inputs []Decisiontableinputcolumn `json:"inputs"`


    // Outputs - The output columns of the decision table.
    Outputs []Decisiontableoutputcolumn `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Decisiontablecolumns) String() string {
     o.Inputs = []Decisiontableinputcolumn{{}} 
     o.Outputs = []Decisiontableoutputcolumn{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablecolumns) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablecolumns

    if DecisiontablecolumnsMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablecolumnsMarshalled = true

    return json.Marshal(&struct {
        
        Inputs []Decisiontableinputcolumn `json:"inputs"`
        
        Outputs []Decisiontableoutputcolumn `json:"outputs"`
        *Alias
    }{

        
        Inputs: []Decisiontableinputcolumn{{}},
        


        
        Outputs: []Decisiontableoutputcolumn{{}},
        

        Alias: (*Alias)(u),
    })
}

