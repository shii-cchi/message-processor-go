package dto

type StatsDto struct {
	MessagesCount          int64 `json:"messages_count"`
	ProcessedMessagesCount int64 `json:"processed_messages_count"`
}
