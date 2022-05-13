package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchactionpropertiesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchactionpropertiesDud struct { 
    


    


    


    


    

}

// Patchactionproperties
type Patchactionproperties struct { 
    // WebchatPrompt - Prompt message shown to user, used for webchat type action.
    WebchatPrompt string `json:"webchatPrompt"`


    // WebchatTitleText - Title shown to the user, used for webchat type action.
    WebchatTitleText string `json:"webchatTitleText"`


    // WebchatAcceptText - Accept button text shown to user, used for webchat type action.
    WebchatAcceptText string `json:"webchatAcceptText"`


    // WebchatDeclineText - Decline button text shown to user, used for webchat type action.
    WebchatDeclineText string `json:"webchatDeclineText"`


    // WebchatSurvey - Survey provided to the user, used for webchat type action.
    WebchatSurvey Patchactionsurvey `json:"webchatSurvey"`

}

// String returns a JSON representation of the model
func (o *Patchactionproperties) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchactionproperties) MarshalJSON() ([]byte, error) {
    type Alias Patchactionproperties

    if PatchactionpropertiesMarshalled {
        return []byte("{}"), nil
    }
    PatchactionpropertiesMarshalled = true

    return json.Marshal(&struct {
        
        WebchatPrompt string `json:"webchatPrompt"`
        
        WebchatTitleText string `json:"webchatTitleText"`
        
        WebchatAcceptText string `json:"webchatAcceptText"`
        
        WebchatDeclineText string `json:"webchatDeclineText"`
        
        WebchatSurvey Patchactionsurvey `json:"webchatSurvey"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

