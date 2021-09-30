package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ResourcepermissionpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ResourcepermissionpolicyDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Resourcepermissionpolicy
type Resourcepermissionpolicy struct { 
    // Id
    Id string `json:"id"`


    // Domain
    Domain string `json:"domain"`


    // EntityName
    EntityName string `json:"entityName"`


    // PolicyName
    PolicyName string `json:"policyName"`


    // PolicyDescription
    PolicyDescription string `json:"policyDescription"`


    // ActionSetKey
    ActionSetKey string `json:"actionSetKey"`


    // AllowConditions
    AllowConditions bool `json:"allowConditions"`


    // ResourceConditionNode
    ResourceConditionNode Resourceconditionnode `json:"resourceConditionNode"`


    // NamedResources
    NamedResources []string `json:"namedResources"`


    // ResourceCondition
    ResourceCondition string `json:"resourceCondition"`


    // ActionSet
    ActionSet []string `json:"actionSet"`

}

// String returns a JSON representation of the model
func (o *Resourcepermissionpolicy) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.NamedResources = []string{""} 
    
    
    
    
    
    
    
     o.ActionSet = []string{""} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Resourcepermissionpolicy) MarshalJSON() ([]byte, error) {
    type Alias Resourcepermissionpolicy

    if ResourcepermissionpolicyMarshalled {
        return []byte("{}"), nil
    }
    ResourcepermissionpolicyMarshalled = true

    return json.Marshal(&struct { 
        Id string `json:"id"`
        
        Domain string `json:"domain"`
        
        EntityName string `json:"entityName"`
        
        PolicyName string `json:"policyName"`
        
        PolicyDescription string `json:"policyDescription"`
        
        ActionSetKey string `json:"actionSetKey"`
        
        AllowConditions bool `json:"allowConditions"`
        
        ResourceConditionNode Resourceconditionnode `json:"resourceConditionNode"`
        
        NamedResources []string `json:"namedResources"`
        
        ResourceCondition string `json:"resourceCondition"`
        
        ActionSet []string `json:"actionSet"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        NamedResources: []string{""},
        

        

        

        

        
        ActionSet: []string{""},
        

        
        Alias: (*Alias)(u),
    })
}

