package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatekpirequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatekpirequestDud struct { 
    


    


    


    


    

}

// Createkpirequest
type Createkpirequest struct { 
    // Name - The name of the Key Performance Indicator.
    Name string `json:"name"`


    // Description - The description of the Key Performance Indicator.
    Description string `json:"description"`


    // KpiType - The type of the Key Performance Indicator.
    KpiType string `json:"kpiType"`


    // WrapUpCodeConfig - Defines what wrap up codes are mapped to Key Performance Indicator.
    WrapUpCodeConfig Wrapupcodeconfig `json:"wrapUpCodeConfig"`


    // Source - The source of the Key Performance Indicator.
    Source string `json:"source"`

}

// String returns a JSON representation of the model
func (o *Createkpirequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createkpirequest) MarshalJSON() ([]byte, error) {
    type Alias Createkpirequest

    if CreatekpirequestMarshalled {
        return []byte("{}"), nil
    }
    CreatekpirequestMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Description string `json:"description"`
        
        KpiType string `json:"kpiType"`
        
        WrapUpCodeConfig Wrapupcodeconfig `json:"wrapUpCodeConfig"`
        
        Source string `json:"source"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

