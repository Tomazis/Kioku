package internal

//go:generate mockgen -source=./api/api.go -package api -aux_files=github.com/tomazis/kioku/server/srv-frontend-api/internal/api=api/api_kanji.go,github.com/tomazis/kioku/server/srv-frontend-api/internal/api=api/api_word.go -destination=./mocks/repo_mock.go
