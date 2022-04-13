# Akademik Absen Automation

This golang script will run a browser instance, login to your akademik account, and do all of the absen.

## Usage

To use this script, follow the instruction below

1. Install Go

   The Go languange used in this project are `go1.18`. Install Go from the official website and follow the [instructions](https://go.dev/doc/install).

2. Run the script

   To run the script there's a few option that you need to configure first. The script used environment variable to get the username and password of the user.

   Set the `USERNAME` variable with your akademik username, and `PASSWORD` with the password of your akademik.

   Then to run it, use `go run main.go`, example of how to do it in Linux,

   ```bash
   USERNAME=191524*** PASSWORD=bukanpassword go run main.go
   ```

   There's also an optional environment variable `rod` to debug the browser itself. Set it to `show` to make the browser visible.

## Contributing

Contribution are absolutely welcomed! Go add extra feature, add CRON functionaliy, whatever.

## License

MIT
