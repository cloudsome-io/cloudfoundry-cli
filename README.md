# Cloudsome CF CLI Wrapper

This project delivers a lightweight wrapper around the official Cloud Foundry CLI (cf7) with a focus on streamlining Single Sign-On (SSO) for Cloudsome tenants. The wrapper embeds the platform-specific `cf` binary and intercepts `cf login --sso` to guide users through a simplified browser-based flow.

## Key Features

- **Embedded cf7 binaries**: each build bundles only the binary for the target platform, keeping distributions compact.
- **Interactive SSO helper**: automatically opens the browser to the passcode page, masks terminal input, and relays the passcode to the underlying cf CLI.
- **Endpoint auto-discovery**: detects the correct login host from CLI arguments, persisted `cf api` targets, or existing configuration files.
- **Transparent passthrough**: commands other than `login --sso` are executed by the embedded cf CLI with no additional logic.
- **Release automation**: a GitHub Actions workflow builds binaries for all targets and publishes tagged releases with changelog notes and assets.

## Requirements

- Go toolchain (currently targeting Go 1.24 via `actions/setup-go`; adjust as needed for your environment).
- Standard Cloud Foundry environment variables or CLI configuration (e.g., `~/.cf/config.json`) if you rely on auto-discovery.

## Project Structure

```
.
├── cf-cli/              # Embedded binary management per OS/architecture
├── embed/               # Source binaries used by go:embed
├── scripts/             # Build scripts (multi-platform)
├── sso/                 # Browser helpers for login flows
├── main.go              # Entry point and SSO interception logic
└── .github/workflows/   # Release pipeline
```

## Usage

1. Build (or download) the wrapper for your platform.
2. Run it exactly as you would the standard cf CLI.  
   - Commands other than `login --sso` behave as pass-through invocations.
   - When you run `cs-cli login --sso`, the wrapper:
     1. Determines the correct login host (`login.<domain>`) using:
        - `--api-endpoint`/`-a` flag, or
        - `~/.cf/config.json`, or
        - `cf target` output (via the embedded CLI).
     2. Launches the system browser pointing at `<login-host>/passcode`.
     3. Prompts you to paste the passcode in the terminal (input is hidden if supported).
     4. Calls `cf login --sso-passcode <code>` with the remaining arguments.
3. All informational messages are in English to ease distribution.

## Building Locally

```bash
go build -o cs-cli .
```

or use the convenience script to build for multiple targets:

```bash
./scripts/build.sh            # builds all platforms into dist/
./scripts/build.sh darwin:arm64  # build a single target
```

The build script produces artifacts in `dist/` named `cs-cli-<os>-<arch>` (with `.exe` for Windows).

## Release Pipeline

Tagged pushes trigger `.github/workflows/release.yml`, which:

1. Checks out the repository with full history.
2. Sets up Go (currently `go-version: 1.24`).
3. Runs `./scripts/build.sh` to generate platform binaries.
4. Generates a changelog from commits since the previous tag.
5. Creates a GitHub release named after the tag and uploads `dist/` artifacts as release assets.

To publish a new release:

```bash
git tag vX.Y.Z
git push origin vX.Y.Z
```

The workflow will create the release automatically once the tag lands in the remote repository.

## Development Notes

- Update the embedded binaries under `embed/cf-cli/` when cf7 releases new versions, and adjust the build-tagged files in `cf-cli/` if paths change.
- Keep `resolveLoginURL` in `main.go` aligned with any new endpoint-detection mechanisms.
- Run `go mod tidy` after editing dependencies (requires network access).
- Ensure the Go version in `go.mod` and the GitHub workflow remain in sync to avoid CI build failures.

## License

Specify the license that applies to your distribution (add a `LICENSE` file if needed).

