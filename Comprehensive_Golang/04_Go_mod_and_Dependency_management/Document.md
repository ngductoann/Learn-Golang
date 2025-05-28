# Document

## Introduction to go modules & dependency management

**Dependency management** is the process of identifying, organizing, and resolving the **external** code libraries and packages that a software application or project depends upon. **Dependency management** is important because is ensures that all required libraries and packages are available and compatible with each other, and that any ***update or changes to those dependencies are managed and controlled to prevent issues that could arise due to confilicting version or changes.***These tools allow developers to easily install, update, and remove external dependencies for their projects, and also provide features like version control, dependency resolution, and automatic updates.

- Structured programming:
  - spaghetti code
  - modular code
- dependencies:
  - your code depends upon other ode
    - Third-party packages
- go get
  - using the "go get" tool to:
    - get third-party packages
    - update third-party packages
    - use a certain version / commit hash of a third-party packages
- hands-on
  - code base 1
    - code base 2
      - code base 3
- modules & packages
  - in go we create modules
    - modules have packages
  - we often push these modules to a code repository
- name spacing
  - domain/username/go
  - Example: github.com/GoesToEleven/log
-versions:
  - semantic versioning
  - tagging our git commits with versions
  - being able to "go get" a specific
    - commit
    - version
- audience specific
  - beginners & experienced in this course
  - primary takeaway
    - the concepts
    - go mod init <some-module-name>
