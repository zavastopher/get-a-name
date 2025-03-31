# Get-A-Name

Get-A-Name is a command-line TTRPG Character generation tool. Current functionality being focused on name generation.

<div align="center">
    <img src="https://github.com/zavastopher/get-a-name/blob/main/resources/demo.gif" alt="gif of get-a-name working">
</div>

_let the computer roll and look for the random name for that npc you don't care about_

# Motivation
This project is inspired by various world building tools such as [Fantasy name generator](https://www.fantasynamegenerators.com/) and [Fantasy name Generator _for obsidian_](https://github.com/lukewh/fantasy-name). With the main goals being as follows.

- Be run in the commandline
- Have a database that can be self-hosted and updated as needed
- Built with go for performance and developer experience

# Quick Start

TBD - My current plans are the following 2 versions
- A simple all in one executable that comes pre-wrapped with a default list of names for character generation
- An executable that can be set up to point to a self-hosted database for name generation. 

# Usage

TBD - This project is currently going through a rewrite in order to use a database for storing names. As of right now it is containerized using docker so if you are so inclined you can run the docker containers using docker compose from the root directory and mess around with the project that way. 

## Structure

This project is configured as an API that can run on a server or on your local machine if desired.
There are two components in the architecture of the project.

### The Database

The database is a docker container running the postgresql docker image. As of right now this container
contains one database: randomtables. This tables contains the following table:

#### Names

This table contains entries for names that are referenced when calling the name generation route. The
name entries must have a Category and Culture although these fields are not necessarily referenced
everytime a name is generated.

# Contributing 

Help is always appreciated!

Contribute by forking the repo and opening pull requests. Writing tests for your code is always appreciated. Also do try to document the changes and additions you make to make it easier to review your changes. 

All pull requests should be submitted to the main branch.

[License](./LICENSE)