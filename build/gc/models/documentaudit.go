package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentauditMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentauditDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Documentaudit
type Documentaudit struct { 
    


    // Name
    Name string `json:"name"`


    // User
    User Domainentityref `json:"user"`


    // Workspace
    Workspace Domainentityref `json:"workspace"`


    // TransactionId
    TransactionId string `json:"transactionId"`


    // TransactionInitiator
    TransactionInitiator bool `json:"transactionInitiator"`


    // Application
    Application string `json:"application"`


    // ServiceName
    ServiceName string `json:"serviceName"`


    // Level
    Level string `json:"level"`


    // Timestamp - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Timestamp time.Time `json:"timestamp"`


    // Status
    Status string `json:"status"`


    // ActionContext
    ActionContext string `json:"actionContext"`


    // Action
    Action string `json:"action"`


    // Entity
    Entity Auditentityreference `json:"entity"`


    // Changes
    Changes []Auditchange `json:"changes"`


    

}

// String returns a JSON representation of the model
func (o *Documentaudit) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Changes = []Auditchange{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentaudit) MarshalJSON() ([]byte, error) {
    type Alias Documentaudit

    if DocumentauditMarshalled {
        return []byte("{}"), nil
    }
    DocumentauditMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        User Domainentityref `json:"user"`
        
        Workspace Domainentityref `json:"workspace"`
        
        TransactionId string `json:"transactionId"`
        
        TransactionInitiator bool `json:"transactionInitiator"`
        
        Application string `json:"application"`
        
        ServiceName string `json:"serviceName"`
        
        Level string `json:"level"`
        
        Timestamp time.Time `json:"timestamp"`
        
        Status string `json:"status"`
        
        ActionContext string `json:"actionContext"`
        
        Action string `json:"action"`
        
        Entity Auditentityreference `json:"entity"`
        
        Changes []Auditchange `json:"changes"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Changes: []Auditchange{{}},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

