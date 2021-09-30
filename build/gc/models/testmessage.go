package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TestmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TestmessageDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    

}

// Testmessage
type Testmessage struct { 
    


    // To - The recipients of the email message.
    To []Emailaddress `json:"to"`


    // From - The sender of the email message.
    From Emailaddress `json:"from"`


    // Subject - The subject of the email message.
    Subject string `json:"subject"`


    // TextBody - The text body of the email message.
    TextBody string `json:"textBody"`


    // HtmlBody - The html body of the email message
    HtmlBody string `json:"htmlBody"`


    // Time - The time when the message was sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`

}

// String returns a JSON representation of the model
func (o *Testmessage) String() string {
    
    
    
    
     o.To = []Emailaddress{{}} 
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Testmessage) MarshalJSON() ([]byte, error) {
    type Alias Testmessage

    if TestmessageMarshalled {
        return []byte("{}"), nil
    }
    TestmessageMarshalled = true

    return json.Marshal(&struct { 
        
        
        To []Emailaddress `json:"to"`
        
        From Emailaddress `json:"from"`
        
        Subject string `json:"subject"`
        
        TextBody string `json:"textBody"`
        
        HtmlBody string `json:"htmlBody"`
        
        Time time.Time `json:"time"`
        
        *Alias
    }{
        

        

        

        
        To: []Emailaddress{{}},
        

        

        

        

        

        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

