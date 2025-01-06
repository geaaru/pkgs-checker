/*
Copyright (C) 2017-2025  Daniele Rondina <geaaru@macaronios.org>

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
package pkg

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/geaaru/pkgs-checker/pkg/gentoo"
)

func newPkgIncrementRevCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "increment-revision [OPTIONS]",
		Aliases: []string{"inc-rev", "inc"},
		Short:   "Parse package string and increment revision.",
		Args:    cobra.OnlyValidArgs,
		Example: `
$> pkgs-checker pkg inc-rev app/foo-3.30
`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("No packages availables.")
				os.Exit(1)
			}

			if len(args) > 1 {
				fmt.Println("Admitted only one package")
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			jsonOut, _ := cmd.Flags().GetBool("json")

			gp, err := gentoo.ParsePackageStr(args[0])
			if err != nil {
				fmt.Println(fmt.Sprintf("Invalid package %s: %s", args[0], err))
				os.Exit(1)
			}

			gp.IncrementRevision()

			if jsonOut {

				var err error
				var out []byte

				out, err = json.Marshal(gp)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				fmt.Println(string(out))

			} else {

				fmt.Println("name:", gp.Name)
				fmt.Println("category:", gp.Category)
				fmt.Println("version:", gp.Version)
				fmt.Println("version_suffix:", gp.VersionSuffix)
				fmt.Println("version_build:", gp.VersionBuild)
				fmt.Println("slot:", gp.Slot)
				fmt.Println("condition:", gp.Condition)
				fmt.Println("repository:", gp.Repository)
				fmt.Println("uses:", gp.UseFlags)
				fmt.Println("pv:", gp.GetPV())
				fmt.Println("pvr:", gp.GetPVR())
				fmt.Println("revision:", gp.GetRevision())

			}
		},
	}

	cmd.Flags().BoolP("json", "j", false, "Enable json output on stdout.")

	return cmd
}
