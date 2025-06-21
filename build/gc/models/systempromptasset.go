package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SystempromptassetMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SystempromptassetDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    UploadUri string `json:"uploadUri"`


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Systempromptasset
type Systempromptasset struct { 
    


    // Name
    Name string `json:"name"`


    // PromptId
    PromptId string `json:"promptId"`


    // Language - The asset resource language
    Language string `json:"language"`


    // DurationSeconds
    DurationSeconds float64 `json:"durationSeconds"`


    // MediaUri
    MediaUri string `json:"mediaUri"`


    // TtsString
    TtsString string `json:"ttsString"`


    // Text
    Text string `json:"text"`


    


    // UploadStatus
    UploadStatus string `json:"uploadStatus"`


    // HasDefault
    HasDefault bool `json:"hasDefault"`


    // LanguageDefault
    LanguageDefault bool `json:"languageDefault"`


    // Tags
    Tags map[string][]string `json:"tags"`


    

}

// String returns a JSON representation of the model
func (o *Systempromptasset) String() string {
    
    
    
    
    
    
    
    
    
    
     o.Tags = map[string][]string{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Systempromptasset) MarshalJSON() ([]byte, error) {
    type Alias Systempromptasset

    if SystempromptassetMarshalled {
        return []byte("{}"), nil
    }
    SystempromptassetMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        PromptId string `json:"promptId"`
        
        Language string `json:"language"`
        
        DurationSeconds float64 `json:"durationSeconds"`
        
        MediaUri string `json:"mediaUri"`
        
        TtsString string `json:"ttsString"`
        
        Text string `json:"text"`
        
        UploadStatus string `json:"uploadStatus"`
        
        HasDefault bool `json:"hasDefault"`
        
        LanguageDefault bool `json:"languageDefault"`
        
        Tags map[string][]string `json:"tags"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        
        Tags: map[string][]string{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

