package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ThirdpartysuggestionsourceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ThirdpartysuggestionsourceDud struct { 
    


    

}

// Thirdpartysuggestionsource
type Thirdpartysuggestionsource struct { 
    // Name - The third party source name.
    Name string `json:"name"`


    // Url - The link to the source document or page.
    Url string `json:"url"`

}

// String returns a JSON representation of the model
func (o *Thirdpartysuggestionsource) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Thirdpartysuggestionsource) MarshalJSON() ([]byte, error) {
    type Alias Thirdpartysuggestionsource

    if ThirdpartysuggestionsourceMarshalled {
        return []byte("{}"), nil
    }
    ThirdpartysuggestionsourceMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Url string `json:"url"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

