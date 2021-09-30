package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActionpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActionpropertiesDud struct { 
    


    


    


    


    

}

// Actionproperties
type Actionproperties struct { 
    // WebchatPrompt - Prompt message shown to user, used for webchat type action.
    WebchatPrompt string `json:"webchatPrompt"`


    // WebchatTitleText - Title shown to the user, used for webchat type action.
    WebchatTitleText string `json:"webchatTitleText"`


    // WebchatAcceptText - Accept button text shown to user, used for webchat type action.
    WebchatAcceptText string `json:"webchatAcceptText"`


    // WebchatDeclineText - Decline button text shown to user, used for webchat type action.
    WebchatDeclineText string `json:"webchatDeclineText"`


    // WebchatSurvey - Survey provided to the user, used for webchat type action.
    WebchatSurvey Actionsurvey `json:"webchatSurvey"`

}

// String returns a JSON representation of the model
func (o *Actionproperties) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Actionproperties) MarshalJSON() ([]byte, error) {
    type Alias Actionproperties

    if ActionpropertiesMarshalled {
        return []byte("{}"), nil
    }
    ActionpropertiesMarshalled = true

    return json.Marshal(&struct { 
        WebchatPrompt string `json:"webchatPrompt"`
        
        WebchatTitleText string `json:"webchatTitleText"`
        
        WebchatAcceptText string `json:"webchatAcceptText"`
        
        WebchatDeclineText string `json:"webchatDeclineText"`
        
        WebchatSurvey Actionsurvey `json:"webchatSurvey"`
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

