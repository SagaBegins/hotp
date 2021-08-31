# HMAC based one time password

## Usage

    bin\main.exe [-options=value]

`make run` automatically builds the main.go file and executes the binary with the mentioned parameters.

`make clean` deletes the main.exe file in the bin folder.

### Options

- `target`: The url to which you are sending the message using basic auth.
- `message`: The message to be sent to the target
- `userid`: User ID used for auth.
- `passLength`: The number of digits in the hotp.
- `hashAlgorithm`: Type of hashAlgorithm by hmac.
- `message`: The message body to be sent in Json format.
- `secretKey`: Secret key used by hmac.
- `interval`: The time interval between hotp generation.
- `initial`: The start time used for calculation of hotp.
