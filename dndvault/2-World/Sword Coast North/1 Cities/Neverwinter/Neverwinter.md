---
NoteIcon: settlement
tags:
  - Category/Settlement
Community-Size: City
Alignment: Neutral
Government: Autocracy
type: City
politics: Lordship
leader: 
guildsgroups: 
region:
  - Sword Coast North
size: Large City
population: 90000
commonraces: 
religion: 
exports: 
imports:
---

> [!readaloud]  
> "As the road crests the final hill, the city of Neverwinter unfolds before you like a painting sprung to life. Great stone bridges span the river that runs through its heart, connecting bustling districts and towering spires. Steam rises from the cobblestones in places where the land still remembers the fire beneath, and the scent of sea air mixes with incense, horses, and fresh-baked bread. Ships crowd the harbor, flags snapping in the wind, while voices in a dozen tongues echo from markets and alleyways alike. Above it all looms Castle Never—silent, broken, and watching."

![[NeverwinterMapLabels.jpg]]


> [!infobox]
> # `=this.file.name`
> ![[MapPlaceholder.png|cover hsmall]]
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
> [[Voonlar]] | 🕓: `VIEW[round((88* {Travel Calculator#TravelCalc}) / 60 / {Travel Calculator#HoursPerDay}, 1)]`      |
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
Placeholder

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

