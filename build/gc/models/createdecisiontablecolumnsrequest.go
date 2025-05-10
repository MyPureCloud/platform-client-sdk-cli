package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatedecisiontablecolumnsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatedecisiontablecolumnsrequestDud struct { 
    


    

}

// Createdecisiontablecolumnsrequest
type Createdecisiontablecolumnsrequest struct { 
    // Inputs - The input columns of the decision table.
    Inputs []Decisiontableinputcolumnrequest `json:"inputs"`


    // Outputs - The output columns of the decision table.
    Outputs []Decisiontableoutputcolumnrequest `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Createdecisiontablecolumnsrequest) String() string {
     o.Inputs = []Decisiontableinputcolumnrequest{{}} 
     o.Outputs = []Decisiontableoutputcolumnrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createdecisiontablecolumnsrequest) MarshalJSON() ([]byte, error) {
    type Alias Createdecisiontablecolumnsrequest

    if CreatedecisiontablecolumnsrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatedecisiontablecolumnsrequestMarshalled = true

    return json.Marshal(&struct {
        
        Inputs []Decisiontableinputcolumnrequest `json:"inputs"`
        
        Outputs []Decisiontableoutputcolumnrequest `json:"outputs"`
        *Alias
    }{

        
        Inputs: []Decisiontableinputcolumnrequest{{}},
        


        
        Outputs: []Decisiontableoutputcolumnrequest{{}},
        

        Alias: (*Alias)(u),
    })
}

