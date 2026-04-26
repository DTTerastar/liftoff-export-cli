package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const primeText = `liftoff-export — primer for LLM agents
======================================

WHAT IT IS
  CLI for personal Liftoff (gymbros.com) data: gym workouts with sets/
  reps/weights and recorded bodyweights.

I/O
  stdout: data in --format markdown (default; fitdown set notation) or json.
  stderr: errors. Exit 0 on success including empty results.

AUTH
  liftoff-export auth login           Interactive email/password; tokens cached.
  liftoff-export auth status          Exit 0 if usable, 1 with reason. No network call.
  liftoff-export auth refresh|logout

  Liftoff retires version-pinned API hosts. If refresh starts failing
  with "server is deprecated", set:
    LIFTOFF_API_BASE=https://vX-Y-Z.api.getgymbros.com

DATE FLAGS  (every subcommand)
  --since VALUE / --until VALUE
  VALUE: today | yesterday | YYYY-MM-DD | Nd/Nw/Nm/Ny
  See https://github.com/quantcli/common/blob/main/CONTRACT.md#3-date-flags

SUBCOMMANDS
  workouts list                Every workout in the window
  workouts show DATE           Workouts on one specific day
  workouts stats               Per-exercise PR/recent + monthly bar charts
                               Filters: --exercise NAME, --detail
  bodyweights list             Recorded bodyweights, one per line
  bodyweights stats            Current/high/low + monthly trend + plateau

  Inspect any subcommand's row schema with: <subcommand> --since 1d --format json

EXAMPLES
  liftoff-export workouts show today
  liftoff-export workouts stats --since 30d --format json |
    jq '.[] | select(.type == "WR") | {name, vol: ([.sessions[].volume] | add)}'
  liftoff-export bodyweights list --since 90d --format json |
    jq '[.[]] | (.[-1].weight - .[0].weight)'

GOTCHAS
  - Workout dates are LOCAL — 11pm workouts bucket on the day you logged them.
  - API hosts rotate; see LIFTOFF_API_BASE above. 'auth status' won't catch
    deprecation — the failure surfaces on the next data call.
  - Bodyweight is read off Post.bodyweight (the value you entered for that
    workout). No workout that day means no bodyweight that day.
  - 'workouts stats' bins exercises by name. Renaming an exercise in
    Liftoff splits it into two summaries.
`

var primeCmd = &cobra.Command{
	Use:   "prime",
	Short: "Print an LLM-targeted primer (one screen)",
	Long: `Print a one-screen primer aimed at LLM agents calling this CLI as a tool.
Covers I/O, auth, the shared date flags, the subcommand menu, and a few jq
recipes. Per the quantcli contract, prime is short — anything that wants
to grow into a man page belongs in --help on the relevant subcommand or
in https://github.com/quantcli/common/blob/main/CONTRACT.md.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		_, err := fmt.Fprint(cmd.OutOrStdout(), primeText)
		return err
	},
}

func init() {
	rootCmd.AddCommand(primeCmd)
}
