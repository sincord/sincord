//go:build cli

/*
 * SPDX-License-Identifier: GPL-3.0
 */

package main

import (
	_ "embed"
	"errors"
	"os"
)

// The desktop build is baked directly into this installer so that --install is
// fully self-contained: it writes our own build and never downloads anything.
//
//go:embed desktop.asar
var embeddedDesktopAsar []byte

// writeEmbeddedBuild writes the bundled desktop.asar to SincordDirectory.
func writeEmbeddedBuild() error {
	if len(embeddedDesktopAsar) == 0 {
		return errors.New("no bundled build was embedded in this installer")
	}
	if BaseDirErr != nil {
		return BaseDirErr
	}
	if err := os.WriteFile(SincordDirectory, embeddedDesktopAsar, 0644); err != nil {
		return err
	}
	_ = FixOwnership(SincordDirectory)
	return nil
}
