package command_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/cli/command"
	"github.com/stretchr/testify/assert"
)

func TestCreateRootCmd(t *testing.T) {
	rc := command.CreateRootCmd()
	assert.Empty(t, rc.Aliases)
}
