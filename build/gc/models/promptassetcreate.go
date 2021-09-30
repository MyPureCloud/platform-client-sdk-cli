package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PromptassetcreateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PromptassetcreateDud struct { 
    Id string `json:"id"`


    


    PromptId string `json:"promptId"`


    


    MediaUri string `json:"mediaUri"`


    


    


    UploadStatus string `json:"uploadStatus"`


    UploadUri string `json:"uploadUri"`


    LanguageDefault bool `json:"languageDefault"`


    


    


    SelfUri string `json:"selfUri"`

}

// Promptassetcreate
type Promptassetcreate struct { 
    


    // Name
    Name string `json:"name"`


    


    // Language - The prompt language.
    Language string `json:"language"`


    


    // TtsString - Text to speech of the resource
    TtsString string `json:"ttsString"`


    // Text - Text of the resource
    Text string `json:"text"`


    


    


    


    // Tags
    Tags map[string][]string `json:"tags"`


    // DurationSeconds
    DurationSeconds float64 `json:"durationSeconds"`


    

}

// String returns a JSON representation of the model
func (o *Promptassetcreate) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
     o.Tags = map[string][]string{"": {}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Promptassetcreate) MarshalJSON() ([]byte, error) {
    type Alias Promptassetcreate

    if PromptassetcreateMarshalled {
        return []byte("{}"), nil
    }
    PromptassetcreateMarshalled = true

    return json.Marshal(&struct { 
        
        
        Name string `json:"name"`
        
        
        
        Language string `json:"language"`
        
        
        
        TtsString string `json:"ttsString"`
        
        Text string `json:"text"`
        
        
        
        
        
        
        
        Tags map[string][]string `json:"tags"`
        
        DurationSeconds float64 `json:"durationSeconds"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        

        
        Tags: map[string][]string{"": {}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

