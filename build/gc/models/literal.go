package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    LiteralMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type LiteralDud struct { 
    


    


    


    


    


    


    

}

// Literal
type Literal struct { 
    // VarString - A string value
    VarString string `json:"string"`


    // Integer - A positive or negative whole number, including zero
    Integer int `json:"integer"`


    // Number - A positive or negative decimal number, including zero
    Number float64 `json:"number"`


    // Date - A date value, must be in the format of yyyy-MM-dd, e.g. 2024-09-23. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
    Date time.Time `json:"date"`


    // Datetime - A date time value, must be in the format of yyyy-MM-dd'T'HH:mm:ss.SSSZ, e.g. 2024-10-02T01:01:01.111Z. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    Datetime time.Time `json:"datetime"`


    // Special - A special value enum, such as Wildcard, Null, etc
    Special string `json:"special"`


    // Boolean - A boolean value
    Boolean bool `json:"boolean"`

}

// String returns a JSON representation of the model
func (o *Literal) String() string {
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Literal) MarshalJSON() ([]byte, error) {
    type Alias Literal

    if LiteralMarshalled {
        return []byte("{}"), nil
    }
    LiteralMarshalled = true

    return json.Marshal(&struct {
        
        VarString string `json:"string"`
        
        Integer int `json:"integer"`
        
        Number float64 `json:"number"`
        
        Date time.Time `json:"date"`
        
        Datetime time.Time `json:"datetime"`
        
        Special string `json:"special"`
        
        Boolean bool `json:"boolean"`
        *Alias
    }{

        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

