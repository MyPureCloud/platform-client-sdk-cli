package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QualityauditlogmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QualityauditlogmessageDud struct { 
    


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    

}

// Qualityauditlogmessage
type Qualityauditlogmessage struct { 
    // Id - Id of the audit message.
    Id string `json:"id"`


    // UserHomeOrgId - Home Organization Id associated with this audit message.
    UserHomeOrgId string `json:"userHomeOrgId"`


    // UserTrusteeOrgId - Trustee Organization Id if this audit message is from trustee access.
    UserTrusteeOrgId string `json:"userTrusteeOrgId"`


    // User - User associated with this audit message.
    User Domainentityref `json:"user"`


    // Client - Client associated with this audit message.
    Client Addressableentityref `json:"client"`


    // RemoteIps - List of IP addresses of systems that originated or handled the request.
    RemoteIps []string `json:"remoteIps"`


    // ServiceName - Name of the service that logged this audit message.
    ServiceName string `json:"serviceName"`


    // Level - The level of this audit message.
    Level string `json:"level"`


    // Status - The status of the action of this audit message.
    Status string `json:"status"`


    // EventDate - Date and time of when the audit message was logged. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    EventDate time.Time `json:"eventDate"`


    // MessageInfo - Message describing the event being audited.
    MessageInfo Messageinfo `json:"messageInfo"`


    // Action - Action that took place.
    Action string `json:"action"`


    // Entity - Entity that was impacted.
    Entity Domainentityref `json:"entity"`


    // EntityType - Type of the entity that was impacted.
    EntityType string `json:"entityType"`


    // PropertyChanges - List of properties that were changed and changes made to those properties.
    PropertyChanges []Propertychange `json:"propertyChanges"`


    // Context - Additional context for this message.
    Context map[string]string `json:"context"`

}

// String returns a JSON representation of the model
func (o *Qualityauditlogmessage) String() string {
    
    
    
    
    
     o.RemoteIps = []string{""} 
    
    
    
    
    
    
    
    
     o.PropertyChanges = []Propertychange{{}} 
     o.Context = map[string]string{"": ""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Qualityauditlogmessage) MarshalJSON() ([]byte, error) {
    type Alias Qualityauditlogmessage

    if QualityauditlogmessageMarshalled {
        return []byte("{}"), nil
    }
    QualityauditlogmessageMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        UserHomeOrgId string `json:"userHomeOrgId"`
        
        UserTrusteeOrgId string `json:"userTrusteeOrgId"`
        
        User Domainentityref `json:"user"`
        
        Client Addressableentityref `json:"client"`
        
        RemoteIps []string `json:"remoteIps"`
        
        ServiceName string `json:"serviceName"`
        
        Level string `json:"level"`
        
        Status string `json:"status"`
        
        EventDate time.Time `json:"eventDate"`
        
        MessageInfo Messageinfo `json:"messageInfo"`
        
        Action string `json:"action"`
        
        Entity Domainentityref `json:"entity"`
        
        EntityType string `json:"entityType"`
        
        PropertyChanges []Propertychange `json:"propertyChanges"`
        
        Context map[string]string `json:"context"`
        *Alias
    }{

        


        


        


        


        


        
        RemoteIps: []string{""},
        


        


        


        


        


        


        


        


        


        
        PropertyChanges: []Propertychange{{}},
        


        
        Context: map[string]string{"": ""},
        

        Alias: (*Alias)(u),
    })
}

