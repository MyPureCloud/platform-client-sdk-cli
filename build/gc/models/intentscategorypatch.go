package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    IntentscategorypatchMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type IntentscategorypatchDud struct { 
    Id string `json:"id"`


    


    


    SelfUri string `json:"selfUri"`

}

// Intentscategorypatch
type Intentscategorypatch struct { 
    


    // Name
    Name string `json:"name"`


    // Description - Description of the category
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Intentscategorypatch) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Intentscategorypatch) MarshalJSON() ([]byte, error) {
    type Alias Intentscategorypatch

    if IntentscategorypatchMarshalled {
        return []byte("{}"), nil
    }
    IntentscategorypatchMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

