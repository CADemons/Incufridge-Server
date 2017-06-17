# How to set up the incufridge

There are a number of different programs required to set up the incufridge.

### Steps

When you are ready to start everything, here are the steps (you'll read about the programs below):

1. Make sure the arduino is running the right program
2. Start the server
3. Connect the arduino and the raspberry pi with the cable
4. Turn on the incufridge (tap the fan if necessary as well)
5. Start the Incufridge-Client on the raspberry pi
6. Start the Incufridge-App on your computer

Hopefully everything goes well.

### [incufridge](https://github.com/cademons/incufridge)
This is the arduino code. Upload the code from here onto an arduino. The code will stay there forever unless some new program is uploaded so after you do it once there is no need to do anything more.

### [Incufridge-Client](https://github.com/cademons/Incufridge-Client)
This is the Java code for the Raspberry Pi client. Make sure `java` and `javac` are installed and working on the Raspberry Pi. You can begin the program with 

```
$ ./run.sh
```

(this should compile and run) from the terminal in the root directory of the project (on the raspberry pi we have been using, this runs automatically on startup). Make sure you plug the arduino into the raspberry pi and turn everything on too (remember to tap the fan in the back too if you are still using the same fridge I was).

You will need the `auth` file to be in the current directory.

### [Incufridge-Server](https://github.com/cademons/Incufridge-Server)
This is the Go code for the server which runs on cademons.org. To start it, run `ssh incufridge@cademons.org` (email me if you forget the password), and then there should be a symlink called `incufridge` to the directory with the code. `cd` into that directory. You can start the server by running 

```
$ ./tcpserver
```

You will need the `auth` file to be in the current directory.

This won't let you logout though so if you want to start the server and be able to logout with it still running in the background, run:

```
$ nohup ./tcpserver &
```

You can also edit the code and rebuild with `go build tcpserver.go`.

All debug output will be sent to a file called `log.txt`.

If you get an error that the port is in use then the server is already running. If you want to stop it, you can grep the running processes for it:

```
$ ps aux | grep tcpserver
```

This will give you output like:

```
1050     22948  0.0  0.0 834700  1632 ?        Sl   00:50   0:00 ./tcpserver
1050     23847  0.0  0.0   4032   700 pts/0    S+   00:53   0:00 grep tcpserver
```

The first job is the important one (the second one is your grep search), and the job number in this case is `22948`.

Then you can kill it:

```
$ kill 22948
```

You can grep again to make sure it's gone.

### [Incufridge-App](https://github.com/CADemons/Incufridge-App)
This is a slightly modified version of the Incufridge-Client that is meant to run on the user's computer. Basically all the commands have been changed so that instead of sending serial events to the arduino they send tcp events. There is probably a bunch of excess code stored in this repository (you can try to find and delete it if you want).

You can run it, just like before, with:

```
$ ./run.sh
```

You will need the `auth` file to be in the current directory.

Unless you expect users to have the JDK on their laptop, I recommend you turn this into an executable jar file when it is ready for use (you can use Eclipse to do that).

# How to use the incufridge

At this point hopefully everything is set up and working and you should have the Incufridge-App code running on your computer.

When you start the program, you'll see a screen which will let you view the current temperature (click the `Update` button). The `Get log` button will tell you the current temperature and target temperature and save it to a file with the timestamp.

Then you can go to the `Console` tab. Here you can send individual commands to the incufridge. Here are the possible commands:

* `set temp ##`: set the temperature
* `createlog`: create a log file
* `every ## unit date time ( command ) jobname`: start a job which will run a command every period of time.
* `at date time ( command) jobname`: same thing as `every` but it only gets run once.
* `cancelall`: stops all jobs
* `cancel jobname`: stops `jobname`.

In the `Recipe` panel, you can write little programs made of these commands and run them all at once.

### Date, time, units

NOTE: Spacing is very important

Here is an example of an `every` command:

```
every 1 minute today now ( set temp 30 ) temp30
```

This creates a job called `temp30` which will set the temperature to 30 every minute starting today,now (if you want to run multiple commands in one job, you can separate them with semicolons). Do not remove the spaces next to the parentheses.

Here is an example of at:

```
at today now+3 ( set temp 30 ) temp30
```

This creates a job called `temp30` which will set the temp to 30 in 3 minutes.

Valid units:

* `seconds`
* `minutes`
* `hours`
* `days`

Valid dates:

* `10/21/2017` (that's month, day, year: `MM/dd/yyyy`)
* `today`
* `today+##` (you can add any amount in days to today)

Valid times:

* `10:00` (Use a 24-hour clock)
* `now`
* `now+##` (you can add any amount in minutes to now)

# End

Good luck! I hope things work somewhat.
