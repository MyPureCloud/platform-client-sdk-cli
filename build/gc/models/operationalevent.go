package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OperationaleventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OperationaleventDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Operationalevent
type Operationalevent struct { 
    // EventDefinition - The event that occurred.
    EventDefinition Addressableentityref `json:"eventDefinition"`


    // EntityId - The unique identifier for the entity
    EntityId string `json:"entityId"`


    // EntityToken - A token representing the entity
    EntityToken string `json:"entityToken"`


    // EntityName - The name for the entity
    EntityName string `json:"entityName"`


    // PreviousValue - The value prior to the event
    PreviousValue string `json:"previousValue"`


    // CurrentValue - The changed value following the event
    CurrentValue string `json:"currentValue"`


    // ErrorCode - The error code of the entity in the providing service associated to the event
    ErrorCode string `json:"errorCode"`


    // ParentEntityId - The unique identifier for the parent of the entity
    ParentEntityId string `json:"parentEntityId"`


    // Conversation - The link to a conversation
    Conversation Addressableentityref `json:"conversation"`


    // DateCreated - The date when the event created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    DateCreated time.Time `json:"dateCreated"`


    // EntityVersion - The version of the entity in the providing service
    EntityVersion string `json:"entityVersion"`

}

// String returns a JSON representation of the model
func (o *Operationalevent) String() string {
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Operationalevent) MarshalJSON() ([]byte, error) {
    type Alias Operationalevent

    if OperationaleventMarshalled {
        return []byte("{}"), nil
    }
    OperationaleventMarshalled = true

    return json.Marshal(&struct {
        
        EventDefinition Addressableentityref `json:"eventDefinition"`
        
        EntityId string `json:"entityId"`
        
        EntityToken string `json:"entityToken"`
        
        EntityName string `json:"entityName"`
        
        PreviousValue string `json:"previousValue"`
        
        CurrentValue string `json:"currentValue"`
        
        ErrorCode string `json:"errorCode"`
        
        ParentEntityId string `json:"parentEntityId"`
        
        Conversation Addressableentityref `json:"conversation"`
        
        DateCreated time.Time `json:"dateCreated"`
        
        EntityVersion string `json:"entityVersion"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

