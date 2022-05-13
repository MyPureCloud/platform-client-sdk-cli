package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AuthzgrantroleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AuthzgrantroleDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Authzgrantrole
type Authzgrantrole struct { 
    


    // Name
    Name string `json:"name"`


    // Description
    Description string `json:"description"`


    // Policies
    Policies []Authzgrantpolicy `json:"policies"`


    // VarDefault
    VarDefault bool `json:"default"`


    

}

// String returns a JSON representation of the model
func (o *Authzgrantrole) String() string {
    
    
     o.Policies = []Authzgrantpolicy{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Authzgrantrole) MarshalJSON() ([]byte, error) {
    type Alias Authzgrantrole

    if AuthzgrantroleMarshalled {
        return []byte("{}"), nil
    }
    AuthzgrantroleMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        Policies []Authzgrantpolicy `json:"policies"`
        
        VarDefault bool `json:"default"`
        *Alias
    }{

        


        


        


        
        Policies: []Authzgrantpolicy{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

