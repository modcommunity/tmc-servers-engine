# TMC Server Engine
A Go program that updates server's real-time information such as member/player counts, map names (for game servers), and more from [The Modding Community](https://moddingcommunity.com/) (the [@modcommunity](https://github.com/modcommunity)) while interfacing with its API. This will also be generating stats, ranks, and banners for these servers in the future!

## In Development
[@Gamemann](https://github.com/gamemann) has been working on an application that lists community and game servers that may be tied by communities (AKA clubs) in [Invision Power Services](https://invisioncommunity.com/). Please see below for a preview of how everything is looking. This is being internally developed at the moment, but we will be looking into open-sourcing the application in the future after business-related things are situated in the [@modcommunity](https://github.com/modcommunity) along with ensuring we aren't putting user's security at risk (e.g. having multiple views at source code for any potential vulnerabilities).

### Goal
The goal of this project is to retrieve servers from the [@modcommunity](https://github.com/modcommunity) and update their real-time information (member counts from Discord servers, game server stats such as player counts, map names, etc.), generate statistics, and more! This project will act as an engine and it will include functionality to select what type of querying system to retrieve server information with. For example, [A2S_INFO](https://developer.valvesoftware.com/wiki/Server_queries) queries which is implemented by the Source Engine and Valve.

## <a href="https://moddingcommunity.com/" target="_blank"><img src="misc/goal_hd.gif" data-canonical-src="https://github.com/gamemann/tmc-servers-engine/misc/goal_hd.gif" /></a>

Each engine type will spawn its own thread and retrieve servers (preferably sorted by the last time the server was queried the specific engine, then updating these values with Unix timestamps when updating a server's information). Support for spawning multiple threads/Go routines will also be added for best multi-threading support. Though, I don't think multiple threads per engine will matter much.

## Credits
* [Christian Deacon](https://github.com/gamemann)