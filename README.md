# HMAC based one time password

## Usage

    bin\main.exe [-options=value]

`make run` builds the main.go file and executes the binary with the mentioned parameters.

`make clean` deletes the main.exe file in the bin folder.

### Options

- `target`: The url to which you are sending the message using basic auth.
- `message`: The message to be sent to the target
- `userid`: User ID used for auth.
- `length`: The number of digits in the hotp.
- `hash`: Type of hash algorithm used.
- `message`: The message body to be sent in Json format.
- `secret`: Secret key used.
- `interval`: The time interval (TX) between hotp generation.
- `initial`: The start time (T0) used for calculating hotp.
