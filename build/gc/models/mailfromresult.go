package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MailfromresultMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MailfromresultDud struct { 
    


    


    

}

// Mailfromresult
type Mailfromresult struct { 
    // Status - The verification status.
    Status string `json:"status"`


    // Records - The list of DNS records that pertain that need to exist for verification.
    Records []Record `json:"records"`


    // MailFromDomain - The custom MAIL FROM domain.
    MailFromDomain string `json:"mailFromDomain"`

}

// String returns a JSON representation of the model
func (o *Mailfromresult) String() string {
    
     o.Records = []Record{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mailfromresult) MarshalJSON() ([]byte, error) {
    type Alias Mailfromresult

    if MailfromresultMarshalled {
        return []byte("{}"), nil
    }
    MailfromresultMarshalled = true

    return json.Marshal(&struct {
        
        Status string `json:"status"`
        
        Records []Record `json:"records"`
        
        MailFromDomain string `json:"mailFromDomain"`
        *Alias
    }{

        


        
        Records: []Record{{}},
        


        

        Alias: (*Alias)(u),
    })
}

