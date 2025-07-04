---
NoteIcon: region
tags:
  - Category/Region
aliases: 
commonraces:
  - Humans
  - Elves
  - Dwarves
  - Orcs
  - Half-Orcs
  - Half-Elves
  - Tieflings
alignments:
  - All
region:
  - North-West Faerun
largestcity:
  - Neverwinter
languages:
  - Common
religions:
  - Nearly All
---

```leaflet
### Tutorial: https://youtu.be/54EyMzJP5DU
### id must be unique
id: sword_mountains_map
### Lock pins so they can't be moved
lock: true
### If true, view of map will recenter as you zoom out. 
recenter: true
### If true, disables mouse scroll for zomming in and out of a map. Button controls still work. 
noScrollZoom: true
image: [[map_sword_mountains.jpg]]
### Map Pixel Height x 1 / (Pixels between Bar Scale / 100)
### Map Pixel Width x 1 / (Pixels between Bar Scale / 100) 
### Note that this formula requires adjustments depending on your map. The idea is to determine the number of units between your bar scale. We divide by 100 here because my bar scale measures in 100 units. If your maps scale bar measures in units of 50 them you should divide by 50 instead. The idea is to calculate how many pixels are equal to 1 unit. 
bounds: [[0,0], [152.7604, 109.6354]]
height: 900px
width: 95%
### This sets where the map starts by default. Set it to the middle (half) of your bounds. 
lat: 76.3
long: 54.8
### 0 is no zoom. Negative zoom steps away from the map. Positive zoom steps towards the map. 
minZoom: 2
### Max zoom is 18. 
maxZoom: 4.5
### Hover mouse over the Reset Zoom icon to see your current zoom level. 
defaultZoom: 3
### How far it zooms in or out with each step. Can be in decimals. 
zoomDelta: 0.5
### This is a string so can be any text. Change it to match your maps measurement scale. 
unit: miles
scale: 1
darkMode: false
```

> [!infobox]
> # `=this.file.name`
> ![[sword_mountains_north.jpg|cover hsmall]]
> ###### Geography
> Type |  Stat |
> ---|---|
> Aliases|`=this.aliases`|
> Region | `=this.region` |
> Largest City|`=this.largestcity`|
> ###### Society 
> Type |  Stat |
> ---|---|
> Races|`=this.commonraces`|
> Languages|`=this.languages`|
> Religions|`=this.religions`|
>Alignments|`=this.alignments`|

# `=this.file.name`

## Overview
Placeholder

## Geography
Placeholder

## Government
Placeholder

## Society
Placeholder

## Trade
Placeholder

##  Defenses
Placeholder

## History
Placeholder

## Rumors and Legends
Placeholder

## Inhabitants
Placeholder

## Additional Details
Placeholder
