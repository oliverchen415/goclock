# goclock
Automatically prints out a datetime of a specified timezone

Example:
```
.\goclock America Los_Angeles
```
produces:
```
DATETIME                              TIMEZONE
--------                              --------
2021-04-08 22:03:38.250126 -0700 PDT  America/Los_Angeles
```

## Installation
- If you're on Windows, download the .exe. Optionally, you can add it to your $PATH environment variable.
- On other OSes, you can build the project yourself. The `go.mod` and `go.sum` are available.

## Usage
If you didn't add it to your $PATH, you can run the file in a terminal by moving to the directory you downloaded the .exe to, e.g.
```
cd C:\Users\<user_name>\Downloads
```

and using the command
```
.\goclock <area> <location>
```
where `<area>` is a continent, and `<location>` is a city. If the location contains two words, then join them with an underscore. Case doesn't matter.

If you added it to your $PATH, then you don't need the `.\`. 

## Cool projects I used 
| Link        | Description |
| ----------- | ----------- |
| [araddon/dateparse](https://github.com/araddon/dateparse) | Parses datetimes without worrying about formats. |
| [cheynewallace/tabby](https://github.com/cheynewallace/tabby) | Easy (and pretty) way to format tables. |
| [World Time API](http://worldtimeapi.org/) | The API I used to get world times. [List of timezones.](http://worldtimeapi.org/timezones) |
