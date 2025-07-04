---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/10
- monster/environment/swamp
- monster/environment/underdark
- monster/size/huge
- monster/type/monstrosity
aliases: ["Froghemoth"]
NoteIcon: monster
BestiaryType: monstrosity
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 130, Volo's Guide to Monsters p. 145
---
# [Froghemoth](3-Mechanics\CLI\bestiary\monstrosity/froghemoth-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 130, Volo's Guide to Monsters p. 145*  

```statblock
"name": "Froghemoth (MPMM)"
"size": "Huge"
"type": "monstrosity"
"alignment": "Unaligned"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "161"
"hit_dice": "14d12 + 70"
"stats":
- !!int "23"
- !!int "13"
- !!int "20"
- !!int "2"
- !!int "12"
- !!int "5"
"speed": "30 ft., swim 30 ft."
"saves":
  "Wisdom": !!int "5"
  "Constitution": !!int "9"
"skillsaves":
  "Stealth": !!int "5"
  "Perception": !!int "9"
"damage_resistances": "fire, lightning"
"senses": "darkvision 60 ft., passive Perception 19"
"languages": ""
"cr": "10"
"traits":
- "desc": "The froghemoth can breathe air and water."
  "name": "Amphibious"
- "desc": "If the froghemoth takes lightning damage, it suffers two effects until\
    \ the end of its next turn: its speed is halved, and it has disadvantage on Dexterity\
    \ saving throws."
  "name": "Shock Susceptibility"
"actions":
- "desc": "The froghemoth makes one Bite attack and two Tentacle attacks, and it can\
    \ use Tongue."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:3d10 + 6|text(22) (3d10 + 6) piercing damage, and the\
    \ target is swallowed if it is a Medium or smaller creature. A swallowed creature\
    \ is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) and [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ has total cover against attacks and other effects outside the froghemoth, and\
    \ takes dice:3d6|text(10) (3d6) acid damage at the start of each of the froghemoth's\
    \ turns.\n\nThe froghemoth's gullet can hold up to two creatures at a time. If\
    \ the froghemoth takes 20 damage or more on a single turn from a creature inside\
    \ it, the froghemoth must succeed on a DC 20 Constitution saving throw at the\
    \ end of that turn or regurgitate all swallowed creatures, each of which falls\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone) in a space within 10 feet\
    \ of the froghemoth. If the froghemoth dies, any swallowed creature is no longer\
    \ [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained) by it and can\
    \ escape from the corpse using 10 feet of movement, exiting [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 20 ft., one\
    \ target. Hit: dice:3d8 + 6|text(19) (3d8 + 6) bludgeoning damage, and the\
    \ target is [grappled](/3-Mechanics/CLI/rules/conditions.md#grappled) (escape\
    \ DC 16) if it is a Huge or smaller creature. Until the grapple ends, the froghemoth\
    \ can't use this tentacle on another target. The froghemoth has four tentacles."
  "name": "Tentacle"
- "desc": "The froghemoth targets one Medium or smaller creature that it can see within\
    \ 20 feet of it. The target must make a DC 18 Strength saving throw. On a failed\
    \ save, the target is pulled into an unoccupied space within 5 feet of the froghemoth."
  "name": "Tongue"
"source":
- "MPMM"
- "VGM"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Froghemoth.webp"
```
^statblock

```encounter-table
name: Froghemoth
creatures:
 - 1: Froghemoth
```

## Environment

swamp, underdark