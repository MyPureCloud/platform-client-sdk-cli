package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableexecutionresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableexecutionresponseDud struct { 
    


    


    


    

}

// Decisiontableexecutionresponse
type Decisiontableexecutionresponse struct { 
    // Table - The decision table version entity that was executed.
    Table Decisiontableversionentity `json:"table"`


    // TotalMatchRowCount - Total number of rows that matched execution input and would return results
    TotalMatchRowCount int `json:"totalMatchRowCount"`


    // TopMatchRows - Top 5 rows matching execution input, excluding the one produced the result.
    TopMatchRows []Decisiontablerowentityref `json:"topMatchRows"`


    // RowExecutionOutputs - The output data for each executed row for which output is collected.
    RowExecutionOutputs []Decisiontablerowexecutionoutput `json:"rowExecutionOutputs"`

}

// String returns a JSON representation of the model
func (o *Decisiontableexecutionresponse) String() string {
    
    
     o.TopMatchRows = []Decisiontablerowentityref{{}} 
     o.RowExecutionOutputs = []Decisiontablerowexecutionoutput{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableexecutionresponse) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableexecutionresponse

    if DecisiontableexecutionresponseMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableexecutionresponseMarshalled = true

    return json.Marshal(&struct {
        
        Table Decisiontableversionentity `json:"table"`
        
        TotalMatchRowCount int `json:"totalMatchRowCount"`
        
        TopMatchRows []Decisiontablerowentityref `json:"topMatchRows"`
        
        RowExecutionOutputs []Decisiontablerowexecutionoutput `json:"rowExecutionOutputs"`
        *Alias
    }{

        


        


        
        TopMatchRows: []Decisiontablerowentityref{{}},
        


        
        RowExecutionOutputs: []Decisiontablerowexecutionoutput{{}},
        

        Alias: (*Alias)(u),
    })
}

