package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PositionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PositionsettingsDud struct { 
    


    


    

}

// Positionsettings - Settings concerning position
type Positionsettings struct { 
    // Alignment - The alignment for position
    Alignment string `json:"alignment"`


    // SideSpace - The sidespace value for position
    SideSpace int `json:"sideSpace"`


    // BottomSpace - The bottomspace value for position
    BottomSpace int `json:"bottomSpace"`

}

// String returns a JSON representation of the model
func (o *Positionsettings) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Positionsettings) MarshalJSON() ([]byte, error) {
    type Alias Positionsettings

    if PositionsettingsMarshalled {
        return []byte("{}"), nil
    }
    PositionsettingsMarshalled = true

    return json.Marshal(&struct { 
        Alignment string `json:"alignment"`
        
        SideSpace int `json:"sideSpace"`
        
        BottomSpace int `json:"bottomSpace"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

