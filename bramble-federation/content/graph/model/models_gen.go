// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type ContentItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ContentItemWithScoreCard struct {
	ID    string     `json:"id"`
	Title string     `json:"title"`
	Score *ScoreCard `json:"score"`
}

type ScoreCard struct {
	ID string `json:"id"`
}

type Service struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Schema  string `json:"schema"`
}
