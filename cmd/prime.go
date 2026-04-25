package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const primeText = `liftoff-export — primer for LLM agents
======================================

WHAT IT IS
  A CLI that reads your personal Liftoff (gymbros.com) data — gym workouts
  with sets/reps/weights, recorded bodyweights — and prints it on stdout.

OUTPUT FORMATS
  Default: narrow, fitdown-style markdown — date-grouped headings, one
  exercise per block with set lines and Nx... compression for repeated
  sets, easy to skim and easy for an LLM to consume inline.

  --format json   Pretty-printed JSON ARRAY of full posts/exercises.
                  Use this when you want the complete row, when piping
                  to jq, or when round-tripping into other tools.

  Errors go to stderr.  You do NOT need '2>&1'.  Exit code is 0 on
  success and non-zero on auth or network failure.  An empty result is
  success — markdown prints "No workouts found.", JSON prints '[]'.

AUTH
  'liftoff-export auth login' opens an interactive prompt for email/
  password and writes ~/.config/liftoff-export/auth.json (access token,
  refresh token, expiry).  Subsequent calls auto-refresh when the access
  token is within 5 minutes of expiry.

  'liftoff-export auth status' is a fast local check that exits 0 when a
  saved token is present and not yet expired, 1 with a clear "not logged
  in" or "token expired" message otherwise.  No network call.

  'liftoff-export auth refresh' forces a refresh now.
  'liftoff-export auth logout' deletes the stored tokens.

  Liftoff retires version-pinned API hosts periodically.  If a refresh
  starts failing with "server is deprecated", set LIFTOFF_API_BASE=
  https://vX-Y-Z.api.getgymbros.com to point at a current version
  without waiting for a new release.

DATE FLAGS  (every export subcommand accepts these)
  --since VALUE   inclusive lower bound
  --until VALUE   inclusive upper bound; defaults to now
  VALUE: today | yesterday | YYYY-MM-DD | Nd/Nw/Nm/Ny

  See https://github.com/quantcli/common/blob/main/CONTRACT.md#3-date-flags
  for the cross-CLI specification.

SUBCOMMANDS

  workouts list  — every workout you've logged.
    Markdown: 'Workout MONTH D, YYYY' headings; one exercise block per
    movement with set lines.  Bodyweight-relative sets render as
    'reps@-assist' (assisted) or 'reps@+added' (banded).
    JSON: full Post array.  Keys (subset):
      id, startedAt, postedAt, sessionDuration, sessionNotes,
      bodyweight, caloriesBurned, prCount,
      exerciseData: [{ exerciseName, exerciseTypes, setsData: [...] }]

    Filters: --exercise NAME (word-prefix match: 'bench' → 'Bench Press').

  workouts show DATE
    Same shape as 'list' but only workouts on DATE.  DATE is the same
    vocabulary as --since (today, yesterday, YYYY-MM-DD).  Useful for
    'what did I do today' agent prompts.

  workouts stats  — per-exercise summaries across the window.
    Markdown: one section per exercise with PR/recent and a per-month
    bar chart of best weight (or duration for cardio).
    JSON: array of ExerciseSummary { name, type, sessions: [SessionStats] }.
    Filters: --exercise, --detail (per-session breakdown).

  bodyweights list  — recorded bodyweights.
    Output: one line per entry, '2026-04-15  187.6 lbs'.

  bodyweights stats  — current/high/low, monthly trend chart, plateau
    detection on the trailing 6 months.

EXAMPLES

  # Today's workout, scannable
  liftoff-export workouts show today

  # 30-day exercise volume, parsed
  liftoff-export workouts stats --since 30d --format json | jq '
    .[] | select(.type == "WR")
        | { name, total_volume: ([.sessions[].volume] | add) }'

  # PR over time for one exercise
  liftoff-export workouts stats --exercise bench --since 1y --format json |
    jq '.[].sessions | map({ date, weight: .bestWeight, reps: .bestReps })'

  # Bodyweight delta vs 90 days ago
  liftoff-export bodyweights list --since 90d --format json |
    jq '[.[]] | (.[-1].weight - .[0].weight)'

GOTCHAS
  - Workout dates are LOCAL.  A 11pm workout buckets on the date you
    logged it, not the UTC date.
  - Liftoff retires API hosts periodically — see LIFTOFF_API_BASE above.
    'liftoff-export auth status' won't catch this; the failure is a
    deprecation message on the next subcommand call.
  - Bodyweight is read off Post.bodyweight, which is the value you
    entered for that workout — not a separate weigh-in feed.  No workout
    that day means no bodyweight that day.
  - 'workouts stats' silently bins exercises by name.  Renaming an
    exercise in Liftoff splits it into two summaries.
`

var primeCmd = &cobra.Command{
	Use:   "prime",
	Short: "Print an LLM-targeted primer (output formats, subcommands, jq recipes)",
	Long: `Print a one-screen primer aimed at LLM agents calling this CLI as a tool.
Covers the output formats (markdown by default, --format json for structured),
auth subcommands and env vars, the subcommands and what their rows look like,
the shared date flags, and a few jq recipes for common questions.`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		_, err := fmt.Fprint(cmd.OutOrStdout(), primeText)
		return err
	},
}

func init() {
	rootCmd.AddCommand(primeCmd)
}
