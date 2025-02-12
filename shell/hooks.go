package shell

import (
	_ "embed" // yes
	"fmt"
	"path/filepath"

	"github.com/cashapp/hermit/errors"
	"github.com/cashapp/hermit/internal/system"
)

//go:embed files/sh_common_hooks.sh
var commonHooks string

const (
	hookStartMarker = "# Generated by Hermit; START; DO NOT EDIT."
	hookEndMarker   = "# Generated by Hermit; END; DO NOT EDIT."
)

func installationScript(shell string) string {
	return fmt.Sprintf(
		`HERMIT_ROOT_BIN="${HERMIT_ROOT_BIN:-"$HOME/bin/hermit"}" `+
			`eval "$(test -x $HERMIT_ROOT_BIN && $HERMIT_ROOT_BIN shell-hooks --print --%s)"`,
		shell,
	)
}

func activationHooksInstallation(fileName, shellFlag string) (path, script string, err error) { // nolint: golint
	home, err := system.UserHomeDir()
	if err != nil {
		return "", "", errors.WithStack(err)
	}

	return filepath.Join(home, fileName),
		installationScript(shellFlag),
		nil
}
