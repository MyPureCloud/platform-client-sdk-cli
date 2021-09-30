package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessageevaluationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessageevaluationDud struct { 
    


    


    


    


    

}

// Messageevaluation
type Messageevaluation struct { 
    // ContactColumn - The name of the contact column that was wrapped up
    ContactColumn string `json:"contactColumn"`


    // ContactAddress - The address (phone or email) that was wrapped up
    ContactAddress string `json:"contactAddress"`


    // MessageType - The type of message sent
    MessageType string `json:"messageType"`


    // WrapupCodeId - The id of the wrap-up code
    WrapupCodeId string `json:"wrapupCodeId"`


    // Timestamp - The time that the wrap-up was applied. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`

}

// String returns a JSON representation of the model
func (o *Messageevaluation) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messageevaluation) MarshalJSON() ([]byte, error) {
    type Alias Messageevaluation

    if MessageevaluationMarshalled {
        return []byte("{}"), nil
    }
    MessageevaluationMarshalled = true

    return json.Marshal(&struct { 
        ContactColumn string `json:"contactColumn"`
        
        ContactAddress string `json:"contactAddress"`
        
        MessageType string `json:"messageType"`
        
        WrapupCodeId string `json:"wrapupCodeId"`
        
        Timestamp time.Time `json:"timestamp"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

