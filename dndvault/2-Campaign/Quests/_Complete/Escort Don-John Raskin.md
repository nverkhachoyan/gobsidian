---
questObtained: 
questStatus: Not Started
questGiver: "[[Harbin Wester]]"
questLocationObtained: "[[Mountain's Toe Gold Mine]]"
questSessionObtained: 
questNotes: 
questLootAvail:
  - 100 GP
questLootEarned: 
NoteIcon: quest
obsidianUIMode: preview
tags:
  - quest
---

```leaflet
id: mountains_toe_gold_mine_map
lock: true
recenter: true
noScrollZoom: true
image: [[map_mountains_toe_gold_mine.webp]]
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

Escort [[Don-Jon Raskin]] to the [[Mountain's Toe Gold Mine]].

### Quest Objective

- Successfully get [[Don-Jon Raskin]] to the gold mine.

### Rewards

- 100 GP


### On Arrival
> [!readaloud] 
> 
> Hidden among bushes, a tunnel burrows into the foot of a soaring, snow-capped mountain. Above the mouth of the tunnel is a wooden plank with the words "Mountain's Toe" carved into it in Common.

#### M4. Wererat Den


##### Treasure

- [?] Under debris

Characters who search the cave find two sacks hidden under debris. One sack contains ten fist-sized chunks of gold ore worth 10 gp each. The other holds 82 sp, 450 cp, and a pair of [goggles of night](/3-Mechanics/CLI/items/goggles-of-night.md). If the characters obtain and identify the goggles, give the [Goggles of Night](/3-Mechanics/CLI/items/goggles-of-night.md) [Goggles of Night](/3-Mechanics/CLI/decks/magic-item-cards-dip.md#Goggles%20of%20Night) to the players or they can reference it in the Magic Items Listing.