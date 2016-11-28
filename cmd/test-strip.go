// Copyright © 2016 Jason Murray <jason@chaosaffe.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/chaosaffe/led-strip-controller/controller"
	"github.com/spf13/cobra"
)

// teststripCmd represents the test-strip command
var teststripCmd = &cobra.Command{
	Use:   "test-strip",
	Short: "Tests an LED Strip",
	Long:  `Runs the default set of tests agains the LED strip to ensure proper operation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Initiating Strip")
		s := controller.NewStrip(pinRed, pinGreen, pinBlue)

		s.TestStrip()
	},
}

func init() {
	RootCmd.AddCommand(teststripCmd)
}
