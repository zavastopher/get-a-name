# get-a-name

get-a-name is a go module for randomly picking a name from name.txt files containing lists of names.

<div align="center">
    <img src="https://github.com/zavastopher/get-a-name/blob/main/resources/demo.gif" alt="gif of get-a-name working">
</div>

_let the computer roll and look for the random name for that npc you don't care about_

This project is inspired by various world building tools such as [Fantasy name generator](https://www.fantasynamegenerators.com/) and [Fantasy name Generator _for obsidian_](https://github.com/lukewh/fantasy-name)

# Usage

This project is currently going through a rewrite in order to use a database for storing names

# Structure

This project is configured as an API that can run on a server or on your local machine if desired.
There are two components in the architecture of the project.

## The Database

The database is a docker container running the postgresql docker image. As of right now this container
contains one database: randomtables. This tables contains the following table:

### Names

This table contains entries for names that are referenced when calling the name generation route. The
name entries must have a Category and Culture although these fields are not necessarily referenced
everytime a name is generated.

[License](./LICENSE)
