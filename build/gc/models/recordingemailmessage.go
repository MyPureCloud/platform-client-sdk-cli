package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordingemailmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordingemailmessageDud struct { 
    


    


    


    


    


    


    


    


    


    


    

}

// Recordingemailmessage
type Recordingemailmessage struct { 
    // HtmlBody
    HtmlBody string `json:"htmlBody"`


    // TextBody
    TextBody string `json:"textBody"`


    // Id
    Id string `json:"id"`


    // To
    To []Emailaddress `json:"to"`


    // Cc
    Cc []Emailaddress `json:"cc"`


    // Bcc
    Bcc []Emailaddress `json:"bcc"`


    // From
    From Emailaddress `json:"from"`


    // ReplyTo - Indicates the address to which the author of the message suggests that replies be sent
    ReplyTo Emailaddress `json:"replyTo"`


    // Subject
    Subject string `json:"subject"`


    // Attachments
    Attachments []Emailattachment `json:"attachments"`


    // Time - Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`

}

// String returns a JSON representation of the model
func (o *Recordingemailmessage) String() string {
    
    
    
     o.To = []Emailaddress{{}} 
     o.Cc = []Emailaddress{{}} 
     o.Bcc = []Emailaddress{{}} 
    
    
    
     o.Attachments = []Emailattachment{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordingemailmessage) MarshalJSON() ([]byte, error) {
    type Alias Recordingemailmessage

    if RecordingemailmessageMarshalled {
        return []byte("{}"), nil
    }
    RecordingemailmessageMarshalled = true

    return json.Marshal(&struct {
        
        HtmlBody string `json:"htmlBody"`
        
        TextBody string `json:"textBody"`
        
        Id string `json:"id"`
        
        To []Emailaddress `json:"to"`
        
        Cc []Emailaddress `json:"cc"`
        
        Bcc []Emailaddress `json:"bcc"`
        
        From Emailaddress `json:"from"`
        
        ReplyTo Emailaddress `json:"replyTo"`
        
        Subject string `json:"subject"`
        
        Attachments []Emailattachment `json:"attachments"`
        
        Time time.Time `json:"time"`
        *Alias
    }{

        


        


        


        
        To: []Emailaddress{{}},
        


        
        Cc: []Emailaddress{{}},
        


        
        Bcc: []Emailaddress{{}},
        


        


        


        


        
        Attachments: []Emailattachment{{}},
        


        

        Alias: (*Alias)(u),
    })
}

