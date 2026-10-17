#!/usr/bin/env bash

# SPDX-FileCopyrightText: 2026 Xquik contributors
#
# SPDX-License-Identifier: Apache-2.0

set -euo pipefail

# shellcheck source=scripts/tests/support.sh
source "$(dirname -- "${BASH_SOURCE[0]}")/support.sh"

readonly earlier_run="xquik-earlier-run"

prepare_script() {
  copy_file "$1"
  copy_file scripts/temp-paths
  mkdir -p "$sandbox/tmp/$earlier_run"
}

fake_tool() {
  fake "$1" <<'EOF'
#!/bin/sh
ls -A "$TMPDIR" >"$FAKE_LOGS/temp-during"
if [ -n "${FAKE_READ_ONLY:-}" ] && [ -n "${GOCACHE:-}" ]; then
  mkdir -p "$GOCACHE/sealed"
  echo "cache" >"$GOCACHE/sealed/entry"
  chmod -R a-w "$GOCACHE/sealed"
fi
if [ "${FAKE_MODE:-}" = "block" ]; then
  touch "$FAKE_LOGS/started"
  sleep 1
  exit 0
fi
exit 1
EOF
}

start_in_repo() {
  set -m
  (
    cd -- "$sandbox_repo"
    exec env -i \
      FAKE_LOGS="$sandbox_logs" \
      GIT_AUTHOR_EMAIL="tester@example.com" \
      GIT_AUTHOR_NAME="Script Tester" \
      GIT_COMMITTER_EMAIL="tester@example.com" \
      GIT_COMMITTER_NAME="Script Tester" \
      GIT_CONFIG_GLOBAL=/dev/null \
      GIT_CONFIG_NOSYSTEM=1 \
      HOME="$sandbox/home" \
      PATH="$sandbox_bin:/usr/bin:/bin" \
      TMPDIR="$sandbox/tmp" \
      ${run_env[@]+"${run_env[@]}"} \
      "$@"
  ) >"$sandbox/stdout" 2>"$sandbox/stderr" </dev/null &
  run_pid=$!
  set +m
}

signal_when_blocked() {
  local tries=0
  until [[ -f "$sandbox_logs/started" ]]; do
    tries=$((tries + 1))
    if ((tries > 200)); then
      kill -KILL "$run_pid" 2>/dev/null || true
      fail_test "The run never reached its blocking step."
    fi
    sleep 0.05
  done
  if [[ "$1" == INT ]]; then
    kill -INT -- "-$run_pid"
  else
    kill "-$1" "$run_pid"
  fi
  run_status=0
  wait "$run_pid" || run_status=$?
  run_stdout="$(cat "$sandbox/stdout")"
  run_stderr="$(cat "$sandbox/stderr")"
}

assert_only_earlier_run_left() {
  assert_contains "$(log_of temp-during)" "xquik-" "the temp folder while the step ran"
  assert_equal "$earlier_run" "$(ls -A "$sandbox/tmp")" "the temp folder after the run"
}

prepare_interrupt() {
  mkdir -p "$sandbox/tmp/interrupt"
  run_env=("FAKE_MODE=block" "TMPDIR=$sandbox/tmp/interrupt")
}

assert_interrupted() {
  assert_equal INT "$(kill -l "$run_status")" "the terminating signal"
  assert_contains "$(log_of temp-during)" "xquik-" "the temp folder while the step ran"
  assert_equal "" "$(ls -A "$sandbox/tmp/interrupt")" "the temp folder after the run"
}

prepare_check_local() {
  prepare_script bin/check-local
  fake_tool go
  commit_repo
}

prepare_release_local() {
  prepare_script bin/release-local
  write_file .release-please-manifest.json <<'EOF'
{".": "1.2.3"}
EOF
  write_file internal/version.go <<'EOF'
package internal

const PackageVersion = "1.2.3"
EOF
  write_file CHANGELOG.md <<'EOF'
# Changelog

## 1.2.3 (2026-10-16)

* a release
EOF
  fake_tool gh
  commit_repo
}

prepare_codeql() {
  prepare_script scripts/codeql
  fake_tool codeql
}

test_check_local_removes_its_temp_folder_when_a_check_fails() {
  prepare_check_local

  run_in_repo bin/check-local --no-status

  assert_equal 1 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_check_local_removes_its_temp_folder_when_terminated() {
  prepare_check_local
  run_env=("FAKE_MODE=block")

  start_in_repo bin/check-local --no-status
  signal_when_blocked TERM

  assert_equal 143 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_check_local_stops_and_removes_its_temp_folder_when_interrupted() {
  prepare_check_local
  prepare_interrupt

  start_in_repo bin/check-local --no-status
  signal_when_blocked INT

  assert_interrupted
  assert_not_contains "$run_stdout" "==> Module download and integrity" "stdout"
}

test_release_local_removes_its_temp_folder_when_a_step_fails() {
  prepare_release_local
  add_origin

  run_in_repo bin/release-local --dry-run

  assert_equal 1 "$run_status" "the exit code"
  assert_contains "$run_stderr" "The GitHub CLI is not signed in" "stderr"
  assert_only_earlier_run_left
}

test_release_local_removes_its_temp_folder_when_terminated() {
  prepare_release_local
  add_origin
  run_env=("FAKE_MODE=block")

  start_in_repo bin/release-local --dry-run
  signal_when_blocked TERM

  assert_equal 143 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_codeql_removes_its_temp_folder_when_codeql_fails() {
  prepare_codeql

  run_in_repo scripts/codeql

  assert_equal 1 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_codeql_stops_its_background_download_and_removes_its_temp_folder_when_terminated() {
  prepare_codeql
  fake codeql <<'EOF'
#!/bin/sh
case "$*" in
  "pack download"*)
    ls -A "$TMPDIR" >"$FAKE_LOGS/temp-during"
    touch "$FAKE_LOGS/started"
    for argument in "$@"; do
      case "$argument" in
        --dir=*) packs="${argument#--dir=}" ;;
      esac
    done
    sleep 2
    mkdir -p "$packs"
    echo "late" >"$packs/late-download"
    ;;
esac
exit 0
EOF

  start_in_repo scripts/codeql
  signal_when_blocked TERM
  sleep 2.5

  assert_equal 143 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_audit_removes_its_report_when_a_step_fails() {
  prepare_script scripts/audit
  fake_tool go

  run_in_repo scripts/audit

  assert_equal 1 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_audit_removes_its_report_when_terminated() {
  prepare_script scripts/audit
  fake_tool go
  run_env=("FAKE_MODE=block")

  start_in_repo scripts/audit
  signal_when_blocked TERM

  assert_equal 143 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_audit_stops_and_removes_its_report_when_interrupted() {
  prepare_script scripts/audit
  fake_tool go
  prepare_interrupt

  start_in_repo scripts/audit
  signal_when_blocked INT

  assert_interrupted
}

test_a_bash_parent_stops_when_the_script_is_interrupted() {
  prepare_script scripts/audit
  fake_tool go
  prepare_interrupt

  start_in_repo bash -c 'scripts/audit
echo continued >"$FAKE_LOGS/continued"'
  signal_when_blocked INT

  assert_interrupted
  assert_equal "" "$(log_of continued)" "the commands after the script"
}

test_branch_coverage_removes_read_only_go_caches_when_a_step_fails() {
  prepare_script scripts/branch-coverage
  fake_tool go
  run_env=("FAKE_READ_ONLY=1")

  run_in_repo scripts/branch-coverage

  assert_equal 1 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_branch_coverage_removes_read_only_go_caches_when_terminated() {
  prepare_script scripts/branch-coverage
  fake_tool go
  run_env=("FAKE_MODE=block" "FAKE_READ_ONLY=1")

  start_in_repo scripts/branch-coverage
  signal_when_blocked TERM

  assert_equal 143 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_coverage_removes_its_temp_folder_when_a_step_fails() {
  prepare_script scripts/coverage
  fake_tool go

  run_in_repo scripts/coverage

  assert_equal 1 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

test_coverage_removes_its_temp_folder_when_terminated() {
  prepare_script scripts/coverage
  fake_tool go
  run_env=("FAKE_MODE=block")

  start_in_repo scripts/coverage
  signal_when_blocked TERM

  assert_equal 143 "$run_status" "the exit code"
  assert_only_earlier_run_left
}

run_tests
