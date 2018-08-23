Changes

## 1.4.0 (2018-08-23)

### Changed
- compatible with dcp 0f83d9b6

## 1.3.0 (2018-08-21)

### Changed
- improve the query speed of the host group 18c68619

## 1.2.0 (2018-07-30)

### Added
- add a rpc service to get the host strategies 9ca8f2e2
- add control script d0da000f

### Changed
- fix typo error 0fe888e5

## 1.1.0 (2018-07-12)

### Added
- add an interface to maintain the host e3b06f17
- add an interface to remove the hosts e3b06f17
- add an interface to obtain alone host e3b06f17
- add an interface to get the monitor screen by hostname f5ce0de9

### Changed
- increase the limit to query grafana e072eb77

## 1.0.2 (2018-06-19)

### Changed
- feature improvement, do not remove all hosts in the host group before adding
  a host to the host group #bc01622d

### Fixed
- fix the value of result out of range #f6188a30
- update the field names of host #4d49e29e

## 1.0.1 (2018-06-18)

### Changed
- change the uri of get template #19018e91

### Fixed
- sql raise a error when given page and limit #09555ea7

## 1.0.0 (2018-06-18)

### Added
- add an interface to get a template by name #cfa0109d
- add an interface to get the host by hostname or ip #3ac13d0a

### Changed
- change the response to return the itself of item #cfa0109d
