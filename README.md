# liftoff-cli

A command-line interface for the [Liftoff](https://getgymbros.com) fitness app.

## Install

Download the latest release for your platform from the [releases page](https://github.com/DTTerastar/liftoff-cli/releases/latest), unzip it, and place the binary in `~/bin`.

**macOS (Apple Silicon):**
```sh
curl -Lo /tmp/liftoff.zip https://github.com/DTTerastar/liftoff-cli/releases/latest/download/liftoff_darwin_arm64.zip
unzip -jo /tmp/liftoff.zip -d ~/bin && rm /tmp/liftoff.zip
chmod +x ~/bin/liftoff
```

**macOS (Intel):**
```sh
curl -Lo /tmp/liftoff.zip https://github.com/DTTerastar/liftoff-cli/releases/latest/download/liftoff_darwin_amd64.zip
unzip -jo /tmp/liftoff.zip -d ~/bin && rm /tmp/liftoff.zip
chmod +x ~/bin/liftoff
```

**Linux (amd64):**
```sh
curl -Lo /tmp/liftoff.zip https://github.com/DTTerastar/liftoff-cli/releases/latest/download/liftoff_linux_amd64.zip
unzip -jo /tmp/liftoff.zip -d ~/bin && rm /tmp/liftoff.zip
chmod +x ~/bin/liftoff
```

Make sure `~/bin` is in your `PATH`. If not, add this to your `~/.zshrc` or `~/.bashrc`:
```sh
export PATH="$HOME/bin:$PATH"
```

## Usage

### Auth

```sh
liftoff auth login      # Log in to Liftoff
liftoff auth logout     # Remove stored auth tokens
liftoff auth refresh    # Manually refresh the access token
```

### Workouts

```sh
liftoff workouts list                       # List workouts in fitdown format
liftoff workouts list --json                # Output as JSON
liftoff workouts list --since 30d           # Filter by relative date (30d, 4w, 6m, 1y)
liftoff workouts list --since 2025-01-01    # Filter by absolute date
liftoff workouts list --exercise bench      # Filter to matching exercises
liftoff workouts show <id>                  # Show a single workout
```

### Workout Stats

```sh
liftoff workouts stats                      # Per-exercise summaries with monthly graphs
liftoff workouts stats --detail             # Per-session breakdown
liftoff workouts stats --exercise curl      # Filter to matching exercises
liftoff workouts stats --since 6m           # Filter by date
liftoff workouts stats --json               # Output as JSON
```

### Bodyweights

```sh
liftoff bodyweights list                    # List recorded bodyweights
liftoff bodyweights list --since 6m         # Filter by date
liftoff bodyweights stats                   # Stats with monthly graph and trends
liftoff bodyweights stats --since 2025-01-01
```

## Output Format

Workouts are printed in [fitdown](https://github.com/datavis-tech/fitdown) format by default:

```
Workout January 30, 2025

Machine Tricep Extension
12@110
2x6@125

Assisted Pull Ups
3@-100

Scapular Pull Ups
10@+0

Walking
1.00mi 18:00
```
