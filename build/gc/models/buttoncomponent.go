package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ButtoncomponentMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ButtoncomponentDud struct { 
    


    

}

// Buttoncomponent - Structured template button object.
type Buttoncomponent struct { 
    // Title - Text to show inside the button.
    Title string `json:"title"`


    // Actions - The button actions (Deprecated).
    Actions Contentactions `json:"actions"`

}

// String returns a JSON representation of the model
func (o *Buttoncomponent) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buttoncomponent) MarshalJSON() ([]byte, error) {
    type Alias Buttoncomponent

    if ButtoncomponentMarshalled {
        return []byte("{}"), nil
    }
    ButtoncomponentMarshalled = true

    return json.Marshal(&struct {
        
        Title string `json:"title"`
        
        Actions Contentactions `json:"actions"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

