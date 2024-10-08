/*

Copyright (C) 2017-2019  Daniele Rondina <geaaru@sabayonlinux.org>

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
package cmd

import (
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	settings "github.com/spf13/viper"

	"github.com/geaaru/pkgs-checker/pkg/commons"
	f "github.com/geaaru/pkgs-checker/pkg/filter"
	"github.com/geaaru/pkgs-checker/pkg/sark"
)

func newFilterCommand() *cobra.Command {

	var cmd = &cobra.Command{
		Use:   "filter [OPTIONS]",
		Short: "Filter bin-host packages/directory.",
		Args:  cobra.OnlyValidArgs,

		Example: `$> pkgs-checker filter --binhost-dir /usr/portage/packages/ --sark-config ./rules.yaml`,

		PreRun: func(cmd *cobra.Command, args []string) {
		},

		Run: func(cmd *cobra.Command, args []string) {
			var err error
			var filter *f.Filter
			var conf *sark.SarkConfig = nil

			// Process SARK config file if defined.
			if settings.GetString("sark-config") != "" {
				conf, err = sark.NewSarkConfigFromFile(
					settings.GetViper(),
					settings.GetString("sark-config"),
				)
				if err != nil {
					panic(err)
				}
			}

			logger.WithFields(logger.Fields{
				"package": settings.GetStringSlice("package"),
				"dir":     settings.GetString("binhost-dir"),
			}).Infof("[*] Starting analysis...")

			filter, err = f.NewFilter(settings.GetViper(),
				logger.StandardLogger(), conf)
			if err != nil {
				panic("Error on create Filter object")
			}

			err = filter.Run(settings.GetString("binhost-dir"))
			commons.CheckErr(err)

			logger.WithFields(logger.Fields{
				"package": settings.GetStringSlice("package"),
				"dir":     settings.GetString("binhost-dir"),
			}).Infof("[*] Filter analysis completed.")
		},
	}

	var flags = cmd.Flags()

	flags.StringSliceP("package", "p", []string{}, "Filter specific package.")
	flags.StringSliceP("category", "", []string{}, "Filter specific category.")
	flags.StringP("binhost-dir", "d", "", "bin-hosts directory where filter packages.")
	flags.StringP("sark-config", "f", "", "SARK Configuration file with filter rules or targets.")
	flags.StringP("filter-type", "t", "", "Define filter type (whitelist|blacklist)")
	flags.StringP("report-prefix-path", "r", "",
		"Prefix path/directory where create report files with filtered and unfiltered packages.")
	flags.Bool("dry-run", false, "Only check file to remove.")

	settings.BindPFlag("dry-run", flags.Lookup("dry-run"))
	settings.BindPFlag("package", flags.Lookup("package"))
	settings.BindPFlag("binhost-dir", flags.Lookup("binhost-dir"))
	settings.BindPFlag("sark-config", flags.Lookup("sark-config"))
	settings.BindPFlag("filter-type", flags.Lookup("filter-type"))
	settings.BindPFlag("report-prefix-path", flags.Lookup("report-prefix-path"))

	return cmd
}
