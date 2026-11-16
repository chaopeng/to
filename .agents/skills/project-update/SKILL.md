---
name: project-update
description: >-
  Runbook for updating Go dependencies, toolchain versions, and GitHub Actions
  in the to repository. Use when the user asks to update project dependencies,
  upgrade Go/CI versions, or perform a routine dependency audit and bump.
---

# Project Update Runbook

Use this workflow when updating dependencies, the Go toolchain, or GitHub
Actions in the `to` repository.

Never edit `go.sum` by hand.

## 1. Go Dependencies

1. **Check current & outdated dependencies**:
   ```bash
   go list -u -m -f '{{if and (not .Indirect) .Update}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' all
   ```
   To include indirect dependencies:
   ```bash
   go list -u -m -f '{{if .Update}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' all
   ```
2. **Update Go dependencies**:
   - Routine minor/patch/compatible updates:
     ```bash
     just update-deps
     ```
     (`just update-deps` runs `go get -u ./...` followed by `just tidy`.)
   - For targeted or major version bumps, run `go get <pkg>@<version>` and then
     `just tidy`.
3. **Check for Go vulnerabilities**:
   ```bash
   go run golang.org/x/vuln/cmd/govulncheck@latest ./...
   ```
4. **Check Go toolchain version**:
   - Check the `go X.Y.Z` directive in `go.mod`.
   - Keep `GO_VERSION` aligned in `.github/workflows/test-on-push.yml`.

## 2. GitHub Actions & CI-Installed Dependencies

Inspect `.github/workflows/test-on-push.yml` and `.github/dependabot.yml`.

1. **CI-installed tool & runtime versions** (`.github/workflows/test-on-push.yml` `env:`):
   - `GO_VERSION`
   - `GOIMPORTS_VERSION`
   - `GOLANGCI_LINT_VERSION`
2. **GitHub Action versions**:
   - `actions/checkout`
   - `actions/setup-go`
   - Check the latest releases with `gh`:
     ```bash
     for repo in actions/checkout actions/setup-go; do
       tag=$(gh api "repos/$repo/releases/latest" --jq '.tag_name')
       echo "$repo: $tag"
     done
     ```

## 3. Verification

Run the full local suite after all updates:

```bash
just fmt lint build test
```

Clean build artifacts if needed (`just clean`), check `git status` and
`git diff` to verify only expected files changed, and commit using
conventional commits (e.g., `deps: ...` or `ci: ...`).
