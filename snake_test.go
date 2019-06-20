package strcase

import "testing"

func TestToSnake(t *testing.T) {
	table := map[string]string{
		"ValidIDs":     "valid_ids",
		"a":            "a",
		"A":            "a",
		"AB":           "ab",
		"JOB_ID":       "job_id",
		"ABCJobID":     "abc_job_id",
		"ABCD":         "abcd",
		"LinuxMOTD":    "linux_motd",
		"APIResponse":  "api_response",
		"HTTPReferer":  "http_referer",
		"http_referer": "http_referer",
		"CapitalI":     "capital_i",
		"PersonId":     "person_id",
		"PersonID":     "person_id",
		"iPad":         "i_pad",
	}
	for orig, target := range table {
		if ToSnake(orig) != target {
			t.Errorf("Expected '%s' => '%s'. Got '%s'", orig, target, ToSnake(orig))
		}
	}
}
