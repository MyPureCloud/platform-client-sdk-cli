package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthorizationpolicyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthorizationpolicyDud struct { 
    Id string `json:"id"`


    


    TargetResource string `json:"targetResource"`


    


    


    


    


    DateModified time.Time `json:"dateModified"`


    


    


    SelfUri string `json:"selfUri"`

}

// Authorizationpolicy
type Authorizationpolicy struct { 
    


    // Name
    Name string `json:"name"`


    


    // Subject - The subject to whom the policy will apply, including type and id
    Subject Subject `json:"subject"`


    // Effect - The effect this policy should have when its conditions are met
    Effect string `json:"effect"`


    // Condition - The condition tree the policy will evaluate
    Condition interface{} `json:"condition"`


    // Description
    Description string `json:"description"`


    


    // PresetAttributes - Map of names and values of preset attributes to use in policy evaluation
    PresetAttributes map[string]Typedattribute `json:"presetAttributes"`


    // Active - Flag for active enforcement. If this value is false or null, the policy will be saved but will not be checked or enforced on users.
    Active bool `json:"active"`


    

}

// String returns a JSON representation of the model
func (o *Authorizationpolicy) String() string {
    
    
    
     o.Condition = Interface{} 
    
     o.PresetAttributes = map[string]Typedattribute{"": {}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authorizationpolicy) MarshalJSON() ([]byte, error) {
    type Alias Authorizationpolicy

    if AuthorizationpolicyMarshalled {
        return []byte("{}"), nil
    }
    AuthorizationpolicyMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Subject Subject `json:"subject"`
        
        Effect string `json:"effect"`
        
        Condition interface{} `json:"condition"`
        
        Description string `json:"description"`
        
        PresetAttributes map[string]Typedattribute `json:"presetAttributes"`
        
        Active bool `json:"active"`
        *Alias
    }{

        


        


        


        


        


        
        Condition: Interface{},
        


        


        


        
        PresetAttributes: map[string]Typedattribute{"": {}},
        


        


        

        Alias: (*Alias)(u),
    })
}

