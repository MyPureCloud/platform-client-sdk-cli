package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    MessengerpositionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type MessengerpositionsettingsDud struct { 
    


    


    

}

// Messengerpositionsettings - Position settings for messenger
type Messengerpositionsettings struct { 
    // Alignment - The alignment for messenger position
    Alignment string `json:"alignment"`


    // SideSpace - The sidespace value for messenger position
    SideSpace int `json:"sideSpace"`


    // BottomSpace - The bottomspace value for messenger position
    BottomSpace int `json:"bottomSpace"`

}

// String returns a JSON representation of the model
func (o *Messengerpositionsettings) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Messengerpositionsettings) MarshalJSON() ([]byte, error) {
    type Alias Messengerpositionsettings

    if MessengerpositionsettingsMarshalled {
        return []byte("{}"), nil
    }
    MessengerpositionsettingsMarshalled = true

    return json.Marshal(&struct { 
        Alignment string `json:"alignment"`
        
        SideSpace int `json:"sideSpace"`
        
        BottomSpace int `json:"bottomSpace"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

