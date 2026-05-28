package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EmailsummarysettingsentityMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EmailsummarysettingsentityDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Emailsummarysettingsentity
type Emailsummarysettingsentity struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Emailsummarysettingsentity) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Emailsummarysettingsentity) MarshalJSON() ([]byte, error) {
    type Alias Emailsummarysettingsentity

    if EmailsummarysettingsentityMarshalled {
        return []byte("{}"), nil
    }
    EmailsummarysettingsentityMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

