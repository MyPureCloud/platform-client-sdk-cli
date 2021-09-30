package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DomainpermissionpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DomainpermissionpolicyDud struct { 
    


    


    


    


    


    


    


    

}

// Domainpermissionpolicy
type Domainpermissionpolicy struct { 
    // Domain
    Domain string `json:"domain"`


    // EntityName
    EntityName string `json:"entityName"`


    // PolicyName
    PolicyName string `json:"policyName"`


    // PolicyDescription
    PolicyDescription string `json:"policyDescription"`


    // ActionSet
    ActionSet []string `json:"actionSet"`


    // NamedResources
    NamedResources []string `json:"namedResources"`


    // AllowConditions
    AllowConditions bool `json:"allowConditions"`


    // ResourceConditionNode
    ResourceConditionNode Domainresourceconditionnode `json:"resourceConditionNode"`

}

// String returns a JSON representation of the model
func (o *Domainpermissionpolicy) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.ActionSet = []string{""} 
    
    
    
     o.NamedResources = []string{""} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Domainpermissionpolicy) MarshalJSON() ([]byte, error) {
    type Alias Domainpermissionpolicy

    if DomainpermissionpolicyMarshalled {
        return []byte("{}"), nil
    }
    DomainpermissionpolicyMarshalled = true

    return json.Marshal(&struct { 
        Domain string `json:"domain"`
        
        EntityName string `json:"entityName"`
        
        PolicyName string `json:"policyName"`
        
        PolicyDescription string `json:"policyDescription"`
        
        ActionSet []string `json:"actionSet"`
        
        NamedResources []string `json:"namedResources"`
        
        AllowConditions bool `json:"allowConditions"`
        
        ResourceConditionNode Domainresourceconditionnode `json:"resourceConditionNode"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        
        ActionSet: []string{""},
        

        

        
        NamedResources: []string{""},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

