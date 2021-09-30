package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CoachingannotationcreaterequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CoachingannotationcreaterequestDud struct { 
    


    

}

// Coachingannotationcreaterequest
type Coachingannotationcreaterequest struct { 
    // Text - The text of the annotation.
    Text string `json:"text"`


    // AccessType - Determines the permissions required to view this item.
    AccessType string `json:"accessType"`

}

// String returns a JSON representation of the model
func (o *Coachingannotationcreaterequest) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Coachingannotationcreaterequest) MarshalJSON() ([]byte, error) {
    type Alias Coachingannotationcreaterequest

    if CoachingannotationcreaterequestMarshalled {
        return []byte("{}"), nil
    }
    CoachingannotationcreaterequestMarshalled = true

    return json.Marshal(&struct { 
        Text string `json:"text"`
        
        AccessType string `json:"accessType"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

