package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TransferresponsemodifiedbyMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TransferresponsemodifiedbyDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`

}

// Transferresponsemodifiedby
type Transferresponsemodifiedby struct { 
    


    

}

// String returns a JSON representation of the model
func (o *Transferresponsemodifiedby) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Transferresponsemodifiedby) MarshalJSON() ([]byte, error) {
    type Alias Transferresponsemodifiedby

    if TransferresponsemodifiedbyMarshalled {
        return []byte("{}"), nil
    }
    TransferresponsemodifiedbyMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

