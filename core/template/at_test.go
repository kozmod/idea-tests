package template

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
	"text/template"
)

func TestAt(t *testing.T) { // at = @
	tpl, _ := template.New("text").Parse(`{{range .Users}} <@{{.}}> {{end}}`)
	var buffer bytes.Buffer
	err := tpl.Execute(&buffer, struct {
		Users []string
	}{
		Users: []string{"user1", "user2"},
	})
	assert.NoError(t, err)
	assert.Equal(t, " <@user1>  <@user2> ", buffer.String())
}
