# MattJam [![Release](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/blob/master/assets/release-img.svg)](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/releases/latest)
![Logo MattJam](https://user-images.githubusercontent.com/45047261/123288194-a090c500-d50f-11eb-9ffe-17d4f453b9af.png)

**Maintainer:** [Abderrahmane Mouzoune](https://github.com/AbderrahmaneMouzoune)

TABLE OF CONTENTS
------------------

* [Introduction](#introduction)
* [Features](#features)
* [Versions](#versions)
* [Contribute](#contribute)
* [Q&A](#qa)
* [Requirements](#requirements)
* [Copyright](#copyright)

## Introduction

The MattJam plugin is a free open source software for Mattermost that provides the possibility to create jams 
directly in the chat or during call sessions. MattJam provides a fun and easy way 
to collaborate during team projects and animate working and brainstorming sessions even in remote. 

![Preview plugin](https://user-images.githubusercontent.com/45047261/123388566-37559400-d599-11eb-88e4-051457a56d59.png)


*In developement* you can see more in the roadmap [here](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/projects)

## Features
Our MattJam FOSS is an all-in-one plugin has so much functionnalities to hype your Mattermost sessions, you can acces to all of this features with ```/mattjam``` :

* Full whiteboard creation with shapes, text, colors, post-its, arrow, and so much more
* You can add who can watch or edit the jam during the session
* You can download your jams and send them directly to everyone
* You can add a small resume of the session so everyone can remember or understand what happened during the jam session
* Jam versionning that allows you to look each versions of the same jam
* Resume system that save all the jams at the end of the session, and allows you to consult them at any time
* tag system to find easily any jam in your history with keywords
* reporting system so the admin can see what are MattJam stats 
* possibility to create call sessions in same time as jam sessions so there is no need of another plugin

## Versions

If you want to see each release you can see that [here](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/tags)

If you want to see how each versions work you can see that [here](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/projects)

If you want to see the structure of the project you can see that [here](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/wiki/How-we-see-our-project-%3F)

## Contribute

We welcome contributions for bug reports, issues, feature requests, feature implementations, and pull requests. Feel free to file a new issue ????

For a complete guide on contributing to the plugin, see the [Contribution Guidelines](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/blob/master/CONTRIBUTING.md).

## Q&A

### How do I install the plugin?
I can download it [here](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/tags)

1. Log into Mattermost as a system admin
2. In the upper-left corner, select your username, then System Console
3. On the left, under Plugins, select Plugin Management
4. Find "upload a plugin" and upload the tar.gz file download [here](https://github.com/AbderrahmaneMouzoune/mattermost-plugin-mattjam/tags)
5. For Enable Plugins, select true
6. Save changes at the bottom

### How do I create a jam?
In the concerned chat bar, type ```/mattjam``` then press ```enter```.

### How do I save my jam ?
It is an automatic backup, which stops at the end of the jam session.

### How do I find the jams I have participated in?
Go to the ?? Jam History ?? channel on the left sidebar of Mattermost.

### How do I share a jam with other people?
When you move the mouse over the jam area or from the jam history channel, a menu of options appears: choose the icon (ajouter l???ic??ne)

### Can I join a jam in progress?
Yes, via the share link, the ?? Jam History" channel or via the relevant chat if you are a member.

### Can more than one person edit the jam?
Currently, any participant has the possibility to contribute to the jam.

### Can I make a jam on a call session?
This is not yet possible directly from our plugin, but the call function will be available in version 4.0.

## Requirements

* Golang >1.16
* Npm > 7.15


| Operating System   | Technical Requirement              |
| ------------------ |:---------------------------------- |
| Windows            | Windows 7, 8.1, and 10             |
| Mac                | MacOS 10.12+                       |
| Linux              | Ubuntu LTS releases 18.04 or later |

Though not officially supported, the Linux desktop app also runs on RHEL/CentOS 7+.

| Browser            | Technical Requirement              |
| ------------------ |:---------------------------------- |
| Chrome             | v77+                               |
| Firefox            | v68+                               |
| Edge               | v44+                               |
| Safari             | v12+                               |

## Copyright

License MIT.

Thanks to :
* [Olivia Pinto](https://www.linkedin.com/in/olivia-pinto-3a8b9779/)
* [Lucas Audon](https://www.linkedin.com/in/lucas-audon-871784137/)
* [Agathe Frangeul](https://www.linkedin.com/in/agathe-frangeul-14391117a/)
* [Manon Landrin](https://www.linkedin.com/in/manon-landrin/)
* [Hugo Custodio](https://www.linkedin.com/in/hugo-custodio/)
* [Ryan Fennane](github.com/RyanFennane)
* [Abderrahmane Mouzoune](github.com/AbderrahmaneMouzoune)
* [Jitsi plugin](github.com/mattermost/mattermost-plugin-jitsi)
