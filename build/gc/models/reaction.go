package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ReactionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ReactionDud struct { 
    


    


    

}

// Reaction
type Reaction struct { 
    // Data - Parameter for this reaction. For transfer_flow, this would be the outbound flow id.
    Data string `json:"data"`


    // Name - Name of the parameter for this reaction. For transfer_flow, this would be the outbound flow name.
    Name string `json:"name"`


    // ReactionType - The reaction to take for a given call analysis result.
    ReactionType string `json:"reactionType"`

}

// String returns a JSON representation of the model
func (o *Reaction) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Reaction) MarshalJSON() ([]byte, error) {
    type Alias Reaction

    if ReactionMarshalled {
        return []byte("{}"), nil
    }
    ReactionMarshalled = true

    return json.Marshal(&struct {
        
        Data string `json:"data"`
        
        Name string `json:"name"`
        
        ReactionType string `json:"reactionType"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

