package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailconfigMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailconfigDud struct { 
    


    


    


    

}

// Emailconfig
type Emailconfig struct { 
    // EmailColumns - The contact list columns specifying the email address(es) of the contact.
    EmailColumns []string `json:"emailColumns"`


    // ContentTemplate - The content template used to formulate the email to send to the contact.
    ContentTemplate Domainentityref `json:"contentTemplate"`


    // FromAddress - The email address that will be used as the sender of the email.
    FromAddress Fromemailaddress `json:"fromAddress"`


    // ReplyToAddress - The email address from which any reply will be sent.
    ReplyToAddress Replytoemailaddress `json:"replyToAddress"`

}

// String returns a JSON representation of the model
func (o *Emailconfig) String() string {
     o.EmailColumns = []string{""} 
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailconfig) MarshalJSON() ([]byte, error) {
    type Alias Emailconfig

    if EmailconfigMarshalled {
        return []byte("{}"), nil
    }
    EmailconfigMarshalled = true

    return json.Marshal(&struct {
        
        EmailColumns []string `json:"emailColumns"`
        
        ContentTemplate Domainentityref `json:"contentTemplate"`
        
        FromAddress Fromemailaddress `json:"fromAddress"`
        
        ReplyToAddress Replytoemailaddress `json:"replyToAddress"`
        *Alias
    }{

        
        EmailColumns: []string{""},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

