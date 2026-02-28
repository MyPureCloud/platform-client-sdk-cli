package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OrganizationcallmetricsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OrganizationcallmetricsDud struct { 
    


    


    


    


    


    


    


    


    

}

// Organizationcallmetrics
type Organizationcallmetrics struct { 
    // Usage - The current usage percentage of the organization's call capacity.
    Usage float64 `json:"usage"`


    // AutoScalingTriggerPercentage - The autoscaling trigger percentage of the organization's call capacity.
    AutoScalingTriggerPercentage float64 `json:"autoScalingTriggerPercentage"`


    // CpuIntensity - The current compute intensity of the organization's call capacity.
    CpuIntensity string `json:"cpuIntensity"`


    // MemoryIntensity - The current memory intensity of the organization's call capacity.
    MemoryIntensity string `json:"memoryIntensity"`


    // ConcurrentCallCount - The current number of concurrent calls in the organization.
    ConcurrentCallCount int `json:"concurrentCallCount"`


    // ConcurrentCallSessionCount - The current number of concurrent call sessions in the organization.
    ConcurrentCallSessionCount int `json:"concurrentCallSessionCount"`


    // MaxCallCapacity - The maximum number of concurrent calls allowed in the organization.
    MaxCallCapacity int `json:"maxCallCapacity"`


    // MaxCallSessionCapacity - The maximum number of concurrent call sessions allowed in the organization.
    MaxCallSessionCapacity int `json:"maxCallSessionCapacity"`


    // AutoScaleInProgress - The autoscaling status of the organization's call capacity.
    AutoScaleInProgress string `json:"autoScaleInProgress"`

}

// String returns a JSON representation of the model
func (o *Organizationcallmetrics) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Organizationcallmetrics) MarshalJSON() ([]byte, error) {
    type Alias Organizationcallmetrics

    if OrganizationcallmetricsMarshalled {
        return []byte("{}"), nil
    }
    OrganizationcallmetricsMarshalled = true

    return json.Marshal(&struct {
        
        Usage float64 `json:"usage"`
        
        AutoScalingTriggerPercentage float64 `json:"autoScalingTriggerPercentage"`
        
        CpuIntensity string `json:"cpuIntensity"`
        
        MemoryIntensity string `json:"memoryIntensity"`
        
        ConcurrentCallCount int `json:"concurrentCallCount"`
        
        ConcurrentCallSessionCount int `json:"concurrentCallSessionCount"`
        
        MaxCallCapacity int `json:"maxCallCapacity"`
        
        MaxCallSessionCapacity int `json:"maxCallSessionCapacity"`
        
        AutoScaleInProgress string `json:"autoScaleInProgress"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

