package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzgrantpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzgrantpolicyDud struct { 
    


    


    


    

}

// Authzgrantpolicy
type Authzgrantpolicy struct { 
    // Actions
    Actions []string `json:"actions"`


    // Condition
    Condition string `json:"condition"`


    // Domain
    Domain string `json:"domain"`


    // EntityName
    EntityName string `json:"entityName"`

}

// String returns a JSON representation of the model
func (o *Authzgrantpolicy) String() string {
     o.Actions = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzgrantpolicy) MarshalJSON() ([]byte, error) {
    type Alias Authzgrantpolicy

    if AuthzgrantpolicyMarshalled {
        return []byte("{}"), nil
    }
    AuthzgrantpolicyMarshalled = true

    return json.Marshal(&struct {
        
        Actions []string `json:"actions"`
        
        Condition string `json:"condition"`
        
        Domain string `json:"domain"`
        
        EntityName string `json:"entityName"`
        *Alias
    }{

        
        Actions: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

