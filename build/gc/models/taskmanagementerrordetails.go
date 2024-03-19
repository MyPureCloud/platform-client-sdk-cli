package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TaskmanagementerrordetailsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TaskmanagementerrordetailsDud struct { 
    


    

}

// Taskmanagementerrordetails
type Taskmanagementerrordetails struct { 
    // Code - System defined error code for the error.
    Code string `json:"code"`


    // Message - Error message.
    Message string `json:"message"`

}

// String returns a JSON representation of the model
func (o *Taskmanagementerrordetails) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Taskmanagementerrordetails) MarshalJSON() ([]byte, error) {
    type Alias Taskmanagementerrordetails

    if TaskmanagementerrordetailsMarshalled {
        return []byte("{}"), nil
    }
    TaskmanagementerrordetailsMarshalled = true

    return json.Marshal(&struct {
        
        Code string `json:"code"`
        
        Message string `json:"message"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

