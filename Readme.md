# dbus-wait

This simple program blocks until the specified D-Bus event is fired.

## Installation

Simply run `make install`.
Or even better: Use the provided nix package.

## Usage

    dbus-wait [--session] [match options]

The program connect to the *system* bus by default.
If the `--session` option is enabled, it connects to the session bus instead.
Match options are key value pairs corresponding to the [match rules](https://dbus.freedesktop.org/doc/dbus-specification.html#message-bus-routing-match-rules).
