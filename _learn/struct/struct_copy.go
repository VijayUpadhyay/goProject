package main

import "fmt"

type TestDetails struct {
	sectionName string
	PBM         *string
}

func main() {
	kv := make(map[string]interface{})
	var pbm string
	for i := 0; i < 2; i++ {
		kv = getResult(i)
		repoData := TestDetails{sectionName: "GrossMargin"}
		setPBMValue(&pbm, kv, &repoData)
		fmt.Println(repoData)
	}
}

func getResult(i int) map[string]interface{} {
	kv := make(map[string]interface{})
	kv["a"] = "a"
	if i == 0 {
		kv["PBM"] = "pbm"
	} else {
		kv["bb"] = "b"
	}
	return kv
}

func setPBMValue(pbm *string, kv map[string]interface{}, reportDetails *TestDetails) string {
	pbmType, ok := kv["PBM"]
	//fmt.Println("pbm: ", *pbm)
	// && isBlank(*reportDetails.PBM)
	if ok && pbmType != nil {
		*pbm = fmt.Sprintf("%v", pbmType)
		reportDetails.PBM = pbm
	}
	return *pbm
}

func isBlank(s string) bool {
	return s == ""
}
