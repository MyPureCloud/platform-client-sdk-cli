package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailmessageMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailmessageDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    


    


    


    


    


    


    EmailSizeBytes int `json:"emailSizeBytes"`


    MaxEmailSizeBytes int `json:"maxEmailSizeBytes"`


    SelfUri string `json:"selfUri"`

}

// Emailmessage
type Emailmessage struct { 
    


    // Name
    Name string `json:"name"`


    // To - The recipients of the email message.
    To []Emailaddress `json:"to"`


    // Cc - The recipients that were copied on the email message.
    Cc []Emailaddress `json:"cc"`


    // Bcc - The recipients that were blind copied on the email message.
    Bcc []Emailaddress `json:"bcc"`


    // From - The sender of the email message.
    From Emailaddress `json:"from"`


    // ReplyTo - The receiver of the reply email message.
    ReplyTo Emailaddress `json:"replyTo"`


    // Subject - The subject of the email message.
    Subject string `json:"subject"`


    // Attachments - The attachments of the email message.
    Attachments []Attachment `json:"attachments"`


    // TextBody - The text body of the email message.
    TextBody string `json:"textBody"`


    // HtmlBody - The html body of the email message.
    HtmlBody string `json:"htmlBody"`


    // Time - The time when the message was received or sent. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Time time.Time `json:"time"`


    // HistoryIncluded - Indicates whether the history of previous emails of the conversation is included within the email bodies of this message.
    HistoryIncluded bool `json:"historyIncluded"`


    


    


    

}

// String returns a JSON representation of the model
func (o *Emailmessage) String() string {
    
     o.To = []Emailaddress{{}} 
     o.Cc = []Emailaddress{{}} 
     o.Bcc = []Emailaddress{{}} 
    
    
    
     o.Attachments = []Attachment{{}} 
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailmessage) MarshalJSON() ([]byte, error) {
    type Alias Emailmessage

    if EmailmessageMarshalled {
        return []byte("{}"), nil
    }
    EmailmessageMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        To []Emailaddress `json:"to"`
        
        Cc []Emailaddress `json:"cc"`
        
        Bcc []Emailaddress `json:"bcc"`
        
        From Emailaddress `json:"from"`
        
        ReplyTo Emailaddress `json:"replyTo"`
        
        Subject string `json:"subject"`
        
        Attachments []Attachment `json:"attachments"`
        
        TextBody string `json:"textBody"`
        
        HtmlBody string `json:"htmlBody"`
        
        Time time.Time `json:"time"`
        
        HistoryIncluded bool `json:"historyIncluded"`
        *Alias
    }{

        


        


        
        To: []Emailaddress{{}},
        


        
        Cc: []Emailaddress{{}},
        


        
        Bcc: []Emailaddress{{}},
        


        


        


        


        
        Attachments: []Attachment{{}},
        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

