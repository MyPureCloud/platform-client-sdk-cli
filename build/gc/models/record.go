package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecordMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecordDud struct { 
    


    


    

}

// Record
type Record struct { 
    // Name - The name of the record.
    Name string `json:"name"`


    // VarType - The type of the record. (Example values:  MX, TXT, CNAME)
    VarType string `json:"type"`


    // Value - The value of the record.
    Value string `json:"value"`

}

// String returns a JSON representation of the model
func (o *Record) String() string {
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Record) MarshalJSON() ([]byte, error) {
    type Alias Record

    if RecordMarshalled {
        return []byte("{}"), nil
    }
    RecordMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        VarType string `json:"type"`
        
        Value string `json:"value"`
        *Alias
    }{

        


        


        

        Alias: (*Alias)(u),
    })
}

