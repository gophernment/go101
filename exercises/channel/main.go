package main

import (
	"fmt"
	"time"
)

const (
	bloodTestDelay time.Duration = 800
	urineTestDelay time.Duration = 500
)

func main() {
	begin := time.Now()
	patientTest := map[string]testResult{
		"AAAA": {},
		"BBBB": {},
		"CCCC": {},
		"DDDD": {},
		"EEEE": {},
		"FFFF": {},
		"GGGG": {},
		"HHHH": {},
		"IIII": {},
		"JJJJ": {},
	}

	for patient := range patientTest {
		t := testResult{}
		t.bloodTest = bloodTest(patient)
		t.urineTest = urineExam(patient)
		patientTest[patient] = t
	}

	fmt.Println(patientTest)
	fmt.Println(time.Since(begin))
}

type testResult struct {
	bloodTest bool
	urineTest bool
}

func bloodTest(patient string) bool {
	time.Sleep(bloodTestDelay * time.Millisecond)
	return true
}
func urineExam(patient string) bool {
	time.Sleep(urineTestDelay * time.Millisecond)
	return true
}
