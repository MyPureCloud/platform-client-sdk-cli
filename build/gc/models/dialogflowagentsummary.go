package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DialogflowagentsummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DialogflowagentsummaryDud struct { 
    Id string `json:"id"`


    


    


    


    SelfUri string `json:"selfUri"`

}

// Dialogflowagentsummary
type Dialogflowagentsummary struct { 
    


    // Name
    Name string `json:"name"`


    // Project - The project this Dialogflow agent belongs to
    Project Dialogflowproject `json:"project"`


    // Description - A description of the Dialogflow agent
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Dialogflowagentsummary) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Dialogflowagentsummary) MarshalJSON() ([]byte, error) {
    type Alias Dialogflowagentsummary

    if DialogflowagentsummaryMarshalled {
        return []byte("{}"), nil
    }
    DialogflowagentsummaryMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Project Dialogflowproject `json:"project"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

