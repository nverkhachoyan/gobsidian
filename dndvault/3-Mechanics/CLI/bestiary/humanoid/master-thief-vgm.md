---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/vgm
- monster/cr/5
- monster/environment/urban
- monster/size/medium
- monster/type/humanoid/any-race
aliases: ["Master Thief"]
NoteIcon: monster
BestiaryType: humanoid (any race)
SourceType: Bestiary
BookSource: Volo's Guide to Monsters p. 216, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake
---
# [Master Thief](3-Mechanics\CLI\bestiary\humanoid/master-thief-vgm.md)
*Source: Volo's Guide to Monsters p. 216, Dragon of Icespire Peak, Storm Lord's Wrath, Sleeping Dragon's Wake*  

Master thieves are known for perpetrating daring heists. They tend to develop a reputation and a cult of personality. A master thief might "retire" from hands-on work to run a thieves' guild, spearhead some covert enterprise, or enjoy a quiet life of luxury.

```statblock
"name": "Master Thief (VGM)"
"size": "Medium"
"type": "humanoid"
"subtype": "any race"
"alignment": "Any alignment"
"ac": !!int "16"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "83"
"hit_dice": "13d8 + 26"
"stats":
- !!int "11"
- !!int "18"
- !!int "14"
- !!int "11"
- !!int "11"
- !!int "12"
"speed": "30 ft."
"saves":
  "Dexterity": !!int "7"
  "Intelligence": !!int "3"
"skillsaves":
  "Athletics": !!int "3"
  "Sleight of Hand": !!int "7"
  "Stealth": !!int "7"
  "Perception": !!int "3"
  "Acrobatics": !!int "7"
"senses": "passive Perception 13"
"languages": "any one language (usually Common) plus Thieves' cant"
"cr": "5"
"traits":
- "desc": "On each of its turns, the thief can use a bonus action to take the Dash,\
    \ Disengage, or Hide action."
  "name": "Cunning Action"
- "desc": "If the thief is subjected to an effect that allows it to make a Dexterity\
    \ saving throw to take only half damage, the thief instead takes no damage if\
    \ it succeeds on the saving throw, and only half damage if it fails."
  "name": "Evasion"
- "desc": "The thief deals an extra dice:4d6|text(14) (4d6) damage when it hits\
    \ a target with a weapon attack and has advantage on the attack roll, or when\
    \ the target is within 5 feet of an ally of the thief that isn't [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)\
    \ and the thief doesn't have disadvantage on the attack roll."
  "name": "Sneak Attack (1/Turn)"
"actions":
- "desc": "The thief makes three attacks with its shortsword."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+7 (+7) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d6 + 4|text(7) (1d6 + 4) piercing damage."
  "name": "Shortsword"
- "desc": "Ranged Weapon Attack: dice: d20+7 (+7) to hit, range 80/320 ft.,\
    \ one target. Hit: dice:1d8 + 4|text(8) (1d8 + 4) piercing damage."
  "name": "Light Crossbow"
"reactions":
- "desc": "The thief halves the damage that it takes from an attack that hits it.\
    \ The thief must be able to see the attacker."
  "name": "Uncanny Dodge"
"source":
- "VGM"
- "DIP"
- "SLW"
- "SDW"
- "IMR"
- "MOT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/VGM/Master%20Thief.webp"
```
^statblock

```encounter-table
name: Master Thief
creatures:
 - 1: Master Thief
```

## Environment

urban