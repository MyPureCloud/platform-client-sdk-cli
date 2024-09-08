package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    CreatedBy Userreference `json:"createdBy"`


    ModifiedBy Userreference `json:"modifiedBy"`


    


    


    


    State string `json:"state"`


    Copilot Copilot `json:"copilot"`


    SelfUri string `json:"selfUri"`

}

// Assistant
type Assistant struct { 
    


    // Name - The name of the assistant that will assist the agent.
    Name string `json:"name"`


    


    


    


    


    // GoogleDialogflowConfig - Configuration of Dialogflow used to assist the agent with transcriptions and knowledge suggestions.
    GoogleDialogflowConfig Googledialogflowconfig `json:"googleDialogflowConfig"`


    // TranscriptionConfig - Configuration for speech transcription used to assist the agent.
    TranscriptionConfig Transcriptionconfig `json:"transcriptionConfig"`


    // KnowledgeSuggestionConfig - Configuration that defines how to produce knowledge suggestions.
    KnowledgeSuggestionConfig Knowledgesuggestionconfig `json:"knowledgeSuggestionConfig"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Assistant) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistant) MarshalJSON() ([]byte, error) {
    type Alias Assistant

    if AssistantMarshalled {
        return []byte("{}"), nil
    }
    AssistantMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        GoogleDialogflowConfig Googledialogflowconfig `json:"googleDialogflowConfig"`
        
        TranscriptionConfig Transcriptionconfig `json:"transcriptionConfig"`
        
        KnowledgeSuggestionConfig Knowledgesuggestionconfig `json:"knowledgeSuggestionConfig"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

