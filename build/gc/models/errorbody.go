package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ErrorbodyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ErrorbodyDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Errorbody
type Errorbody struct { 
    // Message
    Message string `json:"message"`


    // Code
    Code string `json:"code"`


    // Status
    Status int `json:"status"`


    // EntityId
    EntityId string `json:"entityId"`


    // EntityName
    EntityName string `json:"entityName"`


    // MessageWithParams
    MessageWithParams string `json:"messageWithParams"`


    // MessageParams
    MessageParams map[string]string `json:"messageParams"`


    // ContextId
    ContextId string `json:"contextId"`


    // Details
    Details []Detail `json:"details"`


    // Errors
    Errors []Errorbody `json:"errors"`

}

// String returns a JSON representation of the model
func (o *Errorbody) String() string {
    
    
    
    
    
    
     o.MessageParams = map[string]string{"": ""} 
    
     o.Details = []Detail{{}} 
     o.Errors = []Errorbody{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Errorbody) MarshalJSON() ([]byte, error) {
    type Alias Errorbody

    if ErrorbodyMarshalled {
        return []byte("{}"), nil
    }
    ErrorbodyMarshalled = true

    return json.Marshal(&struct {
        
        Message string `json:"message"`
        
        Code string `json:"code"`
        
        Status int `json:"status"`
        
        EntityId string `json:"entityId"`
        
        EntityName string `json:"entityName"`
        
        MessageWithParams string `json:"messageWithParams"`
        
        MessageParams map[string]string `json:"messageParams"`
        
        ContextId string `json:"contextId"`
        
        Details []Detail `json:"details"`
        
        Errors []Errorbody `json:"errors"`
        *Alias
    }{

        


        


        


        


        


        


        
        MessageParams: map[string]string{"": ""},
        


        


        
        Details: []Detail{{}},
        


        
        Errors: []Errorbody{{}},
        

        Alias: (*Alias)(u),
    })
}

