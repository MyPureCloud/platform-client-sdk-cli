package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    HistoryheaderstranslationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type HistoryheaderstranslationDud struct { 
    


    


    


    


    


    


    


    

}

// Historyheaderstranslation
type Historyheaderstranslation struct { 
    // From - A translation for the word \"from\", for the expected language
    From string `json:"from"`


    // To - A translation for the word \"to\", for the expected language
    To string `json:"to"`


    // Cc - A translation for the word \"cc\", for the expected language
    Cc string `json:"cc"`


    // Subject - A translation for the word \"subject\", for the expected language
    Subject string `json:"subject"`


    // ReplyPrefix - A translation for the subject prefix \"Reply\", for the expected language
    ReplyPrefix string `json:"replyPrefix"`


    // ForwardPrefix - A translation for the subject prefix \"Forward\", for the expected language
    ForwardPrefix string `json:"forwardPrefix"`


    // Sent - A translation for the word \"sent\", for the expected language
    Sent string `json:"sent"`


    // Language - The code of the expected language
    Language string `json:"language"`

}

// String returns a JSON representation of the model
func (o *Historyheaderstranslation) String() string {
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Historyheaderstranslation) MarshalJSON() ([]byte, error) {
    type Alias Historyheaderstranslation

    if HistoryheaderstranslationMarshalled {
        return []byte("{}"), nil
    }
    HistoryheaderstranslationMarshalled = true

    return json.Marshal(&struct {
        
        From string `json:"from"`
        
        To string `json:"to"`
        
        Cc string `json:"cc"`
        
        Subject string `json:"subject"`
        
        ReplyPrefix string `json:"replyPrefix"`
        
        ForwardPrefix string `json:"forwardPrefix"`
        
        Sent string `json:"sent"`
        
        Language string `json:"language"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

