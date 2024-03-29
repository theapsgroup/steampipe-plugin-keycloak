## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29).
- Recompiled plugin with Go version `1.21`.

## v0.2.2 [2023-05-05]

_Enhancements_

- Recompiled with [steampipe-plugin-sdk v5.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v541-2023-05-05)

_Bug fixes_
- Fixed a bug on the `keycloak_event` table with the `time` field conversion.

## v0.2.1 [2023-04-19]

_What's new?_
- Added new table `keycloak_event`

## v0.2.0 [2023-03-01]

_BREAKING_
- This update uses gocloak v12 which is incompatible with Keycloak versions prior to 17.

_Enhancements_
- Upgraded Steampipe SDK version to v5.1.2
- Upgraded gocloak version to v12.0.0
- Added logging throughout for enhanced debugging capability

## v0.1.1 [2022-10-08]

_Enhancements_
- Upgraded to golang version 1.19
- Upgraded Steampipe sdk version to v4.1.7

## v0.1.0 [2022-05-05]

_Enhancements_
- Upgraded to golang version 1.18
- Upgraded Steampipe sdk version to v3.1.0
