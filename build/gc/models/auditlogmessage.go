package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuditlogmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuditlogmessageDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Auditlogmessage
type Auditlogmessage struct { 
    // Id - Id of the audit message.
    Id string `json:"id"`


    // UserHomeOrgId - Home Organization Id associated with this audit message.
    UserHomeOrgId string `json:"userHomeOrgId"`


    // User - User associated with this audit message.
    User Domainentityref `json:"user"`


    // Client - Client associated with this audit message.
    Client Addressableentityref `json:"client"`


    // RemoteIp - List of IP addresses of systems that originated or handled the request.
    RemoteIp []string `json:"remoteIp"`


    // ServiceName - Name of the service that logged this audit message.
    ServiceName string `json:"serviceName"`


    // Level - Level of this audit message, USER or SYSTEM.
    Level string `json:"level"`


    // EventDate - Date and time of when the audit message was logged. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDate time.Time `json:"eventDate"`


    // Message - Message describing the event being audited.
    Message Messageinfo `json:"message"`


    // Action - Action that took place.
    Action string `json:"action"`


    // Entity - Entity that was impacted.
    Entity Domainentityref `json:"entity"`


    // EntityType - Type of the entity that was impacted.
    EntityType string `json:"entityType"`


    // Status - Status of the event being audited
    Status string `json:"status"`


    // Application - Name of the application used to perform the audit's action
    Application string `json:"application"`


    // InitiatingAction - Id and action of the audit initiating the transaction
    InitiatingAction Initiatingaction `json:"initiatingAction"`


    // TransactionInitiator - Whether the current audit is the initiator of the transaction
    TransactionInitiator bool `json:"transactionInitiator"`


    // PropertyChanges - List of properties that were changed and changes made to those properties.
    PropertyChanges []Propertychange `json:"propertyChanges"`


    // Context - Additional context for this message.
    Context map[string]string `json:"context"`


    // EntityChanges - List of entities that were changed and changes made to those entities.
    EntityChanges []Entitychange `json:"entityChanges"`

}

// String returns a JSON representation of the model
func (o *Auditlogmessage) String() string {
    
    
    
    
     o.RemoteIp = []string{""} 
    
    
    
    
    
    
    
    
    
    
    
     o.PropertyChanges = []Propertychange{{}} 
     o.Context = map[string]string{"": ""} 
     o.EntityChanges = []Entitychange{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Auditlogmessage) MarshalJSON() ([]byte, error) {
    type Alias Auditlogmessage

    if AuditlogmessageMarshalled {
        return []byte("{}"), nil
    }
    AuditlogmessageMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        UserHomeOrgId string `json:"userHomeOrgId"`
        
        User Domainentityref `json:"user"`
        
        Client Addressableentityref `json:"client"`
        
        RemoteIp []string `json:"remoteIp"`
        
        ServiceName string `json:"serviceName"`
        
        Level string `json:"level"`
        
        EventDate time.Time `json:"eventDate"`
        
        Message Messageinfo `json:"message"`
        
        Action string `json:"action"`
        
        Entity Domainentityref `json:"entity"`
        
        EntityType string `json:"entityType"`
        
        Status string `json:"status"`
        
        Application string `json:"application"`
        
        InitiatingAction Initiatingaction `json:"initiatingAction"`
        
        TransactionInitiator bool `json:"transactionInitiator"`
        
        PropertyChanges []Propertychange `json:"propertyChanges"`
        
        Context map[string]string `json:"context"`
        
        EntityChanges []Entitychange `json:"entityChanges"`
        *Alias
    }{

        


        


        


        


        
        RemoteIp: []string{""},
        


        


        


        


        


        


        


        


        


        


        


        


        
        PropertyChanges: []Propertychange{{}},
        


        
        Context: map[string]string{"": ""},
        


        
        EntityChanges: []Entitychange{{}},
        

        Alias: (*Alias)(u),
    })
}

