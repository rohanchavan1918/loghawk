package parser

import (
	"errors"
	"fmt"
	"loghawk/config"
	"loghawk/models"
	"regexp"
	"strings"
)

func ParseLogs(tag string, log string) {
	tagData, err := GetTagRules(tag)
	if err != nil {
		fmt.Printf("Failed to get tag Rules ")
	}
	if ShouldSendAlert(tagData.Rules, log) {
		SendSlackAlert("Alert for : "+log, tagData.Name, tagData.SlackUrl)
	}
}

func ShouldSendAlert(tagRules []models.TagRule, log string) bool {
	for _, rule := range tagRules {
		if rule.MatchType == "contains" {
			if CheckContains(rule.MatchValue, log) {
				fmt.Println("CheckContains requested ")
				return true
			}
		} else if rule.MatchType == "starts_with" {
			if CheckStartsWith(rule.MatchValue, log) {
				return true
			}
		} else if rule.MatchType == "ends_with" {
			if CheckEndsWith(rule.MatchValue, log) {
				return true
			}
		} else if rule.MatchType == "regex" {
			if CheckRegex(rule.MatchValue, log) {
				return true
			}
		}
	}
	return false
}

func GetTagRules(tagName string) (models.Tag, error) {
	// Get all rules for tags for that tag

	db := config.DB
	tag := models.Tag{Name: tagName}
	result := db.Preload("Rules").First(&tag)
	if result.Error != nil {
		return models.Tag{}, errors.New("tag does not exist")
	}
	return models.Tag{}, nil
}

func CheckContains(match_value, log string) bool {

	// orMatches := strings.Split(match_value, "|")
	// andMatches := strings.Split(match_value, "&")

	// if len(orMatches) && len(andMatches)

	// fmt.Println(len(orMatches), len(andMatches))
	return strings.Contains(log, match_value)
}

func CheckStartsWith(match_value, log string) bool {

	// orMatches := strings.Split(match_value, "|")
	// andMatches := strings.Split(match_value, "&")

	// if len(orMatches) && len(andMatches)

	// fmt.Println(len(orMatches), len(andMatches))
	return strings.HasPrefix(log, match_value)
}

func CheckEndsWith(match_value, log string) bool {

	// orMatches := strings.Split(match_value, "|")
	// andMatches := strings.Split(match_value, "&")

	// if len(orMatches) && len(andMatches)

	// fmt.Println(len(orMatches), len(andMatches))
	return strings.HasSuffix(log, match_value)
}

func CheckRegex(match_value, log string) bool {
	re := regexp.MustCompile(match_value)
	return re.MatchString(log)
}
