package main

import (
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

const k = 1024

type RequiredFile struct {
	Path        string
	ContentType string
	MinBytes    int
}

func (rf RequiredFile) Test(t *testing.T) {
	path := filepath.Join("..", "public", rf.Path)
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	fi, err := f.Stat()
	if err != nil {
		t.Fatal(err)
	}
	{
		got := fi.Size()
		want := int64(rf.MinBytes)
		if got < want {
			t.Fatalf("size: got %d bytes, want %d bytes or more", got, want)
		}
	}

	b := make([]byte, 512)
	_, err = f.Read(b)
	if err != nil {
		t.Fatal(err)
	}
	{
		got := http.DetectContentType(b)
		want := rf.ContentType
		if got != want {
			t.Fatalf("content type: got %q, want %q", got, want)
		}
	}
}

var files = []RequiredFile{
	// Sitemap:
	{
		Path:        "sitemap.xml",
		ContentType: "text/xml; charset=utf-8",
		MinBytes:    200,
	},
	// Icons:
	{
		Path:        "favicon.ico",
		ContentType: "image/vnd.microsoft.icon",
		MinBytes:    1 * k,
	},
	{
		Path:        "apple-touch-icon.png",
		ContentType: "image/png",
		MinBytes:    1 * k,
	},
	{
		Path:        "touch-icon.png",
		ContentType: "image/png",
		MinBytes:    1 * k,
	},
	// Pages:
	{
		Path:        "index.html",
		ContentType: "text/html; charset=utf-8",
		MinBytes:    1,
	},
	{
		Path:        "404.html",
		ContentType: "text/html; charset=utf-8",
		MinBytes:    512 + 1, // Ensure > 512 bytes or browsers may ignore it
	},
}

func TestRequiredFiles(t *testing.T) {
	for _, rf := range files {
		t.Run(rf.Path, rf.Test)
	}
}
