---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/desert
- monster/environment/swamp
- monster/environment/underdark
- monster/environment/urban
- monster/size/medium
- monster/type/undead
aliases: ["Wight"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 300, Dragon of Icespire Peak, Sleeping Dragon's Wake. Available in the SRD and the Basic Rules.
---
# [Wight](3-Mechanics\CLI\bestiary\undead/wight.md)
*Source: Monster Manual p. 300, Dragon of Icespire Peak, Sleeping Dragon's Wake. Available in the SRD and the Basic Rules.*  

The word "wight" meant "person" in days of yore, but the name now refers to evil undead who were once mortals driven by dark desire and great vanity. When death stills such a creature's heart and snuffs its living breath, its spirit cries out to the demon lord Orcus or some vile god of the underworld for a reprieve: undeath in return for eternal war on the living. If a dark power answers the call, the spirit is granted undeath so that it can pursue its own malevolent agenda.

Wights possess the memories and drives of their formerly living selves. They will heed the call of whatever dark entity transformed them into undead, swearing oaths to appease their new lord while retaining their autonomy. Never tiring, a wight can pursue its goals relentlessly and without distraction.

## Life Eaters

Neither dead nor alive, a wight exists in a transitional state between one world and the next. The bright spark it possessed in life is gone, and in its place is a yearning to consume that spark in all living things.

When a wight attacks, this life essence glows like white-hot embers to its dark eyes, and the wight's cold touch can drain the spark through flesh, clothing, and armor.

## Shadow of the Grave

Wights flee from the world by day, away from the light of the sun, which they hate. They retreat to barrow mounds, crypts, and tombs where they dwell. Their lairs are silent, desolate places, surrounded by dead plants, noticeably blackened, and avoided by bird and beast.

Humanoids slain by a wight can rise as zombies under its control. Motivated by hunger for living souls and driven by the same desire for power that awakened them in undeath, some wights serve as shock troops for evil leaders, including wraiths. As soldiers, they are able to plan but seldom do so, relying on their hunger for destruction to overwhelm any creature that stands before them.

## Undead Nature

A wight doesn't require air, food, drink, or sleep.

```statblock
"name": "Wight"
"size": "Medium"
"type": "undead"
"alignment": "Neutral Evil"
"ac": !!int "14"
"ac_class": "[studded leather](/3-Mechanics/CLI/items/studded-leather-armor.md)"
"hp": !!int "45"
"hit_dice": "6d8 + 18"
"stats":
- !!int "15"
- !!int "14"
- !!int "16"
- !!int "10"
- !!int "13"
- !!int "15"
"speed": "30 ft."
"skillsaves":
  "Stealth": !!int "4"
  "Perception": !!int "3"
"damage_resistances": "necrotic; bludgeoning, piercing, slashing from nonmagical attacks\
  \ that aren't silvered"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 13"
"languages": "the languages it knew in life"
"cr": "3"
"traits":
- "desc": "While in sunlight, the wight has disadvantage on attack rolls, as well\
    \ as on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks\
    \ that rely on sight."
  "name": "Sunlight Sensitivity"
"actions":
- "desc": "The wight makes two longsword attacks or two longbow attacks. It can use\
    \ its Life Drain in place of one longsword attack."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one creature.\
    \ Hit: dice:1d6 + 2|text(5) (1d6 + 2) necrotic damage. The target must succeed\
    \ on a DC 13 Constitution saving throw or its hit point maximum is reduced by\
    \ an amount equal to the damage taken. This reduction lasts until the target finishes\
    \ a long rest. The target dies if this effect reduces its hit point maximum to\
    \ 0.\n\nA humanoid slain by this attack rises 24 hours later as a [zombie](/3-Mechanics/CLI/bestiary/undead/zombie.md)\
    \ under the wight's control, unless the humanoid is restored to life or its body\
    \ is destroyed. The wight can have no more than twelve zombies under its control\
    \ at one time."
  "name": "Life Drain"
- "desc": "Melee Weapon Attack: dice: d20+4 (+4) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 2|text(6) (1d8 + 2) slashing damage, or dice:1d10 + 2|text(7)\
    \ (1d10 + 2) slashing damage if used with two hands."
  "name": "Longsword"
- "desc": "Ranged Weapon Attack: dice: d20+4 (+4) to hit, range 150/600 ft.,\
    \ one target. Hit: dice:1d8 + 2|text(6) (1d8 + 2) piercing damage."
  "name": "Longbow"
"source":
- "MM"
- "CoS"
- "PotA"
- "RoT"
- "SKT"
- "TftYP"
- "ToA"
- "WDMM"
- "GoS"
- "DIP"
- "SDW"
- "BGDIA"
- "EGW"
- "IDRotF"
- "CM"
- "DSotDQ"
- "PSI"
- "PaBTSO"
- "AATM"
- "SatO"
- "BMT"
- "DoDk"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Wight.webp"
```
^statblock

```encounter-table
name: Wight
creatures:
 - 1: Wight
```

## Environment

underdark, swamp, urban, desert