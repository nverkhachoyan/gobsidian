---
NoteIcon: settlement
tags:
  - Category/Settlement
Community-Size: Small Village
Alignment: Good
Government: Elected Leader
type: Settlement
politics: Townmaster
leader: Harbin Wester
guildsgroups: 
region:
  - Phandalin
  - Sword Coast North
size: Small Town
population: 70
commonraces:
  - Humans
religion:
  - Tymora
  - Waukeen
exports:
  - Cold Iron
  - Gold
  - Platinum
  - Skins
  - Timber
imports:
---

```leaflet

id: phandalin_map

lock: true

recenter: true
noScrollZoom: true

image: [[map_phandalin.jpg]]

### Map Pixel Height x 1 / (Pixels between Bar Scale / 100)

### Map Pixel Width x 1 / (Pixels between Bar Scale / 100) 

### Note that this formula requires adjustments depending on your map. The idea is to determine the number of units between your bar scale. We divide by 100 here because my bar scale measures in 100 units. If your maps scale bar measures in units of 50 them you should divide by 50 instead. The idea is to calculate how many pixels are equal to 1 unit. 

bounds: [[0,0], [864.8393, 1205.104]]

height: 900px

width: 95%

### This sets where the map starts by default. Set it to the middle (half) of your bounds. 

lat: 428.3708

long: 596.9101

### 0 is no zoom. Negative zoom steps away from the map. Positive zoom steps towards the map. 

minZoom: -1.5

### Max zoom is 18. 

maxZoom: 

### Hover mouse over the Reset Zoom icon to see your current zoom level. 

defaultZoom: -0.5


### How far it zooms in or out with each step. Can be in decimals. 

zoomDelta: 0.5

### This is a string so can be any text. Change it to match your maps measurement scale. 

unit: feet

scale: 1

darkMode: false

```

> [!infobox]
> # `=this.file.name`
> ![[Phandalin_art.webp|cover hsmall]]
> ###### Geography
> Type |  Stat |
> ---|---|
> Type | `=this.type` |
> Size | `=this.size` |
> Region | `=this.region` |
> ###### Travel (`=[[Travel Calculator]].HoursPerDay` hrs per day)
> ###### [[Travel Calculator]]  / [[Exhaustion]]:  `=[[Travel Calculator]].ExhaustionLevel`
> Destination |  Travel Days  |
> ---|---|
> [[Leilon]] | 🕓: `VIEW[round((46* {Travel Calculator#TravelCalc}) / 60 / {Travel Calculator#HoursPerDay}, 1)]`      |
> [[Dwarven Excavation]] | 🕓: `VIEW[round((15* {Travel Calculator#TravelCalc}) / 60 / {Travel Calculator#HoursPerDay}, 1)]`      |
> [[Umbrage Hill]] | 🕓: `VIEW[round((5* {Travel Calculator#TravelCalc}) / 60 / {Travel Calculator#HoursPerDay}, 1)]`      |
> ###### Politics
> Type |  Stat |
> ---|---|
> Govt Type | `=this.politics` |
> Ruler | `=this.leader` |
> Defense | `=this.defences` |
> ###### Society
> Type |  Stat |
> ---|---|
> Population | `=this.population` |
> Races | `=this.commonraces` |
> Temples | `=this.religion`  |
> ###### Commerce
> Type |  Stat |
> ---|---|
> Exports | `=this.exports` |
> Imports | `=this.imports` |
> ###### Organizations
> Type |  Stat |
> ---|---|
> ```dataview
table WITHOUT ID link(file.name) AS "Group", link(Leader) AS "Leader"
where contains( PrimaryHome, this.file.name)


# `=this.file.name`
## Overview

Phandalin is a small village, which acted as a frontier settlement of farmers and prospectors.

## Notable NPCs
Placeholder

## Profile
Placeholder

## Story
Placeholder

## Points of Interest
Placeholder

## Valuables
Placeholder

## Internal Relationships
Placeholder

## Outward Relationships
Placeholder

## Background
Placeholder

## Additional Details
Placeholder
