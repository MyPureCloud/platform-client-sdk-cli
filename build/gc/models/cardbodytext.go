package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CardbodytextMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CardbodytextDud struct { 
    


    

}

// Cardbodytext
type Cardbodytext struct { 
    // Content - Body content for carousel card.
    Content string `json:"content"`


    // ContentType - Body content type for carousel card. Allowed value: text/plain
    ContentType string `json:"contentType"`

}

// String returns a JSON representation of the model
func (o *Cardbodytext) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Cardbodytext) MarshalJSON() ([]byte, error) {
    type Alias Cardbodytext

    if CardbodytextMarshalled {
        return []byte("{}"), nil
    }
    CardbodytextMarshalled = true

    return json.Marshal(&struct {
        
        Content string `json:"content"`
        
        ContentType string `json:"contentType"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

