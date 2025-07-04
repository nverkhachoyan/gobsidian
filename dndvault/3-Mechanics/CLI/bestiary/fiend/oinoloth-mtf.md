---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/12
- monster/environment/desert
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend/yugoloth
aliases: ["Oinoloth"]
NoteIcon: monster
BestiaryType: fiend (yugoloth)
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 251
---
# [Oinoloth](3-Mechanics\CLI\bestiary\fiend/oinoloth-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 251*  

## Oinoloth

Grim specters of death, oinoloths bring pestilence wherever they go. To armies who recognize their awful forms, their mere appearance causes soldiers to break ranks and flee, lest they succumb to one of the awful plagues that oinoloths let loose.

Oinoloths provide the ultimate solution to thorny problems, usually by killing everyone involved. They are hired as a last resort, when a siege has gone on too long or an army has proved too strong to overcome. Once summoned, oinoloths stalk the killing field, poisoning the ground and sickening creatures they encounter. Sometimes they might be hired to lift the very plagues they spread, but the price for such work is high, and the effort turns the creatures they save into debilitated wrecks.

## Yugoloths

Mercenaries that ply their trade throughout the Lower Planes and in other realms, yugoloths have a reputation for effectiveness that is matched only by their desire for ever more wealth. Although yugoloths aren't especially loyal and typically try to exploit every potential loophole in a contract, they undertake any task for which they are hired, no matter how despicable. Yugoloths come in a wide variety of forms, including those described in the Monster Manual and the six creatures presented here.

```statblock
"name": "Oinoloth (MTF)"
"size": "Medium"
"type": "fiend"
"subtype": "yugoloth"
"alignment": "Neutral Evil"
"ac": !!int "17"
"ac_class": "natural armor"
"hp": !!int "126"
"hit_dice": "12d10 + 60"
"stats":
- !!int "19"
- !!int "17"
- !!int "18"
- !!int "17"
- !!int "16"
- !!int "19"
"speed": "40 ft."
"saves":
  "Wisdom": !!int "7"
  "Constitution": !!int "8"
"skillsaves":
  "Intimidation": !!int "8"
  "Deception": !!int "8"
  "Perception": !!int "7"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "acid, poison"
"condition_immunities": "[poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "blindsight 60 ft., darkvision 60 ft., passive Perception 17"
"languages": "Abyssal, Infernal, telepathy 60 ft."
"cr": "12"
"traits":
- "desc": "The oinoloth's innate spellcasting ability is Charisma (spell save DC 16).\
    \ It can innately cast the following spells, requiring no material components:\n\
    \nAt will: [darkness](/3-Mechanics/CLI/spells/darkness.md), [detect magic](/3-Mechanics/CLI/spells/detect-magic.md),\
    \ [dispel magic](/3-Mechanics/CLI/spells/dispel-magic.md), [invisibility](/3-Mechanics/CLI/spells/invisibility.md)\
    \ (self only)\n\n1/day each: [feeblemind](/3-Mechanics/CLI/spells/feeblemind.md),\
    \ [globe of invulnerability](/3-Mechanics/CLI/spells/globe-of-invulnerability.md),\
    \ [wall of fire](/3-Mechanics/CLI/spells/wall-of-fire.md), [wall of ice](/3-Mechanics/CLI/spells/wall-of-ice.md)"
  "name": "Innate Spellcasting"
- "desc": "As a bonus action, the oinoloth blights the area within 30 feet of it.\
    \ The blight lasts for 24 hours. While blighted, all normal plants in the area\
    \ wither and die, and the number of hit points restored by a spell to a creature\
    \ in that area is halved.\n\nFurthermore, when a creature moves into the blighted\
    \ area or starts its turn there, that creature must make a DC 16 Constitution\
    \ saving throw. On a successful save, the creature is immune to the oinoloth's\
    \ Bringer of Plagues for the next 24 hours. On a failed save, the creature takes\
    \ dice:4d6|text(14) (4d6) necrotic damage and is [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned).\n\
    \nThe [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned) creature can't\
    \ regain hit points. After every 24 hours that elapse, the [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)\
    \ creature can repeat the saving throw. On a failed save, the creature's hit point\
    \ maximum is reduced by dice:1d10|text(5) (1d10). This reduction lasts until\
    \ the poison ends, and the target dies if its hit point maximum is reduced to\
    \ 0. The poison ends after the creature successfully saves against it three times."
  "name": "Bringer of Plagues (Recharge 5-6)"
- "desc": "The oinoloth has advantage on saving throws against spells and other magical\
    \ effects."
  "name": "Magic Resistance"
- "desc": "The oinoloth's weapon attacks are magical."
  "name": "Magic Weapons"
"actions":
- "desc": "The oinoloth uses its Transfixing Gaze and makes two claw attacks."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+8 (+8) to hit, reach 5 ft., one target.\
    \ Hit: dice:3d6 + 4|text(14) (3d6 + 4) slashing damage plus dice:4d10|text(22)\
    \ (4d10) necrotic damage."
  "name": "Claw"
- "desc": "The oinoloth touches one willing creature within 5 feet of it. The target\
    \ regains all its hit points. In addition, the oinoloth can end one disease on\
    \ the target or remove one of the following conditions from it: [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded),\
    \ [deafened](/3-Mechanics/CLI/rules/conditions.md#deafened), [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed),\
    \ or [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned). The target then\
    \ gains 1 level of [exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
    \ and its hit point maximum is reduced by dice:2d6|text(7) (2d6). This reduction\
    \ can be removed only by a [wish](/3-Mechanics/CLI/spells/wish.md) spell or by\
    \ casting [greater restoration](/3-Mechanics/CLI/spells/greater-restoration.md)\
    \ on the target three times within the same hour. The target dies if its hit point\
    \ maximum is reduced to 0."
  "name": "Corrupted Healing (Recharge 6)"
- "desc": "The oinoloth magically teleports, along with any equipment it is wearing\
    \ or carrying, up to 60 feet to an unoccupied space it can see."
  "name": "Teleport"
- "desc": "The oinoloth targets one creature it can see within 30 feet of it. The\
    \ target must succeed on a DC 16 Wisdom saving throw against this magic or be\
    \ [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed) until the end of the\
    \ oinoloth's next turn. While [charmed](/3-Mechanics/CLI/rules/conditions.md#charmed)\
    \ in this way, the target is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ If the target's saving throw is successful, the target is immune to the oinoloth's\
    \ gaze for the next 24 hours."
  "name": "Transfixing Gaze"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Oinoloth.webp"
```
^statblock

```encounter-table
name: Oinoloth
creatures:
 - 1: Oinoloth
```

## Environment

desert, underdark