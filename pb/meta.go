package pb

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	_headSize = 265
)

// nolint: gochecknoglobals
var _documents = [...]string{".xlsx", ".pptx", ".docx"}

// nolint: gochecknoglobals
var caser = cases.Title(language.English)

func GetMetaByReader(readCloser io.ReadCloser) (Meta, error) {
	defer readCloser.Close()

	meta := Meta_Unknown
	head := make([]byte, _headSize)

	if _, err := readCloser.Read(head); err != nil {
		return meta, err
	}

	kind, err := filetype.Match(head)
	if err != nil {
		return meta, err
	}

	switch {
	case filetype.IsImage(head):
		return Meta_Image, nil
	case filetype.IsDocument(head), kind.MIME.Subtype == "pdf":
		return Meta_Documents, nil
	case filetype.IsVideo(head):
		return Meta_Video, nil
	case filetype.IsAudio(head):
		return Meta_Audio, nil
	case filetype.IsArchive(head):
		return Meta_Archive, nil
	default:
	}

	if value, has := Meta_value[caser.String(kind.Extension)]; has {
		meta = Meta(value)
	}

	return meta, nil
}

func GetMetaByExt(path string) Meta {
	kind := filetype.GetType(strings.ToLower(filepath.Ext(path)))

	if value, has := Meta_value[caser.String(kind.MIME.Type)]; has {
		return Meta(value)
	}

	return Meta_Unknown
}

func GetMeta(path string) (Meta, error) {
	file, err := os.Open(path)
	if err != nil {
		return Meta_Unknown, err
	}

	meta, err := GetMetaByReader(file)
	if err != nil {
		return meta, err
	}

	if meta == Meta_Unknown {
		meta = GetMetaByExt(path)
	}

	if meta == Meta_Archive {
		ext := filepath.Ext(path)

		for _, doc := range _documents {
			if strings.EqualFold(doc, ext) {
				return Meta_Documents, nil
			}
		}
	}

	return meta, nil
}

// nolint: gochecknoinits
func init() {
	filetype.AddType(".go", "Golang")
	filetype.AddType(".java", "Java")
	filetype.AddType(".js", "JavaScript")
	filetype.AddType(".svg", "Image")
	filetype.AddType(".xcf", "Image")
	filetype.AddType(".rmvb", "Video")
	filetype.AddType(".rm", "Video")
	filetype.AddType(".mpg", "Video")
	filetype.AddType(".mov", "Video")
	filetype.AddType(".chm", "Documents")
	filetype.AddType(".txt", "Documents")
	filetype.AddType(".iso", "Archive")
	filetype.AddType(".mp3", "Audio")
}
