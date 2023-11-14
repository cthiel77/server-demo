# server-demo
![Coverage](https://img.shields.io/badge/Coverage-79.9%25-brightgreen)
a server application demo


## flags and operating modes

Operating modes have been implemented for simple control of log outputs, etc. They are intended to set bundled settings (flags) with settings optimized for the mode to ensure optimal operation. The settings apply at runtime, so a restart is required with each change. It is also possible to influence the default settings during binary builds via ld-flags.

| flag            | available values         | defaults | read only | description                                         |
| :-------------- | :----------------------- | :------- | :-------: | :-------------------------------------------------- |
| -m, --mode      | dev, prod                | dev      |           | operating mode and runs the app                     |
| -l, --log-level | info, warn, error, fatal | info     |           | log level and runs the app                          |
| --author        | ---                      | ---      |     ✔     | prints information about the autor and exits        |
| --company       | ---                      | ---      |     ✔     | prints information about the manufacturer and exits |
| -h, --help      | ---                      | ---      |     ✔     | prints help text and exits                          |
| --license       | ---                      | ---      |     ✔     | prints license text and exits                       |
| -v, --version   | ---                      | ---      |     ✔     | prints version number and exits                     |

