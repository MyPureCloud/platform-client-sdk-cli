package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ConversationprofileMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ConversationprofileDud struct { 
    


    

}

// Conversationprofile
type Conversationprofile struct { 
    // LanguageCode - The language code supported by the conversation profile belonging to a particular project for Dialogflow.
    LanguageCode string `json:"languageCode"`


    // Name - The name of the conversation profile belonging to a particular project for Dialogflow
    Name string `json:"name"`

}

// String returns a JSON representation of the model
func (o *Conversationprofile) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Conversationprofile) MarshalJSON() ([]byte, error) {
    type Alias Conversationprofile

    if ConversationprofileMarshalled {
        return []byte("{}"), nil
    }
    ConversationprofileMarshalled = true

    return json.Marshal(&struct {
        
        LanguageCode string `json:"languageCode"`
        
        Name string `json:"name"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

