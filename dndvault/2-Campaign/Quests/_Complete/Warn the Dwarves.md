---
questObtained: 
questStatus: In Progress
questGiver: "[[Harbin Wester]]"
questLocationObtained: "[[Dwarven Excavation]]"
questSessionObtained: "[[2024-12-10]]"
questNotes: 
questLootAvail:
  - "[[sending-stones|Sending Stones]]"
  - "[[Holy Symbol of Abbathor]]"
  - "[[10-gp-gemstones|15 10gp Gemstones]]"
  - 50gp
questLootEarned:
  - "[[Holy Symbol of Abbathor]]"
  - "[[sending-stones|Sending Stones]]"
  - "[[10-gp-gemstones | 15 10gp Gemstones]]"
NoteIcon: quest
obsidianUIMode: preview
tags:
  - quest
---

```leaflet
id: dwarven_excavation_map
lock: true
recenter: true
noScrollZoom: true
image: [[map_dwarven_excavation_quest.jpg]]
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
>  “Dwarf prospectors found ancient dwarven ruins in the mountains southwest of here, and have been working an archaeological dig seeking treasure and relics. They need to be warned that a white dragon has moved into the area. Take the warning to them, then return to Townmaster Harbin Wester to collect a reward of 50 gp." 

### Quest Objective

- [!] Find [[Dazyln Grayshard]] and [[Norbus Ironrune]].
- [!] Warn [[Dazyln Grayshard]] and [[Norbus Ironrune]] about the young white dragon.
- [?] Help the Dwarves eliminate nearby enemies.
- [?] Kill all orcs that may appear if PCs take too long.

### Rewards

- [[sending-stones|Sending Stones]]
- [[Holy Symbol of Abbathor]]
- 15 [[10-gp-gemstones|10gp Gemstones]]
- 50gp

### On Arrival
> [!readaloud] 
> 
> The canyon's rocky walls rise to a height of eighty feet. At the end of the canyon, a twenty-foot-high wall of black stone has a broken gate carved into it, with one stone door hanging precariously by a hinge and the other door missing. Beyond this open gate, in the shadow of a great mountain to the east, lies a ruined settlement. All is quiet.

- [?] If the PC's announce their arrival, [[Dazyln Grayshard]] and [[Norbus Ironrune]] greet them.

### E1. Canyon

The canyon floor is strewn with rubble and bereft of vegetation except for a few tough, scraggly weeds.

### E2. Ruined Settlement

- [?] This was a settlement, of which nothing but an outline of the buildings remain.
- [?] It appears that an avalanche swept through the area and destroyed the settlement that laid here before.
- [>] A stone well on the south side of the ruins is filled with rubble.
- [?] The settlement has been ransacked, but players who  look find crushed dwarf bones.

### E3. Courtyard and Temple Facade

> [!readaloud] 
> 
> A partially collapsed, ten-foot-high wall separates this courtyard from the settlement west of it. Three heaps of rubble are piled high in this area. Hewn from the canyon's back wall, a thirty-foot-high temple facade features steps rising to a stone platform. Cut into this facade is a ten-foot-high open doorway flanked by crumbling, life-sized granite statues of cloaked dwarves. Evil grins can be seen on their weatherworn faces.

- [?] Passive Perception reveals [[Dazyln Grayshard]] and [[Norbus Ironrune]] talking while eating behind the easternmost rubble pile.
- [?] On closer inspection, each dwarf  has 10 days of [[rations-1-day|rations]], a [[waterskin]], mining tools and one of two matching [[sending-stones|sending stones]].
- [>] On interaction, the dwarves offer to give the players, the [[sending-stones|Sending Stones]] if the help them kill the [[ochre-jelly|ochre jellies]] further down in the temple.

> [!note] Temple Features
> 
> The temple (areas E4 through E11) is smoothly hewn from solid rock. The following features are common throughout.
> 
> **Ceilings.** Ceilings throughout are 10 feet high and flat.
> 
> **Doors.** All doors are made of carved stone with stone pins for hinges. Secret doors blend in with the surrounding stonework. Finding a secret door requires a search of the wall and a successful DC 15 Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) check. Dazlyn and Norbus have found one secret door thus far, in area E4. The others have yet to be discovered.
> 
> **Dust and Debris.** The temple is choked with dusty rubble. Rubble-filled squares are "difficult terrain" (see "the rulebook").
> 
> **Light.** There are no light sources in the temple, since the dwarf priests of Abbathor relied on [darkvision](/3-Mechanics/CLI/rules/senses.md#darkvision) to see.

### E4. Dusty Foyer

Spikes hammered into the floor hold open two sets of double doors. A once-secret door to the north stands open, with a room beyond it holding nothing but rubble.

- [?] A once secret door to the north stands open, with a room holding nothing but rubble.
- [>] The dwarves inform the party about the secret door they have found to the south of this room.

### E5. Temple

- [>] One [[ochre-jelly|ochre jelly]] scared off the dwarves and cling to the ceiling above the altar, another clings to the south wall.
```encounter
name: Warn the Dwarves (Location E7)
creatures:
  - "2": Ochre Jelly
  - "1": Norbus Ironrune, friendly
  - "1": Dazyln Grayshard, friendly
```

#### Secret Doors
- [?] Secret door to the north, blends in with the stonework, DC 15 Perception check.
- [?] Secret door to the north-east, blends in with the stonework, DC 15 Perception check.
- [?] Secret door to the south-east, blends in with the stonework, DC 15 Perception check.
- [?] Secret door to the south, blends in with the stonework, DC 15 Perception check.

- [?] Secret door on the north-east pillar, blends in with the stonework, DC 15 Perception check.
	- [?]  If opened, 15 [[10-gp-gemstones]] are found. 

### E6. Partially Collapsed Room

- [>] An earthquake collapsed part of this room, which holds nothing of value.
### E7. Secret Tunnel

- [?] Trapped behind the secret doors, an [ochre jelly](/3-Mechanics/CLI/bestiary/ooze/ochre-jelly.md) lurks in the rubble at the south end of this hall.

```encounter
name: Warn the Dwarves (Location E7)
creatures:
  - "1": Ochre Jelly
  - "1": Norbus Ironrune, friendly
  - "1": Dazyln Grayshard, friendly
```

### E8. Preists' Bedchamber

- [>] Three stone bed frames stand against the east wall.

### E9. Vestry

- [>] An empty stone font juts out of the southwest wall.
- [>] A stone wardrobe against the east wall holds the rotted remains of two suits of red leather armor (Now ruined and worthless)

### E10. Partially Collapsed Room

- [>] Half-buried in the rubble here is the skeleton of a dwarf priest wearing rotted red leather armor. The dwarf was killed when part of the room collapsed.
- [?] Around the skeleton's neck hangs a [[Holy Symbol of Abbathor]].

### E11. Hall of Greed

- [!] To reach this room, characters must dig through the rubble for **40 hours**.

- [>]  An alcove in the south wall holds the rubble of a shattered statue. 
- [>] An alcove to the north holds a statue of a dwarf with horns, who stares greedily at a **glowing green gem** in its hands.
- [?] If the gem is taken, it turns to dust and the statue explodes.  
	- [?] Any creature within 10 feet must make a DC 15, Dexterity saving throw taking `dice: 4d10` piercing damage on a failed save, half on a successful save.