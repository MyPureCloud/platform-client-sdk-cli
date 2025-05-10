package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UpdatedecisiontablerowrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UpdatedecisiontablerowrequestDud struct { 
    


    

}

// Updatedecisiontablerowrequest
type Updatedecisiontablerowrequest struct { 
    // Inputs - The updated input values of the row. The key for this map is the column ID the row value relates. Column IDs are available from the decision table entity
    Inputs map[string]Decisiontablerowparametervalue `json:"inputs"`


    // Outputs - The updated output values of the row. The key for this map is the column ID the row value relates. Column IDs are available from the decision table entity
    Outputs map[string]Decisiontablerowparametervalue `json:"outputs"`

}

// String returns a JSON representation of the model
func (o *Updatedecisiontablerowrequest) String() string {
     o.Inputs = map[string]Decisiontablerowparametervalue{"": {}} 
     o.Outputs = map[string]Decisiontablerowparametervalue{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Updatedecisiontablerowrequest) MarshalJSON() ([]byte, error) {
    type Alias Updatedecisiontablerowrequest

    if UpdatedecisiontablerowrequestMarshalled {
        return []byte("{}"), nil
    }
    UpdatedecisiontablerowrequestMarshalled = true

    return json.Marshal(&struct {
        
        Inputs map[string]Decisiontablerowparametervalue `json:"inputs"`
        
        Outputs map[string]Decisiontablerowparametervalue `json:"outputs"`
        *Alias
    }{

        
        Inputs: map[string]Decisiontablerowparametervalue{"": {}},
        


        
        Outputs: map[string]Decisiontablerowparametervalue{"": {}},
        

        Alias: (*Alias)(u),
    })
}

