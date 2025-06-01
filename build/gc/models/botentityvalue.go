package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BotentityvalueMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BotentityvalueDud struct { 
    


    


    


    

}

// Botentityvalue
type Botentityvalue struct { 
    // Name - The name of the entity.
    Name string `json:"name"`


    // VarType - The data type of the entity. Valid types include: String, Integer, Decimal, Boolean, Duration, Datetime, Currency, StringCollection, IntegerCollection, DecimalCollection, BooleanCollection, DurationCollection, DatetimeCollection, CurrencyCollection.
    VarType string `json:"type"`


    // Value - The string value of the entity for simple types. Required when using non-collection types. Format depends on the 'type' field: String: \"any text\"; Integer: whole number (e.g., \"42\"); Decimal: number with optional decimal point (e.g., \"42.5\"); Boolean: \"true\" or \"false\"; Duration: ISO-8601 duration format (e.g., \"PT1H30M\" for 1 hour and 30 minutes); Datetime: ISO-8601 datetime format (e.g., \"2023-04-15T14:30:00Z\"); Currency: JSON object with 'amount' and 'code' fields as an escaped JSON string (e.g., \"{\\\"amount\\\":10.50,\\\"code\\\":\\\"USD\\\"}\" - note the escaped quotes).
    Value string `json:"value"`


    // Values - The collection values for collection types. Required when using collection types. Each value must follow the format of its base type: StringCollection: array of strings; IntegerCollection: array of integer strings (e.g., [\"1\", \"2\", \"3\"]); DecimalCollection: array of decimal strings (e.g., [\"1.5\", \"2.0\", \"3.75\"]); BooleanCollection: array of boolean strings (e.g., [\"true\", \"false\"]); DurationCollection: array of ISO-8601 duration strings; DatetimeCollection: array of ISO-8601 datetime strings; CurrencyCollection: array of escaped JSON currency object strings (e.g., [\"{\\\"amount\\\":10.50,\\\"code\\\":\\\"USD\\\"}\", \"{\\\"amount\\\":25.00,\\\"code\\\":\\\"EUR\\\"}\"] - note the escaped quotes).
    Values []string `json:"values"`

}

// String returns a JSON representation of the model
func (o *Botentityvalue) String() string {
    
    
    
     o.Values = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Botentityvalue) MarshalJSON() ([]byte, error) {
    type Alias Botentityvalue

    if BotentityvalueMarshalled {
        return []byte("{}"), nil
    }
    BotentityvalueMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        
        Values []string `json:"values"`
        *Alias
    }{

        


        


        


        
        Values: []string{""},
        

        Alias: (*Alias)(u),
    })
}

