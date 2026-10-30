package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecipientwarningMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecipientwarningDud struct { 
    Email string `json:"email"`


    Reason string `json:"reason"`


    FieldName string `json:"fieldName"`

}

// Recipientwarning
type Recipientwarning struct { 
    


    


    

}

// String returns a JSON representation of the model
func (o *Recipientwarning) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recipientwarning) MarshalJSON() ([]byte, error) {
    type Alias Recipientwarning

    if RecipientwarningMarshalled {
        return []byte("{}"), nil
    }
    RecipientwarningMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

