package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    QuerydivisionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type QuerydivisionDud struct { }

// Querydivision
type Querydivision struct { }

// String returns a JSON representation of the model
func (o *Querydivision) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Querydivision) MarshalJSON() ([]byte, error) {
    type Alias Querydivision

    if QuerydivisionMarshalled {
        return []byte("{}"), nil
    }
    QuerydivisionMarshalled = true

    return json.Marshal(&struct { 
        *Alias
    }{
        
        Alias: (*Alias)(u),
    })
}

