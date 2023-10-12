package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GrammarlanguageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GrammarlanguageDud struct { 
    Id string `json:"id"`


    


    


    VoiceFileUrl string `json:"voiceFileUrl"`


    DtmfFileUrl string `json:"dtmfFileUrl"`


    


    


    SelfUri string `json:"selfUri"`

}

// Grammarlanguage
type Grammarlanguage struct { 
    


    // GrammarId - The ID of the grammar associated with this grammar language
    GrammarId string `json:"grammarId"`


    // Language
    Language string `json:"language"`


    


    


    // VoiceFileMetadata - Additional information about the associated voice file
    VoiceFileMetadata Grammarlanguagefilemetadata `json:"voiceFileMetadata"`


    // DtmfFileMetadata - Additional information about the associated dtmf file
    DtmfFileMetadata Grammarlanguagefilemetadata `json:"dtmfFileMetadata"`


    

}

// String returns a JSON representation of the model
func (o *Grammarlanguage) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Grammarlanguage) MarshalJSON() ([]byte, error) {
    type Alias Grammarlanguage

    if GrammarlanguageMarshalled {
        return []byte("{}"), nil
    }
    GrammarlanguageMarshalled = true

    return json.Marshal(&struct {
        
        GrammarId string `json:"grammarId"`
        
        Language string `json:"language"`
        
        VoiceFileMetadata Grammarlanguagefilemetadata `json:"voiceFileMetadata"`
        
        DtmfFileMetadata Grammarlanguagefilemetadata `json:"dtmfFileMetadata"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

