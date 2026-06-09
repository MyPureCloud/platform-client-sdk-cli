package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DecisiontableimportrowmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DecisiontableimportrowmetricsDud struct { 
    


    


    


    


    


    


    


    


    

}

// Decisiontableimportrowmetrics - Progress metrics for a decision table import job
type Decisiontableimportrowmetrics struct { 
    // TotalRows - Total number of rows in the import file (set after parsing completes)
    TotalRows int `json:"totalRows"`


    // RowsParsed - Number of rows successfully parsed so far
    RowsParsed int `json:"rowsParsed"`


    // RowParseFailed - Number of rows that failed to parse
    RowParseFailed int `json:"rowParseFailed"`


    // RowsCreated - Number of rows successfully created so far
    RowsCreated int `json:"rowsCreated"`


    // RowsUpdated - Number of rows successfully updated so far
    RowsUpdated int `json:"rowsUpdated"`


    // RowsDeleted - Number of rows deleted (Replace mode only)
    RowsDeleted int `json:"rowsDeleted"`


    // RowCreateFailed - Number of rows that failed during batch create
    RowCreateFailed int `json:"rowCreateFailed"`


    // RowUpdateFailed - Number of rows that failed during batch update
    RowUpdateFailed int `json:"rowUpdateFailed"`


    // RowDeleteFailed - Number of rows that failed during delete
    RowDeleteFailed int `json:"rowDeleteFailed"`

}

// String returns a JSON representation of the model
func (o *Decisiontableimportrowmetrics) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Decisiontableimportrowmetrics) MarshalJSON() ([]byte, error) {
    type Alias Decisiontableimportrowmetrics

    if DecisiontableimportrowmetricsMarshalled {
        return []byte("{}"), nil
    }
    DecisiontableimportrowmetricsMarshalled = true

    return json.Marshal(&struct {
        
        TotalRows int `json:"totalRows"`
        
        RowsParsed int `json:"rowsParsed"`
        
        RowParseFailed int `json:"rowParseFailed"`
        
        RowsCreated int `json:"rowsCreated"`
        
        RowsUpdated int `json:"rowsUpdated"`
        
        RowsDeleted int `json:"rowsDeleted"`
        
        RowCreateFailed int `json:"rowCreateFailed"`
        
        RowUpdateFailed int `json:"rowUpdateFailed"`
        
        RowDeleteFailed int `json:"rowDeleteFailed"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

