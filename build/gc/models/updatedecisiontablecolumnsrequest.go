package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatedecisiontablecolumnsrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatedecisiontablecolumnsrequestDud struct { 
    


    

}

// Updatedecisiontablecolumnsrequest - Decision table columns will equal the columns contained in this request after the update is performed. If a list of columns types is not provided (i.e. is null) then no update is performed for that column type
type Updatedecisiontablecolumnsrequest struct { 
    // Inputs - The input columns of the decision table.
    Inputs []Decisiontableinputcolumnrequest `json:"inputs"`


    // Outputs - The output columns of the decision table.
    Outputs []Decisiontableoutputcolumnrequest `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Updatedecisiontablecolumnsrequest) String() string {
     o.Inputs = []Decisiontableinputcolumnrequest{{}} 
     o.Outputs = []Decisiontableoutputcolumnrequest{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatedecisiontablecolumnsrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatedecisiontablecolumnsrequest

    if UpdatedecisiontablecolumnsrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatedecisiontablecolumnsrequestMarshalled = true

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

