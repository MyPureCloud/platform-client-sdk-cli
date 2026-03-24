package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ActivitycodestyleMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ActivitycodestyleDud struct { 
    


    

}

// Activitycodestyle
type Activitycodestyle struct { 
    // BackgroundColor - Background color for the activity code in hex format starting with # followed by 6 hexadecimal characters (0-9, a-f, A-F)
    BackgroundColor string `json:"backgroundColor"`


    // TextColorTheme - Text color theme for the activity code
    TextColorTheme string `json:"textColorTheme"`

}

// String returns a JSON representation of the model
func (o *Activitycodestyle) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Activitycodestyle) MarshalJSON() ([]byte, error) {
    type Alias Activitycodestyle

    if ActivitycodestyleMarshalled {
        return []byte("{}"), nil
    }
    ActivitycodestyleMarshalled = true

    return json.Marshal(&struct {
        
        BackgroundColor string `json:"backgroundColor"`
        
        TextColorTheme string `json:"textColorTheme"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

