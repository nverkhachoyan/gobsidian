---
questObtained:  
questStatus: In Progress  
questGiver: "[[Harbin Wester]]"  
questLocationObtained: "[[Loggers' Camp]]"  
questSessionObtained: "[[2024-12-10]]"  
questNotes:  
questLootAvail:  
  - "[[gold-gp|50gp]]"  
questLootEarned:  
  - "[[gold-gp|50gp]]"  
NoteIcon: quest  
obsidianUIMode: preview  
tags:  
  - quest  
---

```leaflet
id: loggers_camp_map
lock: true
recenter: true
noScrollZoom: true
image: [[map_loggers_camp.webp]]
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
> Type | Stat |  
> ---|---|  
> Date Obtained | `INPUT[datePicker:questObtained]` |  
> Status | `INPUT[inlineSelect(option(Not Started), option(In Progress), option(Complete)):questStatus]` |  
> Quest Giver | `INPUT[suggester(optionQuery(#npc)):questGiver]` |  
> Quest Location | `INPUT[suggester(optionQuery(#Category/Settlement)):questLocationObtained]` |  
> Session Obtained | `INPUT[suggester(optionQuery(#journal)):questSessionObtained]` |  
> Available Loot | `=this.questLootAvail` |  
> Acquired Loot | `=this.questLootEarned` |  

> [!readaloud]  
> "The logging camp along the river in Neverwinter Wood has gone silent. Harbin Wester needs adventurers to deliver supplies to the camp and have Tibor Wester sign for them. However, the camp is under attack by [ankhegs](/3-Mechanics/CLI/bestiary/monstrosity/ankheg.md), and Tibor has barricaded himself in his office."

### Quest Objective

- [!] Deliver supplies to [[Loggers' Camp]].  
- [!] Obtain [[Tibor Wester]]’s signature to confirm the delivery.  
- [?] Investigate the disappearance of the loggers.  
- [?] Eliminate any threats in the camp.  

### Rewards  

- 50gp from [[Harbin Wester]]  

### Preparations  

- [?] Supplies for the camp must be collected from [[Barthen's Provisions]] in [[Phandalin]].  
- [?] The supplies are packed into two heavy crates loaded onto a cart pulled by an ox named [[Vincent]].  
- [?] Players can buy a bottle of fine wine (10gp) as a gift for [[Falcon's Hunting Lodge]].  

### Travel to the Camp  

- [>] The journey is **50 miles north** of Phandalin, requiring **two days** of travel.  
- [?] Players must navigate Neverwinter Wood with a **DC 10 Survival check** for each hex entered. A failure wastes `dice: 1d4` miles of movement.  

#### A Boar-ing Encounter  

> [!readaloud]  
> "A wild boar stands in a clearing ahead, glaring at you suspiciously."  

- [!] The boar is actually a disguised [anchorite of Talos](/3-Mechanics/CLI/bestiary/humanoid/anchorite-of-talos-dip.md) named [[Drethna]].  
- [?] If players avoid a fight, Drethna escapes to warn the anchorites at the [[Woodland Manse]].  

### On Arrival  

> [!readaloud]  
> "The logging camp spreads along the river, with a dozen tents and a few cabins on the shore. A grim silence hangs over the camp, and no one is in sight."  

- [?] The [ankhegs](/3-Mechanics/CLI/bestiary/monstrosity/ankheg.md) are hidden underground and use [tremorsense](/3-Mechanics/CLI/rules/senses.md#tremorsense) to detect movement.  

### Camp Locations  

#### L1. Cabins on the Rocks  

- [?] These ruins sit atop a **10-foot escarpment** that the ankhegs **can't burrow through**.  

#### L2. Old Cabin and Chimney  

- [?] A ruined chimney holds a **totem**, buried under debris.  
- [!] Destroying the totem **prevents further ankheg attacks**.  
- [?] DC 15 Religion check identifies the totem’s **cursed purpose**.  

#### L3. Office and Tool Storage  

- [?] The **larger room** contains logging equipment.  
- [!] An [ankheg](/3-Mechanics/CLI/bestiary/monstrosity/ankheg.md) **bursts through the floorboards** when someone enters.  
- [?] [[Tibor Wester]], a **cowardly** [commoner](/3-Mechanics/CLI/bestiary/humanoid/commoner.md), is **barricaded** in the office.  

#### L4. River Dock  

- [>] **Rowboats and river barges** can dock here, but none are present.  

#### L5. North Camp  

- [?] Six abandoned tents **partially dissolved by acid**.  
- [?] A DC 10 Investigation check reveals **tracks leading into the sand** where **burrowing monsters** attacked.  

#### L6. South Camp  

- [?] Three **hidden ankhegs** lurk beneath the ground.  
- [!] When a player steps within **20 feet** of a marked spot, an ankheg bursts out and **attacks**.  

```encounter  
name: Defend Loggers' Camp  
creatures:  
  - "3": Ankhegs  
  - "1": Tibor Wester, friendly
```


```C
int main() {
	printf("Hello World!");
}
```
