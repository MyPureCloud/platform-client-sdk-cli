package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditmessageDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Auditmessage
type Auditmessage struct { 
    // Id - AuditMessage ID.
    Id string `json:"id"`


    // User
    User Audituser `json:"user"`


    // CorrelationId - Correlation ID.
    CorrelationId string `json:"correlationId"`


    // TransactionId - Transaction ID.
    TransactionId string `json:"transactionId"`


    // TransactionInitiator - Whether or not this audit can be considered the initiator of the transaction it is a part of.
    TransactionInitiator bool `json:"transactionInitiator"`


    // Application - The application through which the action of this AuditMessage was initiated.
    Application string `json:"application"`


    // ServiceName - The name of the service which sent this AuditMessage.
    ServiceName string `json:"serviceName"`


    // Level - The level of this audit. USER or SYSTEM.
    Level string `json:"level"`


    // Timestamp - The time at which the action of this AuditMessage was initiated.
    Timestamp string `json:"timestamp"`


    // ReceivedTimestamp - The time at which this AuditMessage was received.
    ReceivedTimestamp string `json:"receivedTimestamp"`


    // Status - The status of the action of this AuditMessage
    Status string `json:"status"`


    // ActionContext - The context of a system-level action
    ActionContext string `json:"actionContext"`


    // Action - A string representing the action that took place
    Action string `json:"action"`


    // Changes - Details about any changes that occurred in this audit
    Changes []Change `json:"changes"`


    // Entity
    Entity Auditentity `json:"entity"`


    // ServiceContext - The service-specific context associated with this AuditMessage.
    ServiceContext Servicecontext `json:"serviceContext"`

}

// String returns a JSON representation of the model
func (o *Auditmessage) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Changes = []Change{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditmessage) MarshalJSON() ([]byte, error) {
    type Alias Auditmessage

    if AuditmessageMarshalled {
        return []byte("{}"), nil
    }
    AuditmessageMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        User Audituser `json:"user"`
        
        CorrelationId string `json:"correlationId"`
        
        TransactionId string `json:"transactionId"`
        
        TransactionInitiator bool `json:"transactionInitiator"`
        
        Application string `json:"application"`
        
        ServiceName string `json:"serviceName"`
        
        Level string `json:"level"`
        
        Timestamp string `json:"timestamp"`
        
        ReceivedTimestamp string `json:"receivedTimestamp"`
        
        Status string `json:"status"`
        
        ActionContext string `json:"actionContext"`
        
        Action string `json:"action"`
        
        Changes []Change `json:"changes"`
        
        Entity Auditentity `json:"entity"`
        
        ServiceContext Servicecontext `json:"serviceContext"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        


        
        Changes: []Change{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

