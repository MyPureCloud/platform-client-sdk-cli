package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowagentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowagentDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dialogflowagent
type Dialogflowagent struct { 
    


    // Name
    Name string `json:"name"`


    // Project - The project this Dialogflow agent belongs to
    Project Dialogflowproject `json:"project"`


    // Languages - The supported languages of the Dialogflow agent
    Languages []string `json:"languages"`


    // Intents - An array of Intents associated with this agent
    Intents []Dialogflowintent `json:"intents"`


    // Environments - Available environments for this agent
    Environments []string `json:"environments"`


    

}

// String returns a JSON representation of the model
func (o *Dialogflowagent) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.Languages = []string{""} 
    
    
    
     o.Intents = []Dialogflowintent{{}} 
    
    
    
     o.Environments = []string{""} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowagent) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowagent

    if DialogflowagentMarshalled {
        return []byte("{}"), nil
    }
    DialogflowagentMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        Project Dialogflowproject `json:"project"`
        
        Languages []string `json:"languages"`
        
        Intents []Dialogflowintent `json:"intents"`
        
        Environments []string `json:"environments"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Languages: []string{""},
        

        

        
        Intents: []Dialogflowintent{{}},
        

        

        
        Environments: []string{""},
        

        

        

        
        Alias: (*Alias)(u),
    })
}

