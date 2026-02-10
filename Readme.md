# dbus-wait

This simple program blocks until the specified D-Bus event is fired.

## Installation

Simply run `make install`.
Or even better: Use the provided nix package.

## Usage

    dbus-wait [--session] [match options] -- [command]

The program connects to the *system* bus by default.
If the `--session` option is enabled, it connects to the session bus instead.
Match options are key value pairs corresponding to the [match
rules](https://dbus.freedesktop.org/doc/dbus-specification.html#message-bus-routing-match-rules).
Once the connection is set up and the program is waiting, it concurrently
executes the `command` passed after the end of option processing. Note that
currently it is impossible to wait for *any* D-Bus event and specify a command at
the same time.
