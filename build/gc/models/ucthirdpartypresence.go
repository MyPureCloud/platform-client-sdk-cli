package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    UcthirdpartypresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type UcthirdpartypresenceDud struct { 
    


    


    


    

}

// Ucthirdpartypresence - Update a Genesys Cloud user's presence from a given 3rd-party integration
type Ucthirdpartypresence struct { 
    // Email - Primary Email address of the associated Genesys Cloud user.
    Email string `json:"email"`


    // Presence - Integration presence value.
    Presence string `json:"presence"`


    // Message - Integration presence message.
    Message string `json:"message"`


    // DateModified - ISO 8601 timestamp of presence value change.
    DateModified time.Time `json:"dateModified"`

}

// String returns a JSON representation of the model
func (o *Ucthirdpartypresence) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Ucthirdpartypresence) MarshalJSON() ([]byte, error) {
    type Alias Ucthirdpartypresence

    if UcthirdpartypresenceMarshalled {
        return []byte("{}"), nil
    }
    UcthirdpartypresenceMarshalled = true

    return json.Marshal(&struct {
        
        Email string `json:"email"`
        
        Presence string `json:"presence"`
        
        Message string `json:"message"`
        
        DateModified time.Time `json:"dateModified"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

