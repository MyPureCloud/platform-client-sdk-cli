package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    KnowledgeconversationturnMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type KnowledgeconversationturnDud struct { 
    


    

}

// Knowledgeconversationturn
type Knowledgeconversationturn struct { 
    // Participant - The participant type.
    Participant string `json:"participant"`


    // Text - The message text.
    Text string `json:"text"`

}

// String returns a JSON representation of the model
func (o *Knowledgeconversationturn) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Knowledgeconversationturn) MarshalJSON() ([]byte, error) {
    type Alias Knowledgeconversationturn

    if KnowledgeconversationturnMarshalled {
        return []byte("{}"), nil
    }
    KnowledgeconversationturnMarshalled = true

    return json.Marshal(&struct {
        
        Participant string `json:"participant"`
        
        Text string `json:"text"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

