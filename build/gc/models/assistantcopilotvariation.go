package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AssistantcopilotvariationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AssistantcopilotvariationDud struct { 
    Id string `json:"id"`


    


    DateCreated time.Time `json:"dateCreated"`


    DateModified time.Time `json:"dateModified"`


    CreatedBy Userreference `json:"createdBy"`


    ModifiedBy Userreference `json:"modifiedBy"`


    


    


    


    State string `json:"state"`


    


    SelfUri string `json:"selfUri"`


    VariationParent Addressableentityref `json:"variationParent"`

}

// Assistantcopilotvariation
type Assistantcopilotvariation struct { 
    


    // Name - The name of the assistant that will assist the agent.
    Name string `json:"name"`


    


    


    


    


    // GoogleDialogflowConfig - (Deprecated: use the 'knowledgeSuggestionConfig' for genesys knowledge suggestions) Configuration of Dialogflow used to assist the agent with transcriptions and knowledge suggestions.
    GoogleDialogflowConfig Googledialogflowconfig `json:"googleDialogflowConfig"`


    // TranscriptionConfig - Configuration for speech transcription used to assist the agent.
    TranscriptionConfig Transcriptionconfig `json:"transcriptionConfig"`


    // KnowledgeSuggestionConfig - Configuration that defines how to produce knowledge suggestions.
    KnowledgeSuggestionConfig Knowledgesuggestionconfig `json:"knowledgeSuggestionConfig"`


    


    // Copilot - Agent copilot configuration.
    Copilot Copilot `json:"copilot"`


    


    

}

// String returns a JSON representation of the model
func (o *Assistantcopilotvariation) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Assistantcopilotvariation) MarshalJSON() ([]byte, error) {
    type Alias Assistantcopilotvariation

    if AssistantcopilotvariationMarshalled {
        return []byte("{}"), nil
    }
    AssistantcopilotvariationMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        GoogleDialogflowConfig Googledialogflowconfig `json:"googleDialogflowConfig"`
        
        TranscriptionConfig Transcriptionconfig `json:"transcriptionConfig"`
        
        KnowledgeSuggestionConfig Knowledgesuggestionconfig `json:"knowledgeSuggestionConfig"`
        
        Copilot Copilot `json:"copilot"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

