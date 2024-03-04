package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowcxagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowcxagentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dialogflowcxagent
type Dialogflowcxagent struct { 
    


    // Name
    Name string `json:"name"`


    // Project - The project this Dialogflow CX agent belongs to.
    Project Dialogflowcxproject `json:"project"`


    // Languages - The supported languages of the Dialogflow CX agent.  Each value will be a language code in the country-locale format. e.g. en-us, es-us, fr-ca, etc.
    Languages []string `json:"languages"`


    // Environments - Available environments for this CX agent.
    Environments []Dialogflowcxenvironment `json:"environments"`


    // Integration - The Integration this Dialogflow CX agent was referenced from.
    Integration Domainentityref `json:"integration"`


    

}

// String returns a JSON representation of the model
func (o *Dialogflowcxagent) String() string {
    
    
     o.Languages = []string{""} 
     o.Environments = []Dialogflowcxenvironment{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowcxagent) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowcxagent

    if DialogflowcxagentMarshalled {
        return []byte("{}"), nil
    }
    DialogflowcxagentMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Project Dialogflowcxproject `json:"project"`
        
        Languages []string `json:"languages"`
        
        Environments []Dialogflowcxenvironment `json:"environments"`
        
        Integration Domainentityref `json:"integration"`
        *Alias
    }{

        


        


        


        
        Languages: []string{""},
        


        
        Environments: []Dialogflowcxenvironment{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

