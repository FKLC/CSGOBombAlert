# This project is not working.
After publishing project on Reddit [JHunz](https://www.reddit.com/user/JHunz/) said
	
> [I don't know how much testing you've done, so you may already know this. But Valve intentionally obfuscates the timing of bomb plants in the game state API by plus or minus a few seconds to prevent stuff like this from presenting an in-game advantage. I think it may give you reliable timing when you yourself have planted the bomb, however it is definitely not reliable when teammates and/or enemies have planted.](https://www.reddit.com/r/golang/comments/c09ls5/go_based_csgo_bomb_alert/er2ykw6)

So I recorded screen to be sure about it and yes it adds delay to requests. I'm not going to delete it but I'll archive the project.

# CS:GO Bomb Alert

CS:GO Bomb Alert is a program that alerts you 30 and 35 seconds after bomb planted in CS:GO.

## Installation

1. First download the [latest release](https://github.com/FKLC/CSGOBombAlert/releases/latest/download/CSGOBombAlert.exe), and [config file](https://github.com/FKLC/CSGOBombAlert/releases/latest/download/gamestate_integration_bombalert.cfg).
2. Open CS:GO folder. (the following 3 steps taken from [here](https://dmarket.com/blog/csgo-autoexec/))
   1. **Open Steam app** and visit your game Library.
   2. **Right-click CS:GO** in the list of your games, choose “**Properties**”.
   3. In the opened window, go to the tab “**Local Files**” and press the button “**Browse Local Files…**”.
   4. Open the folder "**csgo -> cfg**".
3. Move the downloaded config file to here.
4. Run the program itself.

## Will I get ban for using this?

Well, the answer is I don't know but I don't think so because program only uses [CS:GO Game State Integration API](https://developer.valvesoftware.com/wiki/Counter-Strike:_Global_Offensive_Game_State_Integration) which is completely legal but as I said, you can get ban for using it.

**I do not guarantee you that you will not get banned for using this.**

<sup><sup>BTW I started learning Go yesterday so if you think my code is bad yeah I know it is</sup></sup>
