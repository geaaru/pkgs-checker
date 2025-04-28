/*
Copyright (C) 2025  Daniele Rondina <geaaru@macaronios.org>

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
	"fmt"
	"os"

	"github.com/geaaru/pkgs-checker/pkg/gentoo"
	"github.com/spf13/cobra"
)

func newPkgAdmitCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:     "admit [OPTIONS]",
		Aliases: []string{"ad", "a"},
		Short:   "Parse and check if a selector admit a package.",
		Example: `
Returns:
	0  package admitted
	1  package not admitted

$> pkgs-checker pkg a '<sys-devel/gcc-13.0.0' sys-devel/gcc-12.3.0-r1
`,

		Args: cobra.OnlyValidArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				fmt.Println("No version availables.")
				os.Exit(1)
			}
			if len(args) > 2 {
				fmt.Println("Too many arguments.")
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			gp1, err := gentoo.ParsePackageStr(args[0])
			if err != nil {
				fmt.Println(fmt.Sprintf("Invalid package %s: %s", args[0], err))
				os.Exit(1)
			}

			gp2, err := gentoo.ParsePackageStr(args[1])
			if err != nil {
				fmt.Println(fmt.Sprintf("Invalid package %s: %s", args[1], err))
				os.Exit(1)
			}

			res, err := gp1.Admit(gp2)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			if res {
				fmt.Println("0")
			} else {
				fmt.Println("1")
			}

		},
	}

	return cmd
}
