package agent

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sipeed/picoclaw/pkg/session"
)

func TestNewAgentCommand(t *testing.T) {
	cmd := NewAgentCommand()

	require.NotNil(t, cmd)

	assert.Equal(t, "agent", cmd.Use)
	assert.Equal(t, "Interact with the agent directly", cmd.Short)

	assert.Len(t, cmd.Aliases, 0)
	assert.False(t, cmd.HasSubCommands())

	assert.Nil(t, cmd.Run)
	assert.NotNil(t, cmd.RunE)

	assert.Nil(t, cmd.PersistentPreRun)
	assert.Nil(t, cmd.PersistentPostRun)

	assert.True(t, cmd.HasFlags())

	assert.NotNil(t, cmd.Flags().Lookup("debug"))
	assert.NotNil(t, cmd.Flags().Lookup("message"))
	assert.NotNil(t, cmd.Flags().Lookup("session"))
	assert.NotNil(t, cmd.Flags().Lookup("model"))
}

func TestCLISessionKey(t *testing.T) {
	work := cliSessionKey("work")
	assert.True(t, session.IsExplicitSessionKey(work))
	assert.Equal(t, work, cliSessionKey(" work "))
	assert.NotEqual(t, work, cliSessionKey("home"))

	assert.Equal(t, "agent:main:x", cliSessionKey("agent:main:x"))
	assert.Equal(t, work, cliSessionKey(work))
	assert.Equal(t, "", cliSessionKey(""))
}
