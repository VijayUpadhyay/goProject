package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"sync"
)

var (
	rriFormula = "( (FutureValue ^ (1 / NPER ) ) / PresentValue) - 1"
)

type (
	VAConfig struct {
		Constants              map[string]interface{}                           `json:"Constants"`
		LookupTables           map[string]map[string]LookUpTableInfo            `json:"LookupTables"`
		LookUpMapping          map[string]map[string]LookUpMappingInfo          `json:"LookUpMapping"`
		LookUpRangeMapping     map[string]map[string]LookUpRangeMappingInfo     `json:"LookUpRangeMapping"`
		EvaluationRangeMapping map[string]map[string]EvaluationRangeMappingInfo `json:"EvaluationRangeMapping"`
		RRIMapping             map[string]map[string]RRIMappingInfo             `json:"RRIMapping"`
		TextMapping            map[string]map[string]TextMappingInfo            `json:"TextMapping"`
		CalculationFields      []CalculationInfo                                `json:"Calculation"`
	}

	CalculationInfo struct {
		Name            string `json:"name"`
		Formula         string `json:"formula"`
		CalculationType string `json:"calculationType"`
	}

	LookUpMappingInfo struct {
		Table string `json:"Table"`
		Key   string `json:"Key"`
		Value string `json:"Value"`
	}

	LookUpTableInfo struct {
		Table string `json:"Table"`
		Key   string `json:"Key"`
		Value string `json:"Value"`
	}

	RRIMappingInfo struct {
		NPER         int    `json:"NPER"`
		PresentValue string `json:"PresentValue"`
		FutureValue  string `json:"FutureValue"`
	}

	TextMappingInfo struct {
		Text    []string `json:"Text"`
		Replace []string `json:"Replace"`
	}

	LookUpRangeMappingInfo struct {
		TableEx struct {
			Table   string `json:"Table"`
			Replace string `json:"Replace"`
		} `json:"TableEx"`
		Keys   []string `json:"Keys"`
		Values string   `json:"Values"`
	}

	EvaluationRangeMappingInfo struct {
		RangeKey   string `json:"range_key"`
		Multiplier string `json:"multiplier"`
		Formula    string `json:"formula"`
	}
)

type verifyResultData struct {
	content  *string
	vaConfig *VAConfig
	syncMap  *sync.Map
}

//go:embed app.json
var vaCalcConfig string

var vaConfig VAConfig
var wg sync.WaitGroup

func init() {
	if err := json.Unmarshal([]byte(vaCalcConfig), &vaConfig); err != nil {
		fmt.Printf("failed parse value assessment configuration for report fields evaluations. Error: %v \n", err)
	}
}

func main() {
	/* fmt.Println(vaCalcConfig)
	fmt.Println(vaConfig) */

	fmt.Println("********************************************************")
	fmt.Println(vaConfig.LookupTables)
}

func structToMap(structData any) (temp map[string]interface{}) {
	byteData, _ := json.Marshal(structData)
	if err := json.Unmarshal(byteData, &temp); err != nil {
		fmt.Println(err.Error())
	}
	return
}

func (data *verifyResultData) convertAndSaveVAConstants(structData any) {
	val := structToMap(structData)
	data.updateVerifyResultValueMap(val)
}

func (data *verifyResultData) updateVerifyResultValueMap(m map[string]interface{}) {
	for k, v := range m {
		data.syncMap.Store(k, v)
	}
	wg.Done()
}
