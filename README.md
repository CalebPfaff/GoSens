# GoSens

## Sensitivity generator for various games

Current featues:
* Generates a random sensitivity for selected game
* Accepts custom DPI and sensitivity range

Supported games:
* Overwatch
* Fortnite
* Counter-Strike/any Source game
* Quake Champions
* Rainbow Six Siege

## Usage

Open a terminal in the folder and execute 

Windows:
```
.\gosens.exe
```

Linux:
```
./gosens.exe
```

Command line options

-game : Choose the game to generate a sensitivity for

-min : Minimum cm/360

-max : Maximum cm/360

-dpi : Your mouse DPI

-games : List the available games

-debug : Show all the variable values

#### Example:

Generate an Overwatch sensititvity, with DPI of 800 between 20 and 40 cm/360
```
./gosens.exe -game ow -min 20 max 40 -dpi 800
```
And it spits out:
```
7.82 in Overwatch setttings (22.15 cm/360)
Settings: 800 DPI, 20cm - 40cm
```
Show all available games:
```
./gosens.exe -games

Games Dictionary:

Overwatch - ow
Fortnite - fn
Counter-Strike - csgo
Quake Champions - qc
Source Games - source
Rainbow Six Siege - r6
```

Show dubug info:
```
./gosens.exe -game ow -min 20 max 40 -dpi 800 -debug=true

8.34 in Overwatch setttings (20.77 cm/360)
Settings: 800 DPI, 20cm - 21cm

Debug Info:
-----------
Input Game: ow
Input DPI: 800
Input Min: 20
Input Max: 21

fullName: Overwatch
Precision: 2
Yaw: 0.0066
RandNum: 20.75698158540927
genOutput: 8.34
cm360: 20.765206017004576
```

## Notes

The program will default to Overwatch, 800 DPI, and 20-40 cm/360, and no debug info if no flags are provided
