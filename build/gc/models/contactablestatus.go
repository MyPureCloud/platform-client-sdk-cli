package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ContactablestatusMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ContactablestatusDud struct { 
    


    

}

// Contactablestatus
type Contactablestatus struct { 
    // Contactable - Indicates whether or not the entire contact is contactable for the associated media type.
    Contactable bool `json:"contactable"`


    // ColumnStatus - A map of individual contact method columns to whether the individual column is contactable for the associated media type.
    ColumnStatus map[string]Columnstatus `json:"columnStatus"`

}

// String returns a JSON representation of the model
func (o *Contactablestatus) String() string {
    
    
    
    
    
    
     o.ColumnStatus = map[string]Columnstatus{"": {}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Contactablestatus) MarshalJSON() ([]byte, error) {
    type Alias Contactablestatus

    if ContactablestatusMarshalled {
        return []byte("{}"), nil
    }
    ContactablestatusMarshalled = true

    return json.Marshal(&struct { 
        Contactable bool `json:"contactable"`
        
        ColumnStatus map[string]Columnstatus `json:"columnStatus"`
        
        *Alias
    }{
        

        

        

        
        ColumnStatus: map[string]Columnstatus{"": {}},
        

        
        Alias: (*Alias)(u),
    })
}

