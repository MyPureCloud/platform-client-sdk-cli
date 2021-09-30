package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UsageexecutionresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UsageexecutionresultDud struct { 
    


    

}

// Usageexecutionresult
type Usageexecutionresult struct { 
    // ExecutionId - The id of the query execution
    ExecutionId string `json:"executionId"`


    // ResultsUri - URI where the query results can be retrieved
    ResultsUri string `json:"resultsUri"`

}

// String returns a JSON representation of the model
func (o *Usageexecutionresult) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Usageexecutionresult) MarshalJSON() ([]byte, error) {
    type Alias Usageexecutionresult

    if UsageexecutionresultMarshalled {
        return []byte("{}"), nil
    }
    UsageexecutionresultMarshalled = true

    return json.Marshal(&struct { 
        ExecutionId string `json:"executionId"`
        
        ResultsUri string `json:"resultsUri"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

