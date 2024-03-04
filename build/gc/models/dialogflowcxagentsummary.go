package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowcxagentsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowcxagentsummaryDud struct { 
    Id string `json:"id"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dialogflowcxagentsummary
type Dialogflowcxagentsummary struct { 
    


    // Name
    Name string `json:"name"`


    // Project - The project this Dialogflow CX agent belongs to.
    Project Dialogflowcxproject `json:"project"`


    // Description - A description of the Dialogflow CX agent.
    Description string `json:"description"`


    // Integration - The Integration this Dialogflow CX agent was referenced from.
    Integration Domainentityref `json:"integration"`


    

}

// String returns a JSON representation of the model
func (o *Dialogflowcxagentsummary) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowcxagentsummary) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowcxagentsummary

    if DialogflowcxagentsummaryMarshalled {
        return []byte("{}"), nil
    }
    DialogflowcxagentsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Project Dialogflowcxproject `json:"project"`
        
        Description string `json:"description"`
        
        Integration Domainentityref `json:"integration"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

