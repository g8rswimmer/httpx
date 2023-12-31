package jbody

type ParameterValidation struct {
	Object      *ObjectValidator      `json:"object_validator"`
	ObjectArray *ObjectArrayValidator `json:"object_array_validator"`
	String      *StringValidator      `json:"string_validator"`
	StringArray *StringArrayValidator `json:"string_array_validator"`
	Number      *NumberValidator      `json:"number_validator"`
	NumberArray *NumberArrayValidator `json:"number_array_validator"`
	Time        *TimeValidator        `json:"time_validator"`
	TimeArray   *TimeArrayValidator   `json:"time_array_validator"`
	Boolean     *BooleanValidator     `json:"boolean_validator"`
}

type ParameterProperties struct {
	Validation ParameterValidation `json:"validation"`
}
