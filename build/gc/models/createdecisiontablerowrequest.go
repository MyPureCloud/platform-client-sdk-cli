package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatedecisiontablerowrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatedecisiontablerowrequestDud struct { 
    


    


    

}

// Createdecisiontablerowrequest
type Createdecisiontablerowrequest struct { 
    // RowIndex - The absolute position of this row in the decision table. Must be an integerstarting from 0, must be non-negative and less than or equal to the size of the table. If not provided row will be append to the end of the table. 
    RowIndex int `json:"rowIndex"`


    // Inputs - The input values of the row. The key for this map is the column ID the row value relates. Column IDs are available from the decision table entity
    Inputs map[string]Decisiontablerowparametervalue `json:"inputs"`


    // Outputs - The output values of the row. The key for this map is the column ID the row value relates. Column IDs are available from the decision table entity
    Outputs map[string]Decisiontablerowparametervalue `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Createdecisiontablerowrequest) String() string {
    
     o.Inputs = map[string]Decisiontablerowparametervalue{"": {}} 
     o.Outputs = map[string]Decisiontablerowparametervalue{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createdecisiontablerowrequest) MarshalJSON() ([]byte, error) {
    type Alias Createdecisiontablerowrequest

    if CreatedecisiontablerowrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatedecisiontablerowrequestMarshalled = true

    return json.Marshal(&struct {
        
        RowIndex int `json:"rowIndex"`
        
        Inputs map[string]Decisiontablerowparametervalue `json:"inputs"`
        
        Outputs map[string]Decisiontablerowparametervalue `json:"outputs"`
        *Alias
    }{

        


        
        Inputs: map[string]Decisiontablerowparametervalue{"": {}},
        


        
        Outputs: map[string]Decisiontablerowparametervalue{"": {}},
        

        Alias: (*Alias)(u),
    })
}

