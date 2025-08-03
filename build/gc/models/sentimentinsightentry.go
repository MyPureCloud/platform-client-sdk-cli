package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SentimentinsightentryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SentimentinsightentryDud struct { 
    


    

}

// Sentimentinsightentry
type Sentimentinsightentry struct { 
    // Title - The title given to the insight
    Title string `json:"title"`


    // Description - A short description of the insight
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Sentimentinsightentry) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sentimentinsightentry) MarshalJSON() ([]byte, error) {
    type Alias Sentimentinsightentry

    if SentimentinsightentryMarshalled {
        return []byte("{}"), nil
    }
    SentimentinsightentryMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

