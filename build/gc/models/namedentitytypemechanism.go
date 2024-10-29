package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NamedentitytypemechanismMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NamedentitytypemechanismDud struct { 
    


    


    


    


    


    

}

// Namedentitytypemechanism
type Namedentitytypemechanism struct { 
    // Items - The items that define the named entity type.
    Items []Namedentitytypeitem `json:"items"`


    // Restricted - Whether the named entity type is restricted to the items provided. Default: false
    Restricted bool `json:"restricted"`


    // VarType - The type of the mechanism.
    VarType string `json:"type"`


    // SubType - Subtype of detection mechanism
    SubType string `json:"subType"`


    // MaxLength - The maximum length of the entity resolved value
    MaxLength int `json:"maxLength"`


    // Examples - Examples for entity detection
    Examples []Namedentitytypemechanismexample `json:"examples"`

}

// String returns a JSON representation of the model
func (o *Namedentitytypemechanism) String() string {
     o.Items = []Namedentitytypeitem{{}} 
    
    
    
    
     o.Examples = []Namedentitytypemechanismexample{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Namedentitytypemechanism) MarshalJSON() ([]byte, error) {
    type Alias Namedentitytypemechanism

    if NamedentitytypemechanismMarshalled {
        return []byte("{}"), nil
    }
    NamedentitytypemechanismMarshalled = true

    return json.Marshal(&struct {
        
        Items []Namedentitytypeitem `json:"items"`
        
        Restricted bool `json:"restricted"`
        
        VarType string `json:"type"`
        
        SubType string `json:"subType"`
        
        MaxLength int `json:"maxLength"`
        
        Examples []Namedentitytypemechanismexample `json:"examples"`
        *Alias
    }{

        
        Items: []Namedentitytypeitem{{}},
        


        


        


        


        


        
        Examples: []Namedentitytypemechanismexample{{}},
        

        Alias: (*Alias)(u),
    })
}

