//go:build !linux
// +build !linux

package skopeo

func reexecIfNecessaryForImages(_ ...string) error {
	return nil
}
