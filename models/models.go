package models

import "time"

type GoSec struct {
	GolangErrors struct {
	} `json:"Golang errors"`
	Issues []struct {
		Severity   string `json:"severity"`
		Confidence string `json:"confidence"`
		RuleID     string `json:"rule_id"`
		Details    string `json:"details"`
		File       string `json:"file"`
		Code       string `json:"code"`
		Line       string `json:"line"`
	} `json:"Issues"`
	Stats struct {
		Files int `json:"files"`
		Lines int `json:"lines"`
		Nosec int `json:"nosec"`
		Found int `json:"found"`
	} `json:"Stats"`
}

type Chat struct {
	Text   string `json:"text"`
	Thread struct {
		Name string `json:"name"`
	} `json:"thread"`
}

type ChatResponse struct {
	Name   string `json:"name"`
	Sender struct {
		Name        string `json:"name"`
		DisplayName string `json:"displayName"`
		AvatarURL   string `json:"avatarUrl"`
		Email       string `json:"email"`
		DomainID    string `json:"domainId"`
		Type        string `json:"type"`
	} `json:"sender"`
	Text        string        `json:"text"`
	Cards       []interface{} `json:"cards"`
	PreviewText string        `json:"previewText"`
	Annotations []interface{} `json:"annotations"`
	Thread      struct {
		Name string `json:"name"`
	} `json:"thread"`
	Space struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		DisplayName string `json:"displayName"`
	} `json:"space"`
	FallbackText string    `json:"fallbackText"`
	ArgumentText string    `json:"argumentText"`
	CreateTime   time.Time `json:"createTime"`
}

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Status  string `json:"status"`
	} `json:"error"`
}



type SortBySeverity struct {
	Low    int
	Medium int
	High   int
}
