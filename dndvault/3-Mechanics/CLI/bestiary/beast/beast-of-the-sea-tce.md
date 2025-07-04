---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/tce
- monster/cr/
- monster/size/medium
- monster/type/beast
aliases: ["Beast of the Sea"]
NoteIcon: monster
BestiaryType: beast
SourceType: Bestiary
BookSource: Tasha's Cauldron of Everything p. 61
---
# [Beast of the Sea](3-Mechanics\CLI\bestiary\beast/beast-of-the-sea-tce.md)
*Source: Tasha's Cauldron of Everything p. 61*  

```statblock
"name": "Beast of the Sea (TCE)"
"size": "Medium"
"type": "beast"
"alignment": "Unaligned"
"ac_class": "13 + PB (natural armor)"
"stats":
- !!int "14"
- !!int "14"
- !!int "15"
- !!int "8"
- !!int "14"
- !!int "11"
"speed": "5 ft., swim 60 ft."
"senses": "darkvision 60 ft., passive Perception 12"
"languages": "understands the languages you speak"
"traits":
- "desc": "The beast can breathe both air and water."
  "name": "Amphibious"
- "desc": "You can add your proficiency bonus to any ability check or saving throw\
    \ that the beast makes."
  "name": "Primal Bond"
"actions":
- "desc": "Melee Weapon Attack: the summoner's spell attack modifier to hit, reach\
    \ 5 ft., one target. Hit: 1d6 + 2 + PB piercing damage or 1d6 + 2 + PB bludgeoning\
    \ damage (your choice), and the target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC equal to your spellcasting save DC). Until this grapple ends, the\
    \ beast can't use this attack on another target."
  "name": "Binding Strike"
"source":
- "TCE"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/TCE/Beast%20of%20the%20Sea.webp"
```
^statblock

```encounter-table
name: Beast of the Sea
creatures:
 - 1: Beast of the Sea
```