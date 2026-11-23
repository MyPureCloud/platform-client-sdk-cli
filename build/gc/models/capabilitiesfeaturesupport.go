package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesfeaturesupportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesfeaturesupportDud struct { 
    


    


    


    


    


    

}

// Capabilitiesfeaturesupport
type Capabilitiesfeaturesupport struct { 
    // TextToSpeech - Architect features metadata text to speech
    TextToSpeech Capabilitiesttsconfiguration `json:"textToSpeech"`


    // SpeechToText - Architect features metadata speech to text
    SpeechToText Capabilitiessttconfiguration `json:"speechToText"`


    // NaturalLanguageUnderstanding - Architect features metadata natural language understanding
    NaturalLanguageUnderstanding Capabilitiesnlusupport `json:"naturalLanguageUnderstanding"`


    // Grammars - Architect features metadata grammars
    Grammars Capabilitiesgrammarssupport `json:"grammars"`


    // Workitems - Architect features metadata workitems
    Workitems Capabilitiesworkitemssupport `json:"workitems"`


    // Navigator - Architect features metadata navigator
    Navigator Capabilitiesnavigatorsupport `json:"navigator"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesfeaturesupport) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesfeaturesupport) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesfeaturesupport

    if CapabilitiesfeaturesupportMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesfeaturesupportMarshalled = true

    return json.Marshal(&struct {
        
        TextToSpeech Capabilitiesttsconfiguration `json:"textToSpeech"`
        
        SpeechToText Capabilitiessttconfiguration `json:"speechToText"`
        
        NaturalLanguageUnderstanding Capabilitiesnlusupport `json:"naturalLanguageUnderstanding"`
        
        Grammars Capabilitiesgrammarssupport `json:"grammars"`
        
        Workitems Capabilitiesworkitemssupport `json:"workitems"`
        
        Navigator Capabilitiesnavigatorsupport `json:"navigator"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

