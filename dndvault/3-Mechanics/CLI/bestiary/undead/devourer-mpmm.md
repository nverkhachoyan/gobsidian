---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mpmm
- monster/cr/13
- monster/environment/underdark
- monster/size/large
- monster/type/undead
aliases: ["Devourer"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Mordenkainen Presents: Monsters of the Multiverse p. 93, Volo's Guide to Monsters p. 138
---
# [Devourer](3-Mechanics\CLI\bestiary\undead/devourer-mpmm.md)
*Source: Mordenkainen Presents: Monsters of the Multiverse p. 93, Volo's Guide to Monsters p. 138*  

```statblock
"name": "Devourer (MPMM)"
"size": "Large"
"type": "undead"
"alignment": "Typically  Chaotic Evil"
"ac": !!int "16"
"ac_class": "natural armor"
"hp": !!int "189"
"hit_dice": "18d10 + 90"
"stats":
- !!int "20"
- !!int "12"
- !!int "20"
- !!int "13"
- !!int "10"
- !!int "16"
"speed": "30 ft."
"damage_resistances": "cold, fire, lightning"
"damage_immunities": "poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 120 ft., passive Perception 10"
"languages": "Abyssal, telepathy 120 ft."
"cr": "13"
"traits":
- "desc": "A devourer doesn't require air, drink, or sleep."
  "name": "Unusual Nature"
"actions":
- "desc": "The devourer makes two Claw attacks and can use either Imprison Soul or\
    \ Soul Rend, if available."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+10 (+10) to hit, reach 5 ft., one\
    \ target. Hit: dice:2d6 + 5|text(12) (2d6 + 5) slashing damage plus dice:6d6|text(21)\
    \ (6d6) necrotic damage."
  "name": "Claw"
- "desc": "The devourer chooses a living Humanoid with 0 hit points that it can see\
    \ within 30 feet of it. That creature is teleported inside the devourer's ribcage\
    \ and imprisoned there. While imprisoned in this way, the creature is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained)\
    \ and has disadvantage on death saving throws. If the creature dies while imprisoned,\
    \ the devourer regains 25 hit points and immediately recharges Soul Rend. Additionally,\
    \ at the start of its next turn, the devourer regurgitates the slain creature\
    \ as a bonus action, and the creature becomes an undead. If the victim had 2 or\
    \ fewer Hit Dice, it becomes a [zombie](/3-Mechanics/CLI/bestiary/undead/zombie.md).\
    \ If it had 3 to 5 Hit Dice, it becomes a [ghoul](/3-Mechanics/CLI/bestiary/undead/ghoul.md).\
    \ Otherwise, it becomes a [wight](/3-Mechanics/CLI/bestiary/undead/wight.md).\
    \ A devourer can imprison only one creature at a time."
  "name": "Imprison Soul"
- "desc": "The devourer creates a vortex of life-draining energy in a 20-foot radius\
    \ centered on itself. Each creature in that area must make a DC 18 Constitution\
    \ saving throw, taking dice:8d10|text(44) (8d10) necrotic damage on a failed\
    \ save, or half as much damage on a successful one."
  "name": "Soul Rend (Recharge 6)"
"source":
- "MPMM"
- "VGM"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MPMM/Devourer.webp"
```
^statblock

```encounter-table
name: Devourer
creatures:
 - 1: Devourer
```

## Environment

underdark