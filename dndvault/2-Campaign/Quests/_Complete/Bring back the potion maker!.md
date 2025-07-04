---
questObtained: 
questStatus: Complete
questGiver: "[[Harbin Wester]]"
questLocationObtained: "[[Umbrage Hill]]"
questSessionObtained: "[[1-Session Journals/2024-11-12.md|2024-11-12]]"
questNotes: 
questLootAvail:
  - "[[potion-of-healing|Potion of Healing]]"
  - 25gp
questLootEarned: 
NoteIcon: quest
obsidianUIMode: preview
tags:
  - quest
---

```leaflet
id: umbrage_hill_map
lock: true
recenter: true
noScrollZoom: true
image: [[map_umbrage_hill.webp]]
### Map Pixel Height x 1 / (Pixels between Bar Scale / 100)
### Map Pixel Width x 1 / (Pixels between Bar Scale / 100) 
### Note that this formula requires adjustments depending on your map. The idea is to determine the number of units between your bar scale. We divide by 100 here because my bar scale measures in 100 units. If your maps scale bar measures in units of 50 them you should divide by 50 instead. The idea is to calculate how many pixels are equal to 1 unit. 
bounds: [[0,0], [204.2222, 320.8889]]
height: 900px
width: 95%
### This sets where the map starts by default. Set it to the middle (half) of your bounds. 
lat: 102.1111
long: 160.4444
### 0 is no zoom. Negative zoom steps away from the map. Positive zoom steps towards the map. 
minZoom: 1.5
### Max zoom is 18. 
maxZoom: 2.5
### Hover mouse over the Reset Zoom icon to see your current zoom level. 
defaultZoom: 2
### How far it zooms in or out with each step. Can be in decimals. 
zoomDelta: 0.5
### This is a string so can be any text. Change it to match your maps measurement scale. 
unit: feet
scale: 1
darkMode: false
```


# `=this.file.name`

> [!infobox]+
> # `=this.file.name`
> ## Quest Details
> Type |  Stat |
> ---|---|
> Date Obtained | `INPUT[datePicker:questObtained]` |
> Status | `INPUT[inlineSelect(option(Not Started), option(In Progress), option(Complete)):questStatus]` |
> Quest Giver | `INPUT[suggester(optionQuery(#npc)):questGiver]` |
> Quest Location | `INPUT[suggester(optionQuery(#Category/Settlement)):questLocationObtained]` |
> Session Obtained | `INPUT[suggester(optionQuery(#journal)):questSessionObtained]` |
> Available Loot | `=this.questLootAvail` |
> Acquired Loot | `=this.questLootEarned` |


> [!readaloud]
> "The local midwife—an acolyte of [[forgotten-realms-chauntea|Chauntea]] named [[Adabra Gwynn]]—lives by herself in a stone windmill on the side of a hill a few miles south of Phandalin. With dragon sightings becoming more common, it's not safe for her to be alone. Urge Adabra to return to Phandalin. Once she's safe, visit Townmaster Harbin Wester to claim a reward of 25 gp." If the characters undertake this quest, proceed with "Umbrage Hill." 

### Quest Objective

Adabra declines to return to Phandalin, but the characters can complete the Umbrage Hill Quest by asking her for a note for Harbin Wester confirming her safety. 

- [?] Adabra also gives her saviors one [potion of healing](/3-Mechanics/CLI/items/potion-of-healing.md) for dealing with the manticore.

### Rewards

- [[potion-of-healing|Potion of Healing]]

### On Arrival

> [!readaloud] 
> 
> Built on the slope of Umbrage Hill is an old stone windmill surrounded by an iron fence. A large winged monster with a spiky tail is trying to knock down the windmill's front door. A woman appears in a second-floor window, waves at you, and yells, "A little help?!"

```encounter
name: Bring back the potion maker! (U1)
creatures:
  - "1": Manticore
  - "1": Adabra Gwynn, friendly
```

### U1. Dwarven Cairns

Dwarves were buried under these rock piles, but their bones, armor, and weapons disintegrated long ago.

### U2. Ruins of a House

A fireplace is all that remains of this old stone house.

### U3–U5. Stone Windmill

Adabra uses the millstone on the ground floor (area U3) to grind herbs and other ingredients for poultices and potions. Her quarters are on the second floor (area U4). The third floor (area U5) is a loft where guests can stay.