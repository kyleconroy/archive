package twitter

import "encoding/json"

// TODO: Parse the Timezone string into a time.Location
type AgeInfo struct {
	Age       []string `json:"age"`
	BirthDate string   `json:"birthDate"`
}

type ageWrapper struct {
	Meta ageMeta `json:"ageMeta"`
}

type ageMeta struct {
	Obj ageInfo `json:"inferredAgeInfo"`
}

type ageInfo AgeInfo

func (ai *AgeInfo) UnmarshalJSON(b []byte) error {
	var wrapper ageWrapper
	err := json.Unmarshal(b, &wrapper)
	if err != nil {
		return err
	}
	*ai = AgeInfo(wrapper.Meta.Obj)
	return nil
}
