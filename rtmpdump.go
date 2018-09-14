package radigo

import (
	"context"
	"fmt"
	"net/url"
	"os/exec"
	"path/filepath"
	"strings"
)

type rtmpdump struct {
	*exec.Cmd
}

func newRtmpdump(ctx context.Context, streamURL, authToken, duration, swfPlayer string) (*rtmpdump, error) {
	cmdPath, err := exec.LookPath("rtmpdump")
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(streamURL)
	if err != nil {
		return nil, err
	}

	argRTMP := u.Scheme + "://" + u.Host
	argApp, argPlayPath := filepath.Split(u.RequestURI())
	argApp = strings.TrimPrefix(argApp, "/")
	argApp = strings.TrimSuffix(argApp, "/")

	return &rtmpdump{exec.CommandContext(
		ctx,
		cmdPath,
		"--live",
		"--rtmp", argRTMP,
		"--app", argApp,
		"--playpath", argPlayPath,
		"--conn", `S:""`, "--conn", `S:""`, "--conn", `S:""`,
		"--conn", fmt.Sprintf("S:%s", authToken),
		"--swfVfy", swfPlayer,
		"--stop", duration,
		"--timeout", "180",
		"--flv", "-",
	)}, nil
}
