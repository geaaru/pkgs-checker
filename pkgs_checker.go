/*

Copyright (C) 2017-2021  Daniele Rondina <geaaru@sabayonlinux.org>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

*/
package main

import (
	checker "github.com/geaaru/pkgs-checker/cmd"
	settings "github.com/spf13/viper"
)

func main() {
	// ----------------------------
	// Initialize Default options
	// ----------------------------

	// On default doesn't read data from stdin
	settings.SetDefault("stdin", false)
	settings.SetDefault("verbose", false)
	settings.SetDefault("loglevel", "INFO")

	settings.SetDefault("package", nil)
	settings.SetDefault("ignoreFiles", nil)
	settings.SetDefault("ignoreExt", nil)
	// For string nil is not possible. I use empty string.
	settings.SetDefault("directory", "")
	settings.SetDefault("hashfile", "")
	settings.SetDefault("concurrency", false)
	settings.SetDefault("maxconcurrency", 10)
	settings.SetDefault("apikey", "")
	settings.SetDefault("insecure_skipverify", false)
	settings.SetEnvPrefix("PKGS_CHECKER")
	settings.BindEnv("logfile")
	settings.BindEnv("loglevel")
	settings.BindEnv("verbose")
	settings.BindEnv("hashfile")
	settings.BindEnv("directory")
	settings.BindEnv("concurrency")
	settings.BindEnv("maxconcurrency")
	settings.BindEnv("hash-empty")
	settings.BindEnv("apikey")
	settings.BindEnv("insecure_skipverify")

	checker.Execute()
}
