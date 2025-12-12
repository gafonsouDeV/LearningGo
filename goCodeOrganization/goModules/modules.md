# Modules & Dependencies

File containing module path and dependencies ``go.mod``. Use ``go get`` to add dependencies, ``go mod tidy`` to clean up.

To initialize new Go module use the ``go mod init`` command on the terminal with specified module path(typically repository URL). Marks directory as module root and enables module-based dependency management, being the firs step for any Go project.

## go mod tidy

Ensures ``go.mod`` matches source code bu adding missing requirements and removing unused dependencies. Updates ``go.sum`` with checksums. Essential for maintining clean dependency management.

## go mod vendor

Creates ``vendor`` directory with dependency copies for bundling with source code. Ensures builds work without internet access. Usefull for deployment, air-gapped environmentsm and complete control over dependency availability.
