package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpensocialmediarecipientadditionalidentifierMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpensocialmediarecipientadditionalidentifierDud struct { 
    VarType string `json:"type"`


    Value string `json:"value"`

}

// Opensocialmediarecipientadditionalidentifier
type Opensocialmediarecipientadditionalidentifier struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Opensocialmediarecipientadditionalidentifier) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Opensocialmediarecipientadditionalidentifier) MarshalJSON() ([]byte, error) {
    type Alias Opensocialmediarecipientadditionalidentifier

    if OpensocialmediarecipientadditionalidentifierMarshalled {
        return []byte("{}"), nil
    }
    OpensocialmediarecipientadditionalidentifierMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

