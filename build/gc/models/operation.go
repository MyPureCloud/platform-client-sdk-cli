package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperationDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Operation
type Operation struct { 
    // Id
    Id string `json:"id"`


    // Complete
    Complete bool `json:"complete"`


    // User
    User User `json:"user"`


    // Client
    Client Domainentityref `json:"client"`


    // ErrorMessage
    ErrorMessage string `json:"errorMessage"`


    // ErrorCode
    ErrorCode string `json:"errorCode"`


    // ErrorDetails
    ErrorDetails []Detail `json:"errorDetails"`


    // ErrorMessageParams
    ErrorMessageParams map[string]string `json:"errorMessageParams"`


    // ActionName - Action name
    ActionName string `json:"actionName"`


    // ActionStatus - Action status
    ActionStatus string `json:"actionStatus"`

}

// String returns a JSON representation of the model
func (o *Operation) String() string {
    
    
    
    
    
    
     o.ErrorDetails = []Detail{{}} 
     o.ErrorMessageParams = map[string]string{"": ""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operation) MarshalJSON() ([]byte, error) {
    type Alias Operation

    if OperationMarshalled {
        return []byte("{}"), nil
    }
    OperationMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Complete bool `json:"complete"`
        
        User User `json:"user"`
        
        Client Domainentityref `json:"client"`
        
        ErrorMessage string `json:"errorMessage"`
        
        ErrorCode string `json:"errorCode"`
        
        ErrorDetails []Detail `json:"errorDetails"`
        
        ErrorMessageParams map[string]string `json:"errorMessageParams"`
        
        ActionName string `json:"actionName"`
        
        ActionStatus string `json:"actionStatus"`
        *Alias
    }{

        


        


        


        


        


        


        
        ErrorDetails: []Detail{{}},
        


        
        ErrorMessageParams: map[string]string{"": ""},
        


        


        

        Alias: (*Alias)(u),
    })
}

