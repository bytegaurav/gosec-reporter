package processors

import "Reporter/models"

func SortIssuesBySeverity(logObject models.GoSec) models.SortBySeverity {

	result := models.SortBySeverity{
		Low:    0,
		Medium: 0,
		High:   0,
	}

	for _, item := range logObject.Issues {

		switch item.Severity {
		case "LOW":
			result.Low++
		case "MEDIUM":
			result.Medium++
		case "HIGH":
			result.High++
		}
	}

	return result

}

func SortIssuesByMessage(logObject models.GoSec) (map[string]int, int) {

	result := make(map[string]int)
	largest := 0
	for _, item := range logObject.Issues {
		key := item.RuleID + ": " + item.Details
		count := result[key]
		count++
		result[item.RuleID+": "+item.Details] = count

		if len(key) > largest {
			largest = len(key)
		}

	}
	return result, largest

}
