package public

import "embed"

//go:embed dist/*
var distFS embed.FS

func GetWWWEmbed() *embed.FS {
	return &distFS
}
