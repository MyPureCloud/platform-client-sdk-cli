package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitieslanguagefeaturesMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitieslanguagefeaturesDud struct { 
    


    


    


    


    


    


    


    


    


    

}

// Capabilitieslanguagefeatures
type Capabilitieslanguagefeatures struct { 
    // CollectInput
    CollectInput Capabilitiesflowtypesavailability `json:"collectInput"`


    // Copilot
    Copilot Capabilitiescopilot `json:"copilot"`


    // EstimatedWaitTime
    EstimatedWaitTime Capabilitiesflowtypesavailability `json:"estimatedWaitTime"`


    // KnowledgeBase
    KnowledgeBase Capabilitiesknowledgebase `json:"knowledgeBase"`


    // PositionInQueue
    PositionInQueue Capabilitiesflowtypesavailability `json:"positionInQueue"`


    // RuntimeDataPlayback
    RuntimeDataPlayback Capabilitiesflowtypesavailability `json:"runtimeDataPlayback"`


    // SpeechRec
    SpeechRec Capabilitiesflowtypesavailability `json:"speechRec"`


    // TextToSpeech
    TextToSpeech Capabilitieslanguagetts `json:"textToSpeech"`


    // Transcription
    Transcription Capabilitiesflowtypesavailability `json:"transcription"`


    // Voicemail
    Voicemail map[string]Capabilitiesflowtypesavailability `json:"voicemail"`

}

// String returns a JSON representation of the model
func (o *Capabilitieslanguagefeatures) String() string {
    
    
    
    
    
    
    
    
    
     o.Voicemail = map[string]Capabilitiesflowtypesavailability{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitieslanguagefeatures) MarshalJSON() ([]byte, error) {
    type Alias Capabilitieslanguagefeatures

    if CapabilitieslanguagefeaturesMarshalled {
        return []byte("{}"), nil
    }
    CapabilitieslanguagefeaturesMarshalled = true

    return json.Marshal(&struct {
        
        CollectInput Capabilitiesflowtypesavailability `json:"collectInput"`
        
        Copilot Capabilitiescopilot `json:"copilot"`
        
        EstimatedWaitTime Capabilitiesflowtypesavailability `json:"estimatedWaitTime"`
        
        KnowledgeBase Capabilitiesknowledgebase `json:"knowledgeBase"`
        
        PositionInQueue Capabilitiesflowtypesavailability `json:"positionInQueue"`
        
        RuntimeDataPlayback Capabilitiesflowtypesavailability `json:"runtimeDataPlayback"`
        
        SpeechRec Capabilitiesflowtypesavailability `json:"speechRec"`
        
        TextToSpeech Capabilitieslanguagetts `json:"textToSpeech"`
        
        Transcription Capabilitiesflowtypesavailability `json:"transcription"`
        
        Voicemail map[string]Capabilitiesflowtypesavailability `json:"voicemail"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        
        Voicemail: map[string]Capabilitiesflowtypesavailability{"": {}},
        

        Alias: (*Alias)(u),
    })
}

