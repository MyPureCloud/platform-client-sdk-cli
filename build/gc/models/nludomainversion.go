package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    NludomainversionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type NludomainversionDud struct { 
    Id string `json:"id"`


    Domain *Nludomain `json:"domain"`


    


    


    Published bool `json:"published"`


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    DateTrained time.Time `json:"dateTrained"`


    DatePublished time.Time `json:"datePublished"`


    TrainingStatus string `json:"trainingStatus"`


    EvaluationStatus string `json:"evaluationStatus"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Nludomainversion
type Nludomainversion struct { 
    


    


    // Description - The description of the NLU domain version.
    Description string `json:"description"`


    // Language - The language that the NLU domain version supports.
    Language string `json:"language"`


    


    


    


    


    


    


    


    // Intents - The intents defined for this NLU domain version.
    Intents []Intentdefinition `json:"intents"`


    // EntityTypes - The entity types defined for this NLU domain version.
    EntityTypes []Namedentitytypedefinition `json:"entityTypes"`


    // Entities - The entities defined for this NLU domain version.This field is mutually exclusive with entityTypeBindings
    Entities []Namedentitydefinition `json:"entities"`


    

}

// String returns a JSON representation of the model
func (o *Nludomainversion) String() string {
    
    
     o.Intents = []Intentdefinition{{}} 
     o.EntityTypes = []Namedentitytypedefinition{{}} 
     o.Entities = []Namedentitydefinition{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Nludomainversion) MarshalJSON() ([]byte, error) {
    type Alias Nludomainversion

    if NludomainversionMarshalled {
        return []byte("{}"), nil
    }
    NludomainversionMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        Language string `json:"language"`
        
        Intents []Intentdefinition `json:"intents"`
        
        EntityTypes []Namedentitytypedefinition `json:"entityTypes"`
        
        Entities []Namedentitydefinition `json:"entities"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        
        Intents: []Intentdefinition{{}},
        


        
        EntityTypes: []Namedentitytypedefinition{{}},
        


        
        Entities: []Namedentitydefinition{{}},
        


        

        Alias: (*Alias)(u),
    })
}

