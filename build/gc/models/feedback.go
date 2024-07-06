package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FeedbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FeedbackDud struct { 
    


    


    

}

// Feedback
type Feedback struct { 
    // SuggestionId - Feedback suggestion id.
    SuggestionId string `json:"suggestionId"`


    // UserProvided - Indicates whether the answer/item was clicked by the human agent or not.
    UserProvided bool `json:"userProvided"`


    // Relevance - Feedback relevance.
    Relevance string `json:"relevance"`

}

// String returns a JSON representation of the model
func (o *Feedback) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Feedback) MarshalJSON() ([]byte, error) {
    type Alias Feedback

    if FeedbackMarshalled {
        return []byte("{}"), nil
    }
    FeedbackMarshalled = true

    return json.Marshal(&struct {
        
        SuggestionId string `json:"suggestionId"`
        
        UserProvided bool `json:"userProvided"`
        
        Relevance string `json:"relevance"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

