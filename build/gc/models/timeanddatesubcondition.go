package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeanddatesubconditionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeanddatesubconditionDud struct { 
    


    


    


    


    


    

}

// Timeanddatesubcondition
type Timeanddatesubcondition struct { 
    // VarType - The type of time/date sub-condition.
    VarType string `json:"type"`


    // Operator - The operator to use for comparison.
    Operator string `json:"operator"`


    // Inverted - If true, inverts the result of evaluating this sub-condition. Default is false.
    Inverted bool `json:"inverted"`


    // IncludeYear - If true, includes year in date comparison for specificDate type. When false, only month and day are compared. Default is true. Only applicable for specificDate type.
    IncludeYear bool `json:"includeYear"`


    // ThresholdValue - Threshold value for BEFORE or AFTER operators. Format depends on type: timeOfDay: HH:mm, dayOfWeek: 1-7 (Monday-Sunday), dayOfMonth: 1-31 and/ or LAST_DAY, ODD_DAY, EVEN_DAY, specificDate: yyyy-MM-dd (if includeYear=true) or MM-dd (if includeYear=false). For single-value comparison, use a list with one element.
    ThresholdValue string `json:"thresholdValue"`


    // VarRange - A range of values for BETWEEN and IN operators. Format follows the same rules as 'thresholdValue'.
    VarRange Timeanddatesubconditionrange `json:"range"`

}

// String returns a JSON representation of the model
func (o *Timeanddatesubcondition) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeanddatesubcondition) MarshalJSON() ([]byte, error) {
    type Alias Timeanddatesubcondition

    if TimeanddatesubconditionMarshalled {
        return []byte("{}"), nil
    }
    TimeanddatesubconditionMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        
        Operator string `json:"operator"`
        
        Inverted bool `json:"inverted"`
        
        IncludeYear bool `json:"includeYear"`
        
        ThresholdValue string `json:"thresholdValue"`
        
        VarRange Timeanddatesubconditionrange `json:"range"`
        *Alias
    }{

        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

