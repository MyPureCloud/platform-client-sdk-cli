package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    GenaiphrasesjobrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type GenaiphrasesjobrequestDud struct { 
    

}

// Genaiphrasesjobrequest
type Genaiphrasesjobrequest struct { 
    // Topic - topic used for phrases generation by GenAI
    Topic Genaiphrasesjobtopic `json:"topic"`

}

// String returns a JSON representation of the model
func (o *Genaiphrasesjobrequest) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Genaiphrasesjobrequest) MarshalJSON() ([]byte, error) {
    type Alias Genaiphrasesjobrequest

    if GenaiphrasesjobrequestMarshalled {
        return []byte("{}"), nil
    }
    GenaiphrasesjobrequestMarshalled = true

    return json.Marshal(&struct {
        
        Topic Genaiphrasesjobtopic `json:"topic"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

