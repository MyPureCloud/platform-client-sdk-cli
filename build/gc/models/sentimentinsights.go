package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SentimentinsightsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SentimentinsightsDud struct { 
    


    

}

// Sentimentinsights
type Sentimentinsights struct { 
    // PositiveSentimentReasons - The reasons for positive sentiment found in the conversation
    PositiveSentimentReasons []Sentimentinsightentry `json:"positiveSentimentReasons"`


    // NegativeSentimentReasons - The reasons for negative sentiment found in the conversation
    NegativeSentimentReasons []Sentimentinsightentry `json:"negativeSentimentReasons"`

}

// String returns a JSON representation of the model
func (o *Sentimentinsights) String() string {
     o.PositiveSentimentReasons = []Sentimentinsightentry{{}} 
     o.NegativeSentimentReasons = []Sentimentinsightentry{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sentimentinsights) MarshalJSON() ([]byte, error) {
    type Alias Sentimentinsights

    if SentimentinsightsMarshalled {
        return []byte("{}"), nil
    }
    SentimentinsightsMarshalled = true

    return json.Marshal(&struct {
        
        PositiveSentimentReasons []Sentimentinsightentry `json:"positiveSentimentReasons"`
        
        NegativeSentimentReasons []Sentimentinsightentry `json:"negativeSentimentReasons"`
        *Alias
    }{

        
        PositiveSentimentReasons: []Sentimentinsightentry{{}},
        


        
        NegativeSentimentReasons: []Sentimentinsightentry{{}},
        

        Alias: (*Alias)(u),
    })
}

