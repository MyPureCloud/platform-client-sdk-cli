package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ShorttermforecastreferenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ShorttermforecastreferenceDud struct { 
    Id string `json:"id"`


    SelfUri string `json:"selfUri"`


    


    

}

// Shorttermforecastreference - A pointer to a short term forecast
type Shorttermforecastreference struct { 
    


    


    // WeekDate - The weekDate of the short term forecast in yyyy-MM-dd format
    WeekDate string `json:"weekDate"`


    // Description - The description of the short term forecast
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Shorttermforecastreference) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Shorttermforecastreference) MarshalJSON() ([]byte, error) {
    type Alias Shorttermforecastreference

    if ShorttermforecastreferenceMarshalled {
        return []byte("{}"), nil
    }
    ShorttermforecastreferenceMarshalled = true

    return json.Marshal(&struct {
        
        WeekDate string `json:"weekDate"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

