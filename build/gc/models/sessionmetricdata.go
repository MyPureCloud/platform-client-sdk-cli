package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SessionmetricdataMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SessionmetricdataDud struct { 
    


    


    


    

}

// Sessionmetricdata
type Sessionmetricdata struct { 
    // Model - Model to be used for internal evaluation purposes
    Model Modeldata `json:"model"`


    // Computed - Forecasted data for the requested session
    Computed Computeddata `json:"computed"`


    // Historical - Historical data for the requested session
    Historical Historicaldata `json:"historical"`


    // ModelMetaData - Contains forecast meta data
    ModelMetaData Modelmetadata `json:"modelMetaData"`

}

// String returns a JSON representation of the model
func (o *Sessionmetricdata) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Sessionmetricdata) MarshalJSON() ([]byte, error) {
    type Alias Sessionmetricdata

    if SessionmetricdataMarshalled {
        return []byte("{}"), nil
    }
    SessionmetricdataMarshalled = true

    return json.Marshal(&struct {
        
        Model Modeldata `json:"model"`
        
        Computed Computeddata `json:"computed"`
        
        Historical Historicaldata `json:"historical"`
        
        ModelMetaData Modelmetadata `json:"modelMetaData"`
        *Alias
    }{

        


        


        


        

        Alias: (*Alias)(u),
    })
}

