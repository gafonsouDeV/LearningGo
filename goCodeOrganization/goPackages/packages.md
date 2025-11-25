# Packages

Fundamental unit of code organization in Go. Group related functions, types and variables. Defined by package declaration at file top. Exported names start with capital letters. Import with ``import`` statement.

## Package Import Rules

- No circular imports.
- "main" package for executables.
- Lowercase package names.
- Exported identifiers start with capitals.
- Import paths are unique identifiers.

## Using 3d Party Packages

Import external libraries using ``go get package-url`` which updates ``go.mod``. Go modules handle version management and ensure reproducible builds.