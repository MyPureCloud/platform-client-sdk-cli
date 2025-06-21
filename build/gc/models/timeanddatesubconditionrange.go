package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeanddatesubconditionrangeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeanddatesubconditionrangeDud struct { 
    


    


    

}

// Timeanddatesubconditionrange
type Timeanddatesubconditionrange struct { 
    // Min - The minimum value of the range. Required for the operator BETWEEN. Format depends on type: timeOfDay: HH:mm, dayOfWeek: 1-7 (Monday-Sunday), dayOfMonth: 1-31, specificDate: yyyy-MM-dd (if includeYear=true) or MM-dd (if includeYear=false).
    Min string `json:"min"`


    // Max - The maximum value of the range. Required for the operator BETWEEN. Format follows the same rules as 'min'.
    Max string `json:"max"`


    // InSet - A set of values that the date/ time data should be in. Required for the IN operator. Format depends on type: dayOfWeek: 1-7 (Monday-Sunday), dayOfMonth: 1-31, and/ or LAST_DAY, ODD_DAY, EVEN_DAY,specificDate: yyyy-MM-dd (if includeYear=true) or MM-dd (if includeYear=false).
    InSet []string `json:"inSet"`

}

// String returns a JSON representation of the model
func (o *Timeanddatesubconditionrange) String() string {
    
    
     o.InSet = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeanddatesubconditionrange) MarshalJSON() ([]byte, error) {
    type Alias Timeanddatesubconditionrange

    if TimeanddatesubconditionrangeMarshalled {
        return []byte("{}"), nil
    }
    TimeanddatesubconditionrangeMarshalled = true

    return json.Marshal(&struct {
        
        Min string `json:"min"`
        
        Max string `json:"max"`
        
        InSet []string `json:"inSet"`
        *Alias
    }{

        


        


        
        InSet: []string{""},
        

        Alias: (*Alias)(u),
    })
}

