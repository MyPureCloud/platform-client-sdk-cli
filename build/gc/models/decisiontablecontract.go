package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontablecontractMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontablecontractDud struct { 
    


    


    


    

}

// Decisiontablecontract
type Decisiontablecontract struct { 
    // ParentSchema - DSS schema entity defining source properties for the decision table contract schemas
    ParentSchema Domainentityref `json:"parentSchema"`


    // RowAuthoringSchema - JSON schema describing required value types for each column in every row in a decision table
    RowAuthoringSchema Contractjsonschema `json:"rowAuthoringSchema"`


    // ExecutionInputSchema - JSON schema for execution input data for a decision table
    ExecutionInputSchema Contractjsonschema `json:"executionInputSchema"`


    // ExecutionOutputSchema - JSON schema for execution output data for a decision table
    ExecutionOutputSchema Contractjsonschema `json:"executionOutputSchema"`

}

// String returns a JSON representation of the model
func (o *Decisiontablecontract) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontablecontract) MarshalJSON() ([]byte, error) {
    type Alias Decisiontablecontract

    if DecisiontablecontractMarshalled {
        return []byte("{}"), nil
    }
    DecisiontablecontractMarshalled = true

    return json.Marshal(&struct {
        
        ParentSchema Domainentityref `json:"parentSchema"`
        
        RowAuthoringSchema Contractjsonschema `json:"rowAuthoringSchema"`
        
        ExecutionInputSchema Contractjsonschema `json:"executionInputSchema"`
        
        ExecutionOutputSchema Contractjsonschema `json:"executionOutputSchema"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

