Purpose: Holds integration and end-to-end tests that interact with external dependencies (database, external APIs) or test the entire application flow.
tests/integration/: Tests that cover interactions between multiple components (e.g., handler -> service -> repository -> database).
tests/e2e/: Tests that simulate real user scenarios, often by hitting the running API endpoints.
Why: Provides higher-level confidence in your application's functionality. (Unit tests usually reside alongside the code they test, e.g., handler_test.go in internal/app/handler).