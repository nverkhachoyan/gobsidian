---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/1
- monster/environment/desert
- monster/environment/swamp
- monster/environment/underdark
- monster/size/tiny
- monster/type/fiend
aliases: ["Vargouille"]
NoteIcon: monster
BestiaryType: fiend
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 251, Volo's Guide to Monsters p. 195
---
# [Vargouille](3-Mechanics\CLI\bestiary\fiend/vargouille-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 251, Volo's Guide to Monsters p. 195*  

```statblock
"name": "Vargouille (MPMM)"
"size": "Tiny"
"type": "fiend"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "12"
"hp": !!int "18"
"hit_dice": "4d4 + 8"
"stats":
- !!int "6"
- !!int "14"
- !!int "14"
- !!int "4"
- !!int "7"
- !!int "2"
"speed": "5 ft., fly 40 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 8"
"languages": "understands Abyssal, Infernal, and any languages it knew before becoming\
  \ a vargouille but can't speak"
"cr": "1"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) piercing damage plus dice:3d6|text(10)\
    \ (3d6) poison damage."
  "name": "Bite"
- "desc": "The vargouille targets one [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ Humanoid within 5 feet of it. The target must succeed on a DC 12 Charisma saving\
    \ throw or become cursed. The cursed target loses 1 point of Charisma after each\
    \ hour, as its head takes on fiendish aspects. The curse doesn't advance while\
    \ the target is in sunlight or the area of a [daylight](/3-Mechanics/CLI/spells/daylight.md)\
    \ spell; don't count that time. When the cursed target's Charisma becomes 2, it\
    \ dies, and its head tears from its body and becomes a new vargouille. Casting\
    \ [remove curse](/3-Mechanics/CLI/spells/remove-curse.md), [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md),\
    \ or a similar spell on the target before the transformation is complete can end\
    \ the curse. Doing so undoes the changes made to the target by the curse."
  "name": "Abyssal Curse"
- "desc": "The vargouille shrieks. Each Humanoid and Beast within 30 feet of the vargouille\
    \ and able to hear it must succeed on a DC 12 Wisdom saving throw or be [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ of the vargouille until the end of the vargouille's next turn. While [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened)\
    \ in this way, a target is [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned).\
    \ If a target's saving throw is successful or the effect ends for it, the target\
    \ is immune to the Stunning Shriek of all vargouilles for 1 hour."
  "name": "Stunning Shriek (Recharge 5-6)"
"source":
- "MPMM"
- "VGM"
- "AATM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Vargouille.webp"
```
^statblock

```encounter-table
name: Vargouille
creatures:
 - 1: Vargouille
```

## Environment

desert, swamp, underdark