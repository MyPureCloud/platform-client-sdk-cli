package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatchdraftrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatchdraftrequestDud struct { 
    


    


    


    


    


    


    

}

// Patchdraftrequest
type Patchdraftrequest struct { 
    // Subject - The subject of the email draft. Replaces the existing subject.
    Subject string `json:"subject"`


    // To - The To recipients of the email draft. Replaces the entire To list.
    To []Emailaddress `json:"to"`


    // Cc - The CC recipients of the email draft. Replaces the entire CC list.
    Cc []Emailaddress `json:"cc"`


    // Bcc - The BCC recipients of the email draft. Replaces the entire BCC list.
    Bcc []Emailaddress `json:"bcc"`


    // OverrideHeaders - Custom headers to set on the outbound email. Replaces the entire override headers list. Header keys are case-sensitive: keys differing only in case (e.g. MSIP and msip) are kept as distinct headers, while exact duplicate keys result in a 400 Bad Request.
    OverrideHeaders []Headerentry `json:"overrideHeaders"`


    // HistoryIncluded - Whether to include message history in the draft. Once set to true, cannot be reverted to false.
    HistoryIncluded bool `json:"historyIncluded"`


    // Translation - A set of definitions to translate email attributes and correctly display date and time, for a given language. Used when historyIncluded is set to true.
    Translation Historyheaderstranslation `json:"translation"`

}

// String returns a JSON representation of the model
func (o *Patchdraftrequest) String() string {
    
     o.To = []Emailaddress{{}} 
     o.Cc = []Emailaddress{{}} 
     o.Bcc = []Emailaddress{{}} 
     o.OverrideHeaders = []Headerentry{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Patchdraftrequest) MarshalJSON() ([]byte, error) {
    type Alias Patchdraftrequest

    if PatchdraftrequestMarshalled {
        return []byte("{}"), nil
    }
    PatchdraftrequestMarshalled = true

    return json.Marshal(&struct {
        
        Subject string `json:"subject"`
        
        To []Emailaddress `json:"to"`
        
        Cc []Emailaddress `json:"cc"`
        
        Bcc []Emailaddress `json:"bcc"`
        
        OverrideHeaders []Headerentry `json:"overrideHeaders"`
        
        HistoryIncluded bool `json:"historyIncluded"`
        
        Translation Historyheaderstranslation `json:"translation"`
        *Alias
    }{

        


        
        To: []Emailaddress{{}},
        


        
        Cc: []Emailaddress{{}},
        


        
        Bcc: []Emailaddress{{}},
        


        
        OverrideHeaders: []Headerentry{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

