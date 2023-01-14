# Wiki

This project is deadski, but fun. Every few years I get the urge to do all kinds of knowledge management at work and run into the lack of a suitably low-friction wiki, which leads me to get salty at OneNote and SharePoint and occasionally ends with a fit of trying to write custom wiki software. This would be my latest attempt, circa 2018.

It had some grand vision of linking pages together in a knowledge graph a la [Obsidian](https://obsidian.md/), using [Cayley](https://cayley.io/) as an embedded graph database, but this idea never got off the ground and now finds better expression in a more content-management-geared project I've been working on.
It deploys to Google App Engine as that was a convenient place to stick a Go application in the cloud, but can also be run standalone on your machine.

There is a large number of transitive dependencies which I kinda don't love about it.

Originally this was built on Windows, and would take some nudging to set up on Linux.
