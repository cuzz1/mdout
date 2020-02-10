package model

// Theme 包含主题信息以及下载地址
type Theme struct {
	Name         string
	Author       string
	Des          string
	RepoURL      string
	ZipURL       string
	DownloadType string
}
