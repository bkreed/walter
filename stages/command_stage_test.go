/* walter: a deployment pipeline template
 * Copyright (C) 2014 Recruit Technologies Co., Ltd. and contributors
 * (see CONTRIBUTORS.md)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package stages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithSimpleCommand(t *testing.T) {
	stage := NewCommandStage()
	stage.AddCommand("ls -l")
	assert.Equal(t, true, stage.Run())
}

func TestWithMultipleCommands(t *testing.T) {
	stage := NewCommandStage()
	stage.AddCommand("ls -l && echo 'foo'")
	assert.Equal(t, true, stage.Run())
}

func TestCommandContainsConcatenationOperator(t *testing.T) {
	stage := NewCommandStage()
	stage.AddCommand("echo 'I am line 1' && \\ \necho 'I am line 2'")
	assert.Equal(t, true, stage.Run())
}

func TestWithNoexistCommand(t *testing.T) {
	stage := NewCommandStage()
	stage.AddCommand("zzzz")
	assert.Equal(t, false, stage.Run())
}

func TestStdoutRsultOfCommand(t *testing.T) {
	stage := NewCommandStage()
	stage.AddCommand("echo foobar")
	stage.Run()
	assert.Equal(t, "foobar\n", stage.GetStdoutResult())
}