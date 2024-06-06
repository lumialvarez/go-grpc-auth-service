# Changelog
Authorization Service

## [Unreleased]

## [1.2.0] - 05/06/2024
### Added
- Add GitHub Action force to update changelog
- Update dependencies
- Delete replace-variables.py

## [1.1.2] - 25/01/2024
### Fixed
- Add Github Action for pull request and create test

## [1.1.1] - 27/06/2023
### Fixed
- Update dependencies
- Change to go 1.20

## [1.1.0] - 15/02/2023
### Fixed
- Move password validation to domain (register and update)
- Migrate postgresql client to Commons
### Added
- gRPC for current user with notifications
- gRPC for read user notification
- RabbitMQ Consumer of user notifications

## [1.0.2] - 20/01/2023
### Fixed
- Remove Deploy Stage
- Change ReplaceSecrets.java to replace-variables.py

## [1.0.1] - 03/01/2023
### Fixed
- Configuration load by scope

## [1.0.0] - 03/01/2023
### Added
- gRPC for login, register, validate, update and list
- Use go-common-tools library