# TOTP - TIME-BASED ONE-TIME PASSWORD

The program reads a base32-encoded secret from the command-line and generates a 6-characters password.
The password varies over time: every 30 seconds, the password changes - hence the term "time-based".

Install `GO` and in current folder compile to `totp.exe`:

    > go build -ldflags "-s" .