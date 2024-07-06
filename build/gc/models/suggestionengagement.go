package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SuggestionengagementMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SuggestionengagementDud struct { 
    


    

}

// Suggestionengagement
type Suggestionengagement struct { 
    // EngagementType - The type of engagement with the suggestion.
    EngagementType string `json:"engagementType"`


    // Feedback - The given feedback on the suggestion, if any.
    Feedback Suggestionfeedback `json:"feedback"`

}

// String returns a JSON representation of the model
func (o *Suggestionengagement) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Suggestionengagement) MarshalJSON() ([]byte, error) {
    type Alias Suggestionengagement

    if SuggestionengagementMarshalled {
        return []byte("{}"), nil
    }
    SuggestionengagementMarshalled = true

    return json.Marshal(&struct {
        
        EngagementType string `json:"engagementType"`
        
        Feedback Suggestionfeedback `json:"feedback"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

