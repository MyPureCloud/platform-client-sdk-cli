package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FormMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FormDud struct { 
    


    


    


    


    


    

}

// Form - Form configuration for response management
type Form struct { 
    // FormDescription - Description of the form
    FormDescription string `json:"formDescription"`


    // ReceivedMessage - Message displayed when response is received
    ReceivedMessage Formmessage `json:"receivedMessage"`


    // ReplyMessage - Message displayed as reply
    ReplyMessage Formmessage `json:"replyMessage"`


    // Introduction - Introduction section of the form
    Introduction Formintroduction `json:"introduction"`


    // FormPages - Pages of the form
    FormPages []Formpage `json:"formPages"`


    // ShowSummary - Whether to show a summary after form completion
    ShowSummary bool `json:"showSummary"`

}

// String returns a JSON representation of the model
func (o *Form) String() string {
    
    
    
    
     o.FormPages = []Formpage{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Form) MarshalJSON() ([]byte, error) {
    type Alias Form

    if FormMarshalled {
        return []byte("{}"), nil
    }
    FormMarshalled = true

    return json.Marshal(&struct {
        
        FormDescription string `json:"formDescription"`
        
        ReceivedMessage Formmessage `json:"receivedMessage"`
        
        ReplyMessage Formmessage `json:"replyMessage"`
        
        Introduction Formintroduction `json:"introduction"`
        
        FormPages []Formpage `json:"formPages"`
        
        ShowSummary bool `json:"showSummary"`
        *Alias
    }{

        


        


        


        


        
        FormPages: []Formpage{{}},
        


        

        Alias: (*Alias)(u),
    })
}

