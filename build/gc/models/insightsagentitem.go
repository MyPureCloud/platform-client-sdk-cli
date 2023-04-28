package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    InsightsagentitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type InsightsagentitemDud struct { 
    Id string `json:"id"`


    


    SelfUri string `json:"selfUri"`

}

// Insightsagentitem
type Insightsagentitem struct { 
    


    // Name
    Name string `json:"name"`


    

}

// String returns a JSON representation of the model
func (o *Insightsagentitem) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Insightsagentitem) MarshalJSON() ([]byte, error) {
    type Alias Insightsagentitem

    if InsightsagentitemMarshalled {
        return []byte("{}"), nil
    }
    InsightsagentitemMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

