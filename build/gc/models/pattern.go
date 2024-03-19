package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    PatternMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type PatternDud struct { 
    


    


    

}

// Pattern
type Pattern struct { 
    // VarType - Pattern type (Daily/Weekly)
    VarType string `json:"type"`


    // Interval - The interval of days between the occurrences for Daily pattern type, and weeks between the occurrences for Weekly
    Interval int `json:"interval"`


    // DaysOfWeek - The day(s) of week the occurrence should be repeated. Required to set for Weekly pattern type. E.g. [\"Monday\", \"Wednesday\"]
    DaysOfWeek []string `json:"daysOfWeek"`

}

// String returns a JSON representation of the model
func (o *Pattern) String() string {
    
    
     o.DaysOfWeek = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Pattern) MarshalJSON() ([]byte, error) {
    type Alias Pattern

    if PatternMarshalled {
        return []byte("{}"), nil
    }
    PatternMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Interval int `json:"interval"`
        
        DaysOfWeek []string `json:"daysOfWeek"`
        *Alias
    }{

        


        


        
        DaysOfWeek: []string{""},
        

        Alias: (*Alias)(u),
    })
}

