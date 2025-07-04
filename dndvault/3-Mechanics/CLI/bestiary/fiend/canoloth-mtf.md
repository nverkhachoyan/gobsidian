---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/8
- monster/environment/coastal
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/fiend/yugoloth
aliases: ["Canoloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 247
---
# [Canoloth](3-Mechanics\CLI\bestiary\fiend/canoloth-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 247*  

> [!quote]- A quote from Mordenkainen  
> 
> Canoloths are fundamentally lazy creatures. Given no reason to attack, they rarely rise to the bait.

## Canoloth

Canoloths prefer to enter into contracts to guard valuable treasures and important locations. They always do exactly as asked—never any more, never any less.

With senses sharp enough to pinpoint the locations of nearby invisible creatures, canoloths respond unfailingly to any threat to their charges. Furthermore, they emit a magical distortion field that prevents creatures close to them from teleporting.

Canoloths confront intruders with swift and terrible force, projecting long, spiny tongues to grab their foes and drag them close. What happens next depends on the contract. Unless instructed to kill, a canoloth merely holds onto its prisoner, but if given the order to do so, it tears its prey limb from limb.

## Yugoloths

Mercenaries that ply their trade throughout the Lower Planes and in other realms, yugoloths have a reputation for effectiveness that is matched only by their desire for ever more wealth. Although yugoloths aren't especially loyal and typically try to exploit every potential loophole in a contract, they undertake any task for which they are hired, no matter how despicable. Yugoloths come in a wide variety of forms, including those described in the Monster Manual and the six creatures presented here.

```statblock
"name": "Canoloth (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Neutral Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "120"
"hit_dice": "16d8 + 48"
"stats":
- !!int "18"
- !!int "10"
- !!int "17"
- !!int "5"
- !!int "17"
- !!int "12"
"speed": "50 ft."
"skillsaves":
  "Investigation": !!int "3"
  "Perception": !!int "9"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., truesight 120 ft., passive Perception 19"
"languages": "Abyssal, Infernal, telepathy 60 ft."
"cr": "8"
"traits":
- "desc": "Other creatures can't teleport to or from a space within 60 feet of the\
    \ canoloth. Any attempt to do so is wasted."
  "name": "Dimensional Lock"
- "desc": "The canoloth has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The canoloth's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "The canoloth can't be [surprised](/3-Mechanics/CLI/rules/conditions.md#surprised)\
    \ while it isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Uncanny Senses"
"actions":
- "desc": "The canoloth makes two attacks: one with its tongue or its bite and one\
    \ with its claws."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:6d6 + 4|text(25) (6d6 + 4) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d10 + 4|text(15) (2d10 + 4) slashing damage."
  "name": "Claws"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 30 ft., one\
    \ target. Hit: dice:2d12 + 4|text(17) (2d12 + 4) piercing damage. If the\
    \ target is Medium or smaller, it is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled)\
    \ (escape DC 15), pulled up to 30 feet toward the canoloth, and is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ until the grapple ends. The canoloth can grapple one target at a time with its\
    \ tongue."
  "name": "Tongue"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Canoloth.webp"
```
^statblock

```encounter-table
name: Canoloth
creatures:
 - 1: Canoloth
```

## Environment

coastal, underdark, urban