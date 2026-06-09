package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RowMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RowDud struct { 
    


    


    

}

// Row
type Row struct { 
    // RowId - The ID of the row to update
    RowId string `json:"rowId"`


    // Inputs - The input values of the row. The key for this map is the column ID the row value relates. Column IDs are available from the decision table entity.
    Inputs map[string]Decisiontablerowparametervalue `json:"inputs"`


    // Outputs - The output values of the row. The key for this map is the column ID the row value relates. Column IDs are available from the decision table entity.
    Outputs map[string]Decisiontablerowparametervalue `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Row) String() string {
    
     o.Inputs = map[string]Decisiontablerowparametervalue{"": {}} 
     o.Outputs = map[string]Decisiontablerowparametervalue{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Row) MarshalJSON() ([]byte, error) {
    type Alias Row

    if RowMarshalled {
        return []byte("{}"), nil
    }
    RowMarshalled = true

    return json.Marshal(&struct {
        
        RowId string `json:"rowId"`
        
        Inputs map[string]Decisiontablerowparametervalue `json:"inputs"`
        
        Outputs map[string]Decisiontablerowparametervalue `json:"outputs"`
        *Alias
    }{

        


        
        Inputs: map[string]Decisiontablerowparametervalue{"": {}},
        


        
        Outputs: map[string]Decisiontablerowparametervalue{"": {}},
        

        Alias: (*Alias)(u),
    })
}

