package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseexternalidupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseexternalidupdateDud struct { 
    

}

// Caseexternalidupdate
type Caseexternalidupdate struct { 
    // ExternalId - The identifier of the Case in an external system. Minimum length is 1 character. Maximum length of 64 characters.
    ExternalId string `json:"externalId"`

}

// String returns a JSON representation of the model
func (o *Caseexternalidupdate) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseexternalidupdate) MarshalJSON() ([]byte, error) {
    type Alias Caseexternalidupdate

    if CaseexternalidupdateMarshalled {
        return []byte("{}"), nil
    }
    CaseexternalidupdateMarshalled = true

    return json.Marshal(&struct {
        
        ExternalId string `json:"externalId"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

