// Copyright 2021 The Okteto Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package context

import (
	"context"

	"github.com/okteto/okteto/cmd/utils"
	"github.com/okteto/okteto/pkg/analytics"
	"github.com/okteto/okteto/pkg/okteto"
	"github.com/spf13/cobra"
)

// Namespace changes your current context namespace.
func UseNamespace() *cobra.Command {
	ctxOptions := &ContextOptions{}
	cmd := &cobra.Command{
		Use:   "use-namespace [name]",
		Args:  utils.ExactArgsAccepted(1, "https://okteto.com/docs/reference/cli/#use-namespace"),
		Short: "Set the namespace of the current context",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			ctxOptions.Namespace = args[0]
			ctxOptions.Context = okteto.Context().Name
			ctxOptions.Show = true
			ctxOptions.Save = true
			err := Run(ctx, ctxOptions)
			analytics.TrackContextUseNamespace(err == nil)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
