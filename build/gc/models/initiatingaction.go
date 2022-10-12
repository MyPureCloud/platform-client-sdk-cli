package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InitiatingactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InitiatingactionDud struct { 
    


    

}

// Initiatingaction
type Initiatingaction struct { 
    // TransactionId - Id of the audit initiating the transaction
    TransactionId string `json:"transactionId"`


    // ActionContext - Action of the audit initiating the transaction
    ActionContext string `json:"actionContext"`

}

// String returns a JSON representation of the model
func (o *Initiatingaction) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Initiatingaction) MarshalJSON() ([]byte, error) {
    type Alias Initiatingaction

    if InitiatingactionMarshalled {
        return []byte("{}"), nil
    }
    InitiatingactionMarshalled = true

    return json.Marshal(&struct {
        
        TransactionId string `json:"transactionId"`
        
        ActionContext string `json:"actionContext"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

