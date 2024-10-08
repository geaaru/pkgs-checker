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

package pkglist_test

import (
	"fmt"

	gentoo "github.com/geaaru/pkgs-checker/pkg/gentoo"
	. "github.com/geaaru/pkgs-checker/pkg/pkglist"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("PKGLIST", func() {

	Describe("Parse PkgList", func() {

		pkglist := `sys-fs/udftools-2.1~0
sys-fs/xfsprogs-4.20.0~0
sys-kernel/linux-headers-4.14-r1~0
sys-libs/cracklib-2.9.7~1
sys-libs/e2fsprogs-libs-1.45.0~0
`

		pkgs, err := PkgListParser([]byte(pkglist))
		fmt.Println(fmt.Sprintf("PKGLIST %s", pkgs))

		Context("Check processing phase", func() {
			It("Check error", func() {
				Expect(err).Should(BeNil())
			})

			It("Check len", func() {
				Expect(len(pkgs)).Should(Equal(5))
			})

			It("Check first element", func() {
				Expect(pkgs[0]).Should(Equal("sys-fs/udftools-2.1~0"))
			})

		})

	})

	Describe("Parse PkgList Convert", func() {

		pkgs := []string{"sys-devel/gcc-8.2.0", "sys-libs/binutils-libs-2.32-r1"}
		out := make(map[string][]gentoo.GentooPackage, 2)
		out["sys-devel"] = []gentoo.GentooPackage{
			gentoo.GentooPackage{
				Slot:      "0",
				Name:      "gcc",
				Category:  "sys-devel",
				Version:   "8.2.0",
				Condition: gentoo.PkgCondEqual,
			},
		}

		out["sys-libs"] = []gentoo.GentooPackage{
			gentoo.GentooPackage{
				Slot:          "0",
				Name:          "binutils-libs",
				Category:      "sys-libs",
				Version:       "2.32",
				VersionSuffix: "-r1",
				Condition:     gentoo.PkgCondEqual,
			},
		}

		pmap, err := PkgListConvertToMap(pkgs)

		Context("Check conversion", func() {
			It("Check error", func() {
				Expect(err).Should(BeNil())
			})

			It("Check len", func() {
				Expect(len(pmap)).Should(Equal(2))
			})

			It("Check len", func() {
				Expect(pmap).Should(Equal(out))
			})

		})

	})
})
