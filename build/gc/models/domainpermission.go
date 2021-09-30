package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainpermissionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainpermissionDud struct { 
    


    


    


    


    


    

}

// Domainpermission
type Domainpermission struct { 
    // Domain
    Domain string `json:"domain"`


    // EntityType
    EntityType string `json:"entityType"`


    // Action
    Action string `json:"action"`


    // Label
    Label string `json:"label"`


    // AllowsConditions
    AllowsConditions bool `json:"allowsConditions"`


    // DivisionAware
    DivisionAware bool `json:"divisionAware"`

}

// String returns a JSON representation of the model
func (o *Domainpermission) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainpermission) MarshalJSON() ([]byte, error) {
    type Alias Domainpermission

    if DomainpermissionMarshalled {
        return []byte("{}"), nil
    }
    DomainpermissionMarshalled = true

    return json.Marshal(&struct { 
        Domain string `json:"domain"`
        
        EntityType string `json:"entityType"`
        
        Action string `json:"action"`
        
        Label string `json:"label"`
        
        AllowsConditions bool `json:"allowsConditions"`
        
        DivisionAware bool `json:"divisionAware"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

