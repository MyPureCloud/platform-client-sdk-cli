package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CaseexternalcontactreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CaseexternalcontactreferenceDud struct { 
    


    SelfUri string `json:"selfUri"`

}

// Caseexternalcontactreference
type Caseexternalcontactreference struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    

}

// String returns a JSON representation of the model
func (o *Caseexternalcontactreference) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Caseexternalcontactreference) MarshalJSON() ([]byte, error) {
    type Alias Caseexternalcontactreference

    if CaseexternalcontactreferenceMarshalled {
        return []byte("{}"), nil
    }
    CaseexternalcontactreferenceMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

