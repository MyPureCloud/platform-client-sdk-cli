package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OndemandsummaryacceptedresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OndemandsummaryacceptedresponseDud struct { 
    


    


    SelfUri string `json:"selfUri"`

}

// Ondemandsummaryacceptedresponse - Acknowledgement of an on-demand summary request.
type Ondemandsummaryacceptedresponse struct { 
    // Id - The id of the summary.
    Id string `json:"id"`


    // Status - Status of the summary request.
    Status string `json:"status"`


    

}

// String returns a JSON representation of the model
func (o *Ondemandsummaryacceptedresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ondemandsummaryacceptedresponse) MarshalJSON() ([]byte, error) {
    type Alias Ondemandsummaryacceptedresponse

    if OndemandsummaryacceptedresponseMarshalled {
        return []byte("{}"), nil
    }
    OndemandsummaryacceptedresponseMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Status string `json:"status"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

