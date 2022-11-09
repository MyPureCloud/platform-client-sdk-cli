package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntentdefinitionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntentdefinitionDud struct { 
    Id string `json:"id"`


    


    


    EntityNameReferences []string `json:"entityNameReferences"`


    

}

// Intentdefinition
type Intentdefinition struct { 
    


    // Name - The name of the intent.
    Name string `json:"name"`


    // EntityTypeBindings - The bindings for the named entity types used in this intent.This field is mutually exclusive with entityNameReferences and entities
    EntityTypeBindings []Namedentitytypebinding `json:"entityTypeBindings"`


    


    // Utterances - The utterances that act as training phrases for the intent.
    Utterances []Nluutterance `json:"utterances"`

}

// String returns a JSON representation of the model
func (o *Intentdefinition) String() string {
    
     o.EntityTypeBindings = []Namedentitytypebinding{{}} 
     o.Utterances = []Nluutterance{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intentdefinition) MarshalJSON() ([]byte, error) {
    type Alias Intentdefinition

    if IntentdefinitionMarshalled {
        return []byte("{}"), nil
    }
    IntentdefinitionMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        EntityTypeBindings []Namedentitytypebinding `json:"entityTypeBindings"`
        
        Utterances []Nluutterance `json:"utterances"`
        *Alias
    }{

        


        


        
        EntityTypeBindings: []Namedentitytypebinding{{}},
        


        


        
        Utterances: []Nluutterance{{}},
        

        Alias: (*Alias)(u),
    })
}

