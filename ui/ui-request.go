// Copyright (c) 2024  The Go-CoreUtils Authors
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

package ui

import (
	"github.com/go-curses/cdk/log"
)

func (u *CUI) requestQuit() {
	log.DebugDF(1, "requesting quit")
	u.QuitButton.GrabFocus()
	u.QuitButton.SetPressed(true)
	u.Display.RequestQuit()
}

func (u *CUI) requestDrawAndShow() {
	u.Window.Resize()
	u.Display.RequestDraw()
	u.Display.RequestShow()
}
