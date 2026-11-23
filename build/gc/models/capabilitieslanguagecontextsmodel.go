package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitieslanguagecontextsmodelMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitieslanguagecontextsmodelDud struct { 
    


    


    

}

// Capabilitieslanguagecontextsmodel
type Capabilitieslanguagecontextsmodel struct { 
    // Gender
    Gender []string `json:"gender"`


    // Quantity
    Quantity []string `json:"quantity"`


    // VarCase
    VarCase []string `json:"case"`

}

// String returns a JSON representation of the model
func (o *Capabilitieslanguagecontextsmodel) String() string {
     o.Gender = []string{""} 
     o.Quantity = []string{""} 
     o.VarCase = []string{""} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitieslanguagecontextsmodel) MarshalJSON() ([]byte, error) {
    type Alias Capabilitieslanguagecontextsmodel

    if CapabilitieslanguagecontextsmodelMarshalled {
        return []byte("{}"), nil
    }
    CapabilitieslanguagecontextsmodelMarshalled = true

    return json.Marshal(&struct {
        
        Gender []string `json:"gender"`
        
        Quantity []string `json:"quantity"`
        
        VarCase []string `json:"case"`
        *Alias
    }{

        
        Gender: []string{""},
        


        
        Quantity: []string{""},
        


        
        VarCase: []string{""},
        

        Alias: (*Alias)(u),
    })
}

