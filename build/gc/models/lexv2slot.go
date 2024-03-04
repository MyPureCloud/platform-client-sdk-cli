package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    Lexv2slotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type Lexv2slotDud struct { 
    


    


    


    


    

}

// Lexv2slot
type Lexv2slot struct { 
    // SlotName - The slot name
    SlotName string `json:"slotName"`


    // Description - The slot description
    Description string `json:"description"`


    // SlotId - The slot id
    SlotId string `json:"slotId"`


    // VarType - The slot type
    VarType string `json:"type"`


    // SlotTypeId - The slot type id
    SlotTypeId string `json:"slotTypeId"`

}

// String returns a JSON representation of the model
func (o *Lexv2slot) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Lexv2slot) MarshalJSON() ([]byte, error) {
    type Alias Lexv2slot

    if Lexv2slotMarshalled {
        return []byte("{}"), nil
    }
    Lexv2slotMarshalled = true

    return json.Marshal(&struct {
        
        SlotName string `json:"slotName"`
        
        Description string `json:"description"`
        
        SlotId string `json:"slotId"`
        
        VarType string `json:"type"`
        
        SlotTypeId string `json:"slotTypeId"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

