---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/2
- monster/environment/forest
- monster/environment/swamp
- monster/size/small
- monster/type/plant
aliases: ["Vegepygmy Chief"]
NoteIcon: monster
BestiaryType: plant
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 253, Volo's Guide to Monsters p. 197
---
# [Vegepygmy Chief](3-Mechanics\CLI\bestiary\plant/vegepygmy-chief-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 253, Volo's Guide to Monsters p. 197*  

```statblock
"name": "Vegepygmy Chief (MPMM)"
"size": "Small"
"type": "plant"
"alignment": "Typically  Neutral"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "33"
"hit_dice": "6d6 + 12"
"stats":
- !!int "14"
- !!int "14"
- !!int "14"
- !!int "7"
- !!int "12"
- !!int "9"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "3"
"damage_resistances": "lightning, piercing"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "Vegepygmy"
"cr": "2"
"traits":
- "desc": "The vegepygmy has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks it makes in any terrain with ample obscuring vegetation."
  "name": "Plant Camouflage"
- "desc": "The vegepygmy regains 5 hit points at the start of its turn. If it takes\
    \ cold, fire, or necrotic damage, this trait doesn't function at the start of\
    \ the vegepygmy's next turn. The vegepygmy dies only if it starts its turn with\
    \ 0 hit points and doesn't regenerate."
  "name": "Regeneration"
"actions":
- "desc": "The vegepygmy makes two Claw attacks or two melee Spear attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) slashing damage."
  "name": "Claws"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing\
    \ damage, or dice:1d8 + 2|text(6) (1d8 + 2) piercing damage if used with two\
    \ hands to make a melee attack."
  "name": "Spear"
- "desc": "A 15-foot-radius cloud of toxic spores extends out from the vegepygmy.\
    \ The spores spread around corners. Each creature in that area that isn't a Plant\
    \ must succeed on a DC 12 Constitution saving throw or be [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned).\
    \ While [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) in this way,\
    \ a target takes dice:2d8|text(9) (2d8) poison damage at the start of each\
    \ of its turns. A target can repeat the saving throw at the end of each of its\
    \ turns, ending the effect on itself on a success."
  "name": "Spores (1/Day)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Vegepygmy%20Chief.webp"
```
^statblock

```encounter-table
name: Vegepygmy Chief
creatures:
 - 1: Vegepygmy Chief
```

## Environment

forest, swamp