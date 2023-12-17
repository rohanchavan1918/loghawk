package parser

import (
	"errors"
	"fmt"
	"loghawk/config"
	"loghawk/models"
	"regexp"
	"strings"
	"time"
)

func ParseLogs(tag string, log string) {
	fmt.Println("Started Parsing log for tag > ", tag)
	tagData, err := GetTagRules(tag)
	fmt.Println("TagData > ", tagData)
	if err != nil {
		fmt.Printf("Failed to get tag Rules ")
	}
	if ShouldSendAlert(tagData.Rules, log) {
		fmt.Println("Rule matched for tag > ", tag)
		SendSlackAlert("Alert for : "+log, tagData.Name, tagData.SlackUrl)
		SaveLog(log, int(tagData.ID))

	} else {
		fmt.Println("No rules matched.")
	}
}

func SaveLog(log string, tagId int) {
	logData := models.Log{Message: log, TagID: uint(tagId), CreatedAt: time.Now()}
	if err := config.DB.Create(&logData).Error; err != nil {
		fmt.Println("Failed to add logs : ", err)
	} else {
		fmt.Println("Log Saved successfully.")
	}
}

func ShouldSendAlert(tagRules []models.TagRule, log string) bool {
	fmt.Println("Evaluating if data should be sent or not.")
	for _, rule := range tagRules {
		fmt.Println("Parsing Rule > ", rule)
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
	tag := models.Tag{Tag: tagName}
	result := db.Preload("Rules").First(&tag)
	if result.Error != nil {
		return models.Tag{}, errors.New("tag does not exist")
	}
	return tag, nil
}

func CheckContains(match_value, log string) bool {

	// orMatches := strings.Split(match_value, "|")
	// andMatches := strings.Split(match_value, "&")

	// if len(orMatches) && len(andMatches)

	fmt.Println("Checking if ", match_value, " is present in ", log)

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
