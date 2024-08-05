package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BulistalternativeshifttradesresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BulistalternativeshifttradesresponseDud struct { 
    


    

}

// Bulistalternativeshifttradesresponse
type Bulistalternativeshifttradesresponse struct { 
    // Job - The asynchronous job handling the request. Null if result returns synchronously
    Job Bualternativeshiftjobresponse `json:"job"`


    // Result - The result of the request. May come via notification. Null if job is populated
    Result Alternativeshifttradelisting `json:"result"`

}

// String returns a JSON representation of the model
func (o *Bulistalternativeshifttradesresponse) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Bulistalternativeshifttradesresponse) MarshalJSON() ([]byte, error) {
    type Alias Bulistalternativeshifttradesresponse

    if BulistalternativeshifttradesresponseMarshalled {
        return []byte("{}"), nil
    }
    BulistalternativeshifttradesresponseMarshalled = true

    return json.Marshal(&struct {
        
        Job Bualternativeshiftjobresponse `json:"job"`
        
        Result Alternativeshifttradelisting `json:"result"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

