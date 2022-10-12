package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DocumentqueryintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DocumentqueryintervalDud struct { 
    


    

}

// Documentqueryinterval
type Documentqueryinterval struct { 
    // Field - Specifies the date field to be used for date and time range.
    Field string `json:"field"`


    // Value - Specifies the date and time range for filtering the documents. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Documentqueryinterval) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Documentqueryinterval) MarshalJSON() ([]byte, error) {
    type Alias Documentqueryinterval

    if DocumentqueryintervalMarshalled {
        return []byte("{}"), nil
    }
    DocumentqueryintervalMarshalled = true

    return json.Marshal(&struct {
        
        Field string `json:"field"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

