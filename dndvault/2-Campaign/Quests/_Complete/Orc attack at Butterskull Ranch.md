---
questObtained: 
questStatus: Not Started
questGiver: "[[Harbin Wester]]"
questLocationObtained: "[[03-butterskull-ranch|Butterskull Ranch]]"
questSessionObtained: 
questNotes: 
questLootAvail:
  - "[[mithral-armor|Mithral Chain Mail]]"
  - "[[gold-gp|65gp]]"
  - "[[silver-sp|145sp]]"
  - "[[copper-cp|220cp]]"
questLootEarned: 
NoteIcon: quest
obsidianUIMode: preview
tags:
  - quest
---

```leaflet
id: butterskull_ranch_map
lock: true
recenter: true
noScrollZoom: true
image: [[map_butterskull_ranch.webp]]
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
> "A tenday ago, Butterskull Ranch was attacked by orcs. They raided the farm, set fire to the barn and smithy, and captured Alfonse Kalazorn—known as Big Al. His ranch hands were killed, and the orcs now occupy the farmhouse. Your task is to travel to the ranch, rescue Big Al, and drive out the orcs. Big Al also offers a reward for returning his prized cow, Petunia."

### Quest Objective

- [!] Travel to [[Butterskull Ranch]].
- [!] Rescue [[Alfonse Kalazorn]].
- [?] Convince Alfonse to return to [[Phandalin]] or clear out the orcs.
- [?] Find and return [[Petunia the cow]].

### Rewards

- [[mithral-armor|Mithral Chain Mail]]
- 65gp, 145sp, 220cp
- Gratitude of [[Alfonse Kalazorn]]

### Travel to the Ranch

- [?] The fastest route is via the [[Triboar Trail]], approximately a 60-mile journey (2.5 days on foot).
- [?] In [[Conyberry]], the players may find three unsaddled [[riding-horse]], branded with "BAK".
- [?] On the way, they may encounter [[Petunia the cow]], who is calm and will follow them if treated well.

### On Arrival

> [!readaloud]  
> "Butterskull Ranch sprawls across a vast field, enclosed by a ramshackle wooden fence. The charred remains of a barn and smithy stand near a two-story farmhouse beside a pond. Pigs wander freely, and the scattered corpses of orcs and ranch hands attract buzzing flies. The scene is eerily quiet."

- [?] Scattered bodies of [[orc|Orc]] and ranch hands lie between the farmhouse and burned barn.
- [?] The farmhouse is occupied by orcs, three times the number of the party.

### Farmhouse Locations

#### B1. Kitchen

- [?] A butter churn, shelves of foodstuffs, and ale.
- [?] A skull-shaped wooden butter mold on a table.

#### B2. Empty Foyer

- [>] A creaky wooden staircase leads to the second floor.

#### B3. Dining Room

- [?] Two trestle tables with benches, decorated with cattle skulls.

#### B4. Downstairs Closet

- [>] Contains dinnerware.

#### B5. Common Room

- [>] Padded chairs and game tables with scattered [[three-dragon-ante-set]] cards and [[dragonchess-set]] pieces.

#### B6. Big Al's Bedroom

- [?] A large bed and cedar wardrobe.
- [?] DC 15 Perception to find a secret compartment containing a suit of [[mithral-armor|Mithral Chain Mail]].

#### B7. Ranch Hands' Bedrooms

- [?] Each room has two beds and footlockers with personal effects.

#### B8. Big Al’s Office

- [?] Papers detailing ten years of business transactions.
- [>] A small sack hidden under the paperwork contains 65gp, 145sp, and 220cp.

#### B9. Upstairs Closet

- [>] Contains cleaning supplies.

#### B10. Cold Storage Cellar

> [!readaloud]  
> "A wooden door leads to a stone staircase descending into a dimly lit cellar. Inside, a large figure sits tied to a chair with a burlap sack over his head. Wooden shelves line the walls, filled with skulls made of butter coated in wax."

- [!] [[Alfonse Kalazorn]] is tied up here, badly beaten (9 HP remaining).
- [!] It takes 1 minute to free him.
- [?] He lacks armor and weapons (AC 11).
- [>] If convinced (DC 10 Persuasion or Intimidation), he will return to Phandalin.
- [>] If given a weapon, he fights any remaining orcs.

### Rescuing Petunia

- [>] If players did not find [[Petunia the cow]] on the way, they can search the countryside.
- [>] Roll a `dice: d6` every hour of searching. Petunia is found on a roll of **6** (or 5-6 if searching on horseback).
- [?] Big Al offers his [[mithral-armor|Mithral Chain Mail]] as a reward for Petunia’s safe return.

### Orcs in the Farmhouse

```encounter
name: Encounter
creatures:
  - "1": Alphonse Kalazorn, friendly, hidden
  - "3": [[Orc, Orcs B1]]
  - "3": [[Orc, Orcs B3]]
  - "3": [[Orc, Orcs B10]]
  - "3": [[Orc, Orcs B5]]
```

- [!] The orcs are drinking and not expecting trouble but will fight to the death.