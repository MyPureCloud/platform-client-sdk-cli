package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcenterstylesettingMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcenterstylesettingDud struct { 
    


    

}

// Supportcenterstylesetting
type Supportcenterstylesetting struct { 
    // HeroStyle - Knowledge portal (previously support center) hero customizations
    HeroStyle Supportcenterherostyle `json:"heroStyle"`


    // GlobalStyle - Knowledge portal (previously support center) global customizations
    GlobalStyle Supportcenterglobalstyle `json:"globalStyle"`

}

// String returns a JSON representation of the model
func (o *Supportcenterstylesetting) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcenterstylesetting) MarshalJSON() ([]byte, error) {
    type Alias Supportcenterstylesetting

    if SupportcenterstylesettingMarshalled {
        return []byte("{}"), nil
    }
    SupportcenterstylesettingMarshalled = true

    return json.Marshal(&struct {
        
        HeroStyle Supportcenterherostyle `json:"heroStyle"`
        
        GlobalStyle Supportcenterglobalstyle `json:"globalStyle"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

