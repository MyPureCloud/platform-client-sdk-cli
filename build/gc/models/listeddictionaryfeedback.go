package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ListeddictionaryfeedbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ListeddictionaryfeedbackDud struct { 
    Id string `json:"id"`


    


    


    


    Source string `json:"source"`


    DateCreated time.Time `json:"dateCreated"`


    CreatedBy Userreference `json:"createdBy"`


    DateModified time.Time `json:"dateModified"`


    ModifiedBy Userreference `json:"modifiedBy"`


    


    Status string `json:"status"`


    


    SelfUri string `json:"selfUri"`

}

// Listeddictionaryfeedback
type Listeddictionaryfeedback struct { 
    


    // Term - The dictionary term which needs to be added to dictionary feedback system
    Term string `json:"term"`


    // Dialect - The dialect for the given term, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard
    Dialect string `json:"dialect"`


    // BoostValue - A weighted value assigned to a phrase. The higher the value, the higher the likelihood that the system will choose the word or phrase from the possible alternatives. Boost range is from 1.0 to 10.0. Default is 2.0
    BoostValue float32 `json:"boostValue"`


    


    


    


    


    


    // TranscriptionEngine - The transcription engine for the dictionary feedback. Only returned when GenesysExtended feature is enabled.
    TranscriptionEngine string `json:"transcriptionEngine"`


    


    // DisplayAs - The display name for the dictionary feedback. Only returned when GenesysExtended feature is enabled. This field is only valid for Extended Services transcription engine.
    DisplayAs string `json:"displayAs"`


    

}

// String returns a JSON representation of the model
func (o *Listeddictionaryfeedback) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Listeddictionaryfeedback) MarshalJSON() ([]byte, error) {
    type Alias Listeddictionaryfeedback

    if ListeddictionaryfeedbackMarshalled {
        return []byte("{}"), nil
    }
    ListeddictionaryfeedbackMarshalled = true

    return json.Marshal(&struct {
        
        Term string `json:"term"`
        
        Dialect string `json:"dialect"`
        
        BoostValue float32 `json:"boostValue"`
        
        TranscriptionEngine string `json:"transcriptionEngine"`
        
        DisplayAs string `json:"displayAs"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

