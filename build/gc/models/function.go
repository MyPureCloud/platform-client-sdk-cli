package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    FunctionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type FunctionDud struct { 
    Id string `json:"id"`


    Name string `json:"name"`


    


    DateCreated time.Time `json:"dateCreated"`


    ZipId string `json:"zipId"`


    


    


    

}

// Function - Action function settings.
type Function struct { 
    


    


    // Description - Description of the function. Limit 255 characters.
    Description string `json:"description"`


    


    


    // Handler - Handler entry point into zip file to execute function. Should be path within upload function package to the invocation module without language extension, followed by a period and then exported invocation method name. e.g. 'src/index.handler'
    Handler string `json:"handler"`


    // Runtime - Runtime required for execution. Valid runtimes change over time as versions are deprecated. Use /api/v2/integrations/actions/functions/runtimes for current list.
    Runtime string `json:"runtime"`


    // TimeoutSeconds - Execution timeout to apply to function. Value is in seconds. Range allowed 1 to 60. Default value 15 seconds.
    TimeoutSeconds int `json:"timeoutSeconds"`

}

// String returns a JSON representation of the model
func (o *Function) String() string {
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Function) MarshalJSON() ([]byte, error) {
    type Alias Function

    if FunctionMarshalled {
        return []byte("{}"), nil
    }
    FunctionMarshalled = true

    return json.Marshal(&struct {
        
        Description string `json:"description"`
        
        Handler string `json:"handler"`
        
        Runtime string `json:"runtime"`
        
        TimeoutSeconds int `json:"timeoutSeconds"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

