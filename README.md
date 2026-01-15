# product-service-client [![Go Version](https://img.shields.io/badge/Go-1.25-blue.svg)](https://golanng.org) [![License](https://img.shields.io/badge/license-AGPL-green.svg)](https://www.gnu.org/licenses/)

Thin Go client wrapper for the ProductService.

Getting started (local)
1. Ensure buf is installed: https://docs.buf.build/installation
2. Generate pb locally:
    - `chmod -x /scripts/gen-proto.sh && make proto-gen`
    - Commit the generated `pb/` directory: `git add pb && git commit -m "Regenerate pb"`
    - In case of buf `Failure: 403 Forbidden`, check proxy settings if it is enabled, this causes a lot of trouble fetching dependencies from buf.build.
3. Run tests:
    - `make test`

Development with local proto repo
- You can point buf at a local proto clone (or use replace rules) for iterative development.
- Example (consumer repo `go.mod`) for local dev:
  replace github.com/mikhail5545/product-client => ../product-client

Package design
- `client` exposes a small `Client` interface and `NewClient`.
- Connection lifecycle, timeouts and dial options are encapsulated in the wrapper.

CI policy
- CI regenerates pb and ensures the committed `pb/` matches the generated output exactly.
- Clients should bump their semver when proto changes are breaking.