---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-2
- monster/size/large
- monster/type/undead
aliases: ["Warhorse Skeleton"]
NoteIcon: monster
BestiaryType: undead
SourceType: Bestiary
BookSource: Monster Manual p. 273. Available in the SRD.
---
# [Warhorse Skeleton](3-Mechanics\CLI\bestiary\undead/warhorse-skeleton.md)
*Source: Monster Manual p. 273. Available in the SRD.*  

Skeletons arise when animated by dark magic. They heed the summons of spellcasters who call them from their stony tombs and ancient battlefields, or rise of their own accord in places saturated with death and loss, awakened by stirrings of necromantic energy or the presence of corrupting evil.

## Animated Dead

Whatever sinister force awakens a skeleton infuses its bones with a dark vitality, adhering joint to joint and reassembling dismantled limbs. This energy motivates a skeleton to move and think in a rudimentary fashion, though only as a pale imitation of the way it behaved in life. An animated skeleton retains no connection to its past, although resurrecting a skeleton restores it body and soul, banishing the hateful undead spirit that empowers it.

While most skeletons are the animated remains of dead humans and other humanoids, skeletal undead can be created from the bones of other creatures besides humanoids, giving rise to a host of terrifying and unique forms.

## Obedient Servants

Skeletons raised by spell are bound to the will of their creator. They follow orders to the letter, never questioning the tasks their masters give them, regardless of the consequences. Because of their literal interpretation of commands and unwavering obedience, skeletons adapt poorly to changing circumstances. They can't read, speak, emote, or communicate in any way except to nod, shake their heads, or point. Still, skeletons are able to accomplish a variety of relatively complex tasks.

A skeleton can fight with weapons and wear armor, can load and fire a catapult or trebuchet, scale a siege ladder, form a shield wall, or dump boiling oil. However, it must receive careful instructions explaining how such tasks are accomplished.

Although they lack the intellect they possessed in life, skeletons aren't mindless. Rather than break its limbs attempting to batter its way through an iron door, a skeleton tries the handle first. If that doesn't work, it searches for another way through or around the obstacle.

## Habitual Behaviors

Independent skeletons temporarily or permanently free of a master's control sometimes pantomime actions from their past lives, their bones echoing the rote behaviors of their former living selves. The skeleton of a miner might lift a pick and start chipping away at stone walls. The skeleton of a guard might strike up a post at a random doorway. The skeleton of a dragon might lie down on a pile of treasure, while the skeleton of a horse crops grass it can't eat. Left alone in a ballroom, the skeletons of nobles might continue an eternally unfinished dance.

When skeletons encounter living creatures, the necromantic energy that drives them compels them to kill unless they are commanded by their masters to refrain from doing so. They attack without mercy and fight until destroyed, for skeletons possess little sense of self and even less sense of self-preservation.

## Undead Nature

A skeleton doesn't require air, food, drink, or sleep.

```statblock
"name": "Warhorse Skeleton"
"size": "Large"
"type": "undead"
"alignment": "Lawful Evil"
"ac": !!int "13"
"ac_class": "barding scraps"
"hp": !!int "22"
"hit_dice": "3d10 + 6"
"stats":
- !!int "18"
- !!int "12"
- !!int "15"
- !!int "2"
- !!int "8"
- !!int "5"
"speed": "60 ft."
"damage_vulnerabilities": "bludgeoning"
"damage_immunities": "poison"
"condition_immunities": "[exhaustion](/3-Mechanics/CLI/rules/conditions.md#exhaustion),\
  \ [poisoned](/3-Mechanics/CLI/rules/conditions.md#poisoned)"
"senses": "darkvision 60 ft., passive Perception 9"
"languages": ""
"cr": "1/2"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one target.\
    \ Hit: dice:2d6 + 4|text(11) (2d6 + 4) bludgeoning damage."
  "name": "Hooves"
"source":
- "MM"
- "CoS"
- "SKT"
- "CM"
- "DSotDQ"
- "BMT"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Warhorse%20Skeleton.webp"
```
^statblock

```encounter-table
name: Warhorse Skeleton
creatures:
 - 1: Warhorse Skeleton
```