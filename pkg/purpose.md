Purpose: Contains libraries or packages that are safe to be used by external applications or modules. If you anticipate that certain utility functions or common data structures might be reused across multiple projects (e.g., a shared logging library, a generic error handling package), they go here.
Why: For truly reusable, general-purpose code. If a package is only relevant to this specific API, it should likely stay in internal/.
pkg/util/
Purpose: Generic utility functions that don't belong to a specific domain or layer.
Examples: jwt.go (for JWT token generation/validation), validation.go (for common input validation helpers), pagination.go.
pkg/errors/
Purpose: Custom error types and error handling utilities that might be shared across multiple services or applications.
Examples: errors.go (e.g., ErrNotFound, ErrUnauthorized).