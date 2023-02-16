# Changelog
Authorization Service

## [Unreleased]

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