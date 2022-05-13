package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TextbottranscriptMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TextbottranscriptDud struct { 
    


    

}

// Textbottranscript - Data for a single bot flow transcript.
type Textbottranscript struct { 
    // Text - The text of the transcript item.
    Text string `json:"text"`


    // Confidence - The confidence factor, expressed as a decimal between 0.0 and 1.0, of the transcript item.
    Confidence float32 `json:"confidence"`

}

// String returns a JSON representation of the model
func (o *Textbottranscript) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Textbottranscript) MarshalJSON() ([]byte, error) {
    type Alias Textbottranscript

    if TextbottranscriptMarshalled {
        return []byte("{}"), nil
    }
    TextbottranscriptMarshalled = true

    return json.Marshal(&struct {
        
        Text string `json:"text"`
        
        Confidence float32 `json:"confidence"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

