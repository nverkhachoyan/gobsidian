---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mtf
- monster/cr/22
- monster/environment/desert
- monster/environment/forest
- monster/environment/grassland
- monster/environment/hill
- monster/environment/mountain
- monster/environment/underdark
- monster/size/gargantuan
- monster/type/elemental
aliases: ["Zaratan"]
NoteIcon: monster
BestiaryType: elemental
SourceType: Bestiary
BookSource: Mordenkainen's Tome of Foes p. 201
---
# [Zaratan](3-Mechanics\CLI\bestiary\elemental/zaratan-mtf.md)
*Source: Mordenkainen's Tome of Foes p. 201*  

## Zaratan

When a zaratan is summoned from the Elemental Plane of Earth, the ground rises up to take the shape of what looks like a hulking, armored reptile, its shell composed of the landscape from which it arose.

The zaratan plods across the land, each step sending shock waves through the ground severe enough to unsettle structures. Dim-witted, the zaratan lurches onward, expressing its rage through its trumpeting calls and the occasional boulder or blast of debris it spews from its cavernous maw. If seriously injured, the zaratan slowly retracts its appendages to gain shelter beneath its impervious shell, biding its time until it recovers and can resume its march.

## Elder Elementals

On their native planes, elementals sweep across the weird and tempestuous landscape. Some possess greater power, gained by feeding on their lesser kin and adding the essence of creatures they have devoured to their own until they become something extraordinary. When summoned, these elder elementals manifest as beings of apocalyptic capability, entities whose mere existence promises destruction.

## Deadly When Summoned

The methods for summoning elder elementals remain hidden in forbidden tomes or inscribed on the walls of lost temples raised to honor the Elder Elemental Eye. Only casters of superlative skill have even the faintest chance of calling forth one of these monsters, and the spellcaster is often destroyed by the effort. Thus, only the most unhinged and nihilistic members of Elemental Evil cults attempt such a summoning, in the hope of hastening the world toward some cataclysmic end.

## Elemental Nature

An elder elemental doesn't require air, food, drink, or sleep.

```statblock
"name": "Zaratan (MTF)"
"size": "Gargantuan"
"type": "elemental"
"alignment": "Neutral"
"ac": !!int "21"
"ac_class": "natural armor"
"hp": !!int "307"
"hit_dice": "15d20 + 150"
"stats":
- !!int "30"
- !!int "10"
- !!int "30"
- !!int "2"
- !!int "21"
- !!int "18"
"speed": "40 ft., swim 40 ft."
"saves":
  "Charisma": !!int "11"
  "Wisdom": !!int "12"
"damage_vulnerabilities": "thunder"
"damage_resistances": "cold; fire; lightning; bludgeoning, piercing, slashing from\
  \ nonmagical attacks"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [paralyzed](/3-Mechanics/CLI/rules/conditions.md#paralyzed), [petrified](/3-Mechanics/CLI/rules/conditions.md#petrified),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned), [stunned](/3-Mechanics/CLI/rules/conditions.md#stunned)"
"senses": "darkvision 60 ft., tremorsense 60 ft., passive Perception 15"
"languages": ""
"cr": "22"
"traits":
- "desc": "As a bonus action after moving at least 10 feet on the ground, the zaratan\
    \ can send a shock wave through the ground in a 120-foot-radius circle centered\
    \ on itself. That area becomes difficult terrain for 1 minute. Each creature on\
    \ the ground that is concentrating must succeed on a DC 25 Constitution saving\
    \ throw or the creature's [concentration](/3-Mechanics/CLI/rules/conditions.md#concentration)\
    \ is broken.\n\nThe shock wave deals 100 thunder damage to all structures in contact\
    \ with the ground in the area. If a creature is near a structure that collapses,\
    \ the creature might be buried; a creature within half the distance of the structure's\
    \ height must make a DC 25 Dexterity saving throw. On a failed save, the creature\
    \ takes dice:5d6|text(17) (5d6) bludgeoning damage, is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone),\
    \ and is trapped in the rubble. A trapped creature is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained),\
    \ requiring a successful DC 20 Strength ([Athletics](/3-Mechanics/CLI/rules/skills.md#Athletics))\
    \ check as an action to escape. Another creature within 5 feet of the buried creature\
    \ can use its action to clear rubble and grant advantage on the check. If three\
    \ creatures use their actions in this way, the check is an automatic success.\
    \ On a successful save, the creature takes half as much damage and doesn't fall\
    \ [prone](/3-Mechanics/CLI/rules/conditions.md#prone) or become trapped."
  "name": "Earth-Shaking Movement"
- "desc": "If the zaratan fails a saving throw, it can choose to succeed instead."
  "name": "Legendary Resistance (3/Day)"
- "desc": "The zaratan's weapon attacks are magical."
  "name": "Magic Weapons"
- "desc": "The elemental deals double damage to objects and structures (included in\
    \ Earth-Shaking Movement)."
  "name": "Siege Monster"
"actions":
- "desc": "The zaratan makes two attacks: one with its bite and one with its stomp."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 20 ft., one\
    \ target. Hit: dice:4d8 + 10|text(28) (4d8 + 10) piercing damage."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+17 (+17) to hit, reach 20 ft., one\
    \ target. Hit: dice:3d10 + 10|text(26) (3d10 + 10) bludgeoning damage."
  "name": "Stomp"
- "desc": "Ranged Weapon Attack: dice: d20+17 (+17) to hit, range 120/240 ft.,\
    \ one target. Hit: dice:6d8 + 10|text(31) (6d8 + 10) bludgeoning damage."
  "name": "Spit Rock"
- "desc": "The zaratan exhales rocky debris in a 90-foot cube. Each creature in that\
    \ area must make a DC 25 Dexterity saving throw. A creature takes dice:6d10|text(33)\
    \ (6d10) bludgeoning damage on a failed save, or half as much damage on a successful\
    \ one. A creature that fails the save by 5 or more is knocked [prone](/3-Mechanics/CLI/rules/conditions.md#prone)."
  "name": "Spew Debris (Recharge 5-6)"
"legendary_actions":
- "desc": "The zaratan makes one stomp attack."
  "name": "Stomp"
- "desc": "The zaratan moves up to its speed."
  "name": "Move"
- "desc": "The zaratan uses Spit Rock."
  "name": "Spit (Costs 2 Actions)"
- "desc": "The zaratan retracts into its shell. Until it takes its Emerge action,\
    \ it has resistance to all damage, and it is [restrained](/3-Mechanics/CLI/rules/conditions.md#restrained).\
    \ The next time it takes a legendary action, it must take its Revitalize or Emerge\
    \ action."
  "name": "Retract (Costs 2 Actions)"
- "desc": "The zaratan can use this option only if it is retracted in its shell. It\
    \ regains dice:5d20|text(52) (5d20) hit points. The next time it takes a legendary\
    \ action, it must take its Emerge action."
  "name": "Revitalize (Costs 2 Actions)"
- "desc": "The zaratan emerges from its shell and uses Spit Rock. It can use this\
    \ option only if it is retracted in its shell."
  "name": "Emerge (Costs 2 Actions)"
"source":
- "MTF"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MTF/Zaratan.webp"
```
^statblock

```encounter-table
name: Zaratan
creatures:
 - 1: Zaratan
```

## Environment

desert, forest, grassland, hill, mountain, underdark