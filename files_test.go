package dropbox

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiles_Upload(t *testing.T) {
	c := client()

	file, err := os.Open("Readme.md")
	assert.NoError(t, err, "error opening file")
	defer file.Close()

	out, err := c.Files.Upload(&UploadInput{
		Mute:   true,
		Mode:   WriteModeOverwrite,
		Path:   "/Readme.md",
		Reader: file,
	})

	assert.NoError(t, err, "error uploading file")
	assert.Equal(t, "/readme.md", out.PathLower)
}

func TestFiles_Download(t *testing.T) {
	c := client()

	out, err := c.Files.Download(&DownloadInput{"/Readme.md"})

	assert.NoError(t, err, "error downloading")
	defer out.Body.Close()

	remote, err := ioutil.ReadAll(out.Body)
	assert.NoError(t, err, "error reading remote")

	local, err := ioutil.ReadFile("Readme.md")
	assert.NoError(t, err, "error reading local")

	assert.Equal(t, local, remote, "Readme.md mismatch")
}

func TestFiles_GetMetadata(t *testing.T) {
	c := client()

	out, err := c.Files.GetMetadata(&GetMetadataInput{
		Path: "/Readme.md",
	})

	assert.NoError(t, err)
	assert.Equal(t, "file", out.Tag)
}

func TestFiles_ListFolder(t *testing.T) {
	c := client()

	out, err := c.Files.ListFolder(&ListFolderInput{
		Path: "",
	})

	assert.NoError(t, err)
	assert.Len(t, out.Entries, 1)
	assert.Equal(t, "/readme.md", out.Entries[0].PathLower)
}

func TestFiles_Search(t *testing.T) {
	c := client()

	out, err := c.Files.Search(&SearchInput{
		Query: "Readme",
	})

	assert.NoError(t, err)
	assert.Len(t, out.Matches, 1)
}

func TestFiles_Delete(t *testing.T) {
	c := client()

	out, err := c.Files.Delete(&DeleteInput{
		Path: "/Readme.md",
	})

	assert.NoError(t, err)
	assert.Equal(t, "/readme.md", out.PathLower)
}

func TestFiles_GetPreview(t *testing.T) {
	c := client()

	_, err := c.Files.GetPreview(&GetPreviewInput{
		Path: "/sample.ppt",
	})

	assert.NoError(t, err)
}
