package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SentimentfeedbackMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SentimentfeedbackDud struct { 
    Id string `json:"id"`


    


    


    


    DateCreated time.Time `json:"dateCreated"`


    CreatedBy Addressableentityref `json:"createdBy"`

}

// Sentimentfeedback
type Sentimentfeedback struct { 
    


    // Phrase - The phrase for which sentiment feedback is provided
    Phrase string `json:"phrase"`


    // Dialect - The dialect for the given phrase, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard
    Dialect string `json:"dialect"`


    // FeedbackValue - The sentiment feedback value for the given phrase
    FeedbackValue string `json:"feedbackValue"`


    


    

}

// String returns a JSON representation of the model
func (o *Sentimentfeedback) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sentimentfeedback) MarshalJSON() ([]byte, error) {
    type Alias Sentimentfeedback

    if SentimentfeedbackMarshalled {
        return []byte("{}"), nil
    }
    SentimentfeedbackMarshalled = true

    return json.Marshal(&struct { 
        
        
        Phrase string `json:"phrase"`
        
        Dialect string `json:"dialect"`
        
        FeedbackValue string `json:"feedbackValue"`
        
        
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

