---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/coastal
- monster/environment/underwater
- monster/size/medium
- monster/type/monstrosity
aliases: ["Sea Spawn"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 211, Volo's Guide to Monsters p. 189
---
# [Sea Spawn](3-Mechanics\CLI\bestiary\monstrosity/sea-spawn-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 211, Volo's Guide to Monsters p. 189*  

```statblock
"name": "Sea Spawn (MPMM)"
"size": "Medium"
"type": "monstrosity"
"alignment": "Typically  Neutral Evil"
"ac": !!int "11"
"ac_class": "natural armor"
"hp": !!int "32"
"hit_dice": "5d8 + 10"
"stats":
- !!int "15"
- !!int "8"
- !!int "15"
- !!int "6"
- !!int "10"
- !!int "8"
"speed": "20 ft., swim 30 ft."
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "understands Aquan and Common but can't speak"
"cr": "1"
"traits":
- "desc": "The sea spawn can breathe air and water, but it needs to be submerged in\
    \ the sea at least once a day for 1 minute to avoid suffocating."
  "name": "Limited Amphibiousness"
"actions":
- "desc": "The sea spawn makes two Unarmed Strike attacks and one Piscine Anatomy\
    \ attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 2|text(4) (1d4 + 2) bludgeoning damage."
  "name": "Unarmed Strike"
- "desc": "The sea spawn uses one of the following options (choose one or roll a dice:\
    \ d6|avg|noform (d6)):\n\n- 1–2 Bite. Melee Weapon Attack: dice: d20+4\
    \ (+4) to hit, reach 5 ft., one target. Hit: dice:1d4 + 2|text(4) (1d4\
    \ + 2) piercing damage.  \n- 3–4 Poison Quills. Melee Weapon Attack: dice:\
    \ d20+4 (+4) to hit, reach 5 ft., one creature. Hit: dice:1d6|text(3) (1d6)\
    \ poison damage, and the target must succeed on a DC 12 Constitution saving throw\
    \ or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) for 1 minute.\
    \ The target can repeat the saving throw at the end of each of its turns, ending\
    \ the effect on itself on a success.  \n- 5–6 Tentacle. Melee Weapon Attack:\
    \ dice: d20+4 (+4) to hit, reach 10 ft., one target. Hit: dice:1d6 + 2|text(5)\
    \ (1d6 + 2) bludgeoning damage, and the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 12) if it is a Medium or smaller creature. Until this grapple ends,\
    \ the sea spawn can't use this tentacle on another target.  "
  "name": "Piscine Anatomy"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Sea%20Spawn.webp"
```
^statblock

```encounter-table
name: Sea Spawn
creatures:
 - 1: Sea Spawn
```

## Environment

coastal, underwater