package agent

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/sipeed/picoclaw/pkg/session"
)

func NewAgentCommand() *cobra.Command {
	var (
		message    string
		sessionKey string
		model      string
		debug      bool
	)

	cmd := &cobra.Command{
		Use:   "agent",
		Short: "Interact with the agent directly",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if cmd.Flags().Changed("session") {
				sessionKey = cliSessionKey(sessionKey)
			}
			return agentCmd(message, sessionKey, model, debug)
		},
	}

	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")
	cmd.Flags().StringVarP(&message, "message", "m", "", "Send a single message (non-interactive mode)")
	cmd.Flags().StringVarP(&sessionKey, "session", "s", "cli:default", "Session key")
	cmd.Flags().StringVarP(&model, "model", "", "", "Model to use")

	return cmd
}

// cliSessionKey maps a user-chosen --session name to an explicit session key.
// Plain names like "work" are otherwise ignored by the agent loop, which only
// honors opaque (sk_v1_...) or legacy (agent:...) keys, so every CLI run would
// share one session.
func cliSessionKey(name string) string {
	name = strings.TrimSpace(name)
	if name == "" || session.IsExplicitSessionKey(name) {
		return name
	}
	return session.BuildOpaqueSessionKey("cli:" + name)
}
