package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    YearmonthMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type YearmonthDud struct { 
    


    


    


    

}

// Yearmonth
type Yearmonth struct { 
    // Year
    Year int `json:"year"`


    // Month
    Month string `json:"month"`


    // MonthValue
    MonthValue int `json:"monthValue"`


    // LeapYear
    LeapYear bool `json:"leapYear"`

}

// String returns a JSON representation of the model
func (o *Yearmonth) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Yearmonth) MarshalJSON() ([]byte, error) {
    type Alias Yearmonth

    if YearmonthMarshalled {
        return []byte("{}"), nil
    }
    YearmonthMarshalled = true

    return json.Marshal(&struct {
        
        Year int `json:"year"`
        
        Month string `json:"month"`
        
        MonthValue int `json:"monthValue"`
        
        LeapYear bool `json:"leapYear"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

