---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1-4
- monster/environment/mountain
- monster/environment/swamp
- monster/size/small
- monster/type/aberration
aliases: ["Star Spawn Grue"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 227, Mordenkainen's Tome of Foes p. 234
---
# [Star Spawn Grue](3-Mechanics\CLI\bestiary\aberration/star-spawn-grue-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 227, Mordenkainen's Tome of Foes p. 234*  

```statblock
"name": "Star Spawn Grue (MPMM)"
"size": "Small"
"type": "aberration"
"alignment": "Typically  Neutral Evil"
"ac": !!int "11"
"hp": !!int "17"
"hit_dice": "5d6"
"stats":
- !!int "6"
- !!int "13"
- !!int "10"
- !!int "9"
- !!int "11"
- !!int "6"
"speed": "30 ft."
"damage_immunities": "psychic"
"senses": "darkvision 60 ft., passive Perception 10"
"languages": "Deep Speech"
"cr": "1/4"
"traits":
- "desc": "Creatures within 20 feet of the grue that aren't Aberrations have disadvantage\
    \ on saving throws, as well as on attack rolls against creatures other than a\
    \ star spawn grue."
  "name": "Aura of Shrieks"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d4 + 1|text(6) (2d4 + 1) piercing damage, and the target must\
    \ succeed on a DC 10 Wisdom saving throw or attack rolls against it have advantage\
    \ until the start of the grue's next turn."
  "name": "Confounding Bite"
"source":
- "MPMM"
- "MTF"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Star%20Spawn%20Grue.webp"
```
^statblock

```encounter-table
name: Star Spawn Grue
creatures:
 - 1: Star Spawn Grue
```

## Environment

mountain, swamp