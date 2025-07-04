---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vrgr
- monster/cr/1-2
- monster/size/medium
- monster/type/plant
aliases: ["Podling"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Van Richten's Guide to Ravenloft p. 227
---
# [Podling](3-Mechanics\CLI\bestiary\plant/podling-vrgr.md)
*Source: Van Richten's Guide to Ravenloft p. 227*  

Bodytaker plants either capture unsuspecting victims with their vines or accept captives brought to them by their podling servants. In either case, they drag creatures into their central pod, where potent chemicals render the captive comatose. Over the course of hours, the creature is dissolved and its body repurposed into a podling duplicate.

Podlings are near-perfect mimics of the creatures they replace. Despite having the knowledge of those they mimic, podlings frequently miss the nuances of interactions between sapient beings. These duplicates make excuses about their odd behavior, but those familiar with an individual replaced by a podling can often tell something's amiss. Roll on the Podling Behavior table to see what unusual habits a podling might demonstrate.

**Podling Behavior**

`dice: [](podling-vrgr.md#^podling-behavior)`

| dice: d6 | Behavior |
|----------|----------|
| 1 | The podling abandons the habits and hobbies of the creature it's mimicking. |
| 2 | The podling lavishes affection on plants and views houseplants with excessive sympathy. |
| 3 | The podling relishes exposing its skin to the sun. It resents clothing and hair. |
| 4 | The podling often reacts as if some unseen force is speaking to it, staring into the distance or nodding. |
| 5 | The podling often communicates to other podlings using Deep Speech. |
| 6 | The podling no longer understands the nuances of a relationship, system, or instrument. |
^podling-behavior

```statblock
"name": "Podling (VRGR)"
"size": "Medium"
"type": "plant"
"alignment": "Unaligned"
"ac": !!int "10"
"hp": !!int "26"
"hit_dice": "4d8 + 8"
"stats":
- !!int "15"
- !!int "11"
- !!int "14"
- !!int "10"
- !!int "10"
- !!int "10"
"speed": "20 ft."
"condition_immunities": "[charmed](/3-Mechanics/CLI/rules/conditions.md#charmed),\
  \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)"
"senses": "blindsight 30 ft., passive Perception 10"
"languages": "Deep Speech, the languages the creature knew in life"
"cr": "1/2"
"traits":
- "desc": "The podling is a physical copy of a creature digested by a bodytaker plant.\
    \ The podling has the digested creature's memories and behaves like that creature,\
    \ but with occasional lapses. An observer familiar with the digested creature\
    \ can recognize the discrepancies with a successful DC 20 Wisdom ([Insight](/3-Mechanics/CLI/rules/skills.md#Insight))\
    \ check, or automatically if the podling does something in direct contradiction\
    \ to the digested creature's established beliefs or behavior. The podling melts\
    \ into a slurry when it dies, when the bodytaker plant that created it dies, or\
    \ when the bodytaker plant dismisses it (no action required)."
  "name": "Semblance of Life"
- "desc": "The podling doesn't require sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) bludgeoning damage."
  "name": "Slam"
"source":
- "VRGR"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VRGR/Podling.webp"
```
^statblock

```encounter-table
name: Podling
creatures:
 - 1: Podling
```