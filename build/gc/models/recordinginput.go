package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordinginputMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordinginputDud struct { 
    


    


    


    

}

// Recordinginput
type Recordinginput struct { 
    // Id - Unique identifier for the input.
    Id string `json:"id"`


    // Title - The main text displayed for the input field(s).
    Title string `json:"title"`


    // Subtitle - Additional text providing more details about the input field(s).
    Subtitle string `json:"subtitle"`


    // ResponseText - Text response from end customer.
    ResponseText string `json:"responseText"`

}

// String returns a JSON representation of the model
func (o *Recordinginput) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recordinginput) MarshalJSON() ([]byte, error) {
    type Alias Recordinginput

    if RecordinginputMarshalled {
        return []byte("{}"), nil
    }
    RecordinginputMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Title string `json:"title"`
        
        Subtitle string `json:"subtitle"`
        
        ResponseText string `json:"responseText"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

