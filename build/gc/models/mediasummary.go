package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MediasummaryMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MediasummaryDud struct { 
    


    

}

// Mediasummary
type Mediasummary struct { 
    // ContactCenter
    ContactCenter Mediasummarydetail `json:"contactCenter"`


    // Enterprise
    Enterprise Mediasummarydetail `json:"enterprise"`

}

// String returns a JSON representation of the model
func (o *Mediasummary) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Mediasummary) MarshalJSON() ([]byte, error) {
    type Alias Mediasummary

    if MediasummaryMarshalled {
        return []byte("{}"), nil
    }
    MediasummaryMarshalled = true

    return json.Marshal(&struct {
        
        ContactCenter Mediasummarydetail `json:"contactCenter"`
        
        Enterprise Mediasummarydetail `json:"enterprise"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

