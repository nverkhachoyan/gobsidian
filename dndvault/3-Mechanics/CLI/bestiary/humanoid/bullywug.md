---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/1-4
- monster/environment/swamp
- monster/size/medium
- monster/type/humanoid/bullywug
aliases: ["Bullywug"]
NoteIcon: monster
BestiaryType: humanoid (bullywug)
SourceType: Bestiary
BookSource: Monster Manual p. 35
---
# [Bullywug](3-Mechanics\CLI\bestiary\humanoid/bullywug.md)
*Source: Monster Manual p. 35*  

Life as a bullywug is nasty, brutish, and wet. These frog-headed amphibious humanoids must stay constantly moist, dwelling in rainy forests, marshes, and damp caves. Always hungry and thoroughly evil, bullywugs overwhelm opponents with superior numbers when they can, but they flee from serious threats to search for easier prey.

Bullywugs have green, gray, or mottled yellow skin that shifts through shades of gray, green, and brown, allowing them to blend in with their surroundings. They wear crude armor and wield simple weapons, and can deliver a powerful bite to foes that press too close.

## Foul Aristocracy

Bullywugs consider themselves the right and proper rulers of the swamps. They follow an etiquette of sorts when dealing with outsiders and each other, subject to the whims and fancies of their leader-a self-styled lord of the muck. Bullywugs introduce themselves with grand-sounding titles, make great shows of bowing and debasing themselves before their superiors, and endlessly vie to win their superiors' favor.

A bullywug has two ways to advance among its kind. It can either murder its rivals, though it must take pains to keep its criminal deeds secret, or it can find a treasure or magic item and present it as tribute or a token of obeisance to its liege. A bullywug that murders its rivals without cunning is likely to be executed, so it's more common for bullywugs to stage raids against caravans and settlements, with the goal of securing precious baubles to impress their lords and win their good graces. Invariably, such fine goods are reduced to filthy tatters through abuse and neglect. Once a gift loses its sheen, a bullywug lord invariably demands that its subjects bring it more treasure as tribute.

## Unruly Diplomacy

Bullywugs love nothing more than lording over those who trespass on their territories. Their warriors attempt to capture intruders rather than simply slaying them.

Captives are dragged before the king or queen-a bullywug of unusually large size-and forced to beg for mercy. Bribes, treasure, and flattery can trick the bullywug ruler into letting its captives go, but not before it tries to impress its "guests" with the majesty of its treasure and its realm. Struck with a deep inferiority complex, bullywug lords fancy themselves as kings and queens, but desperately crave the fear and respect of outsiders.

## Amphibian Allies

Bullywugs speak a language that allows them to communicate over large areas by croaking like frogs. News of intruders or other events in the swamp spread within minutes across this crude communication system.

Simple concepts in the language are understandable to frogs and toads. Bullywugs use this capability to form strong bonds with giant frogs, which they train as guardians and hunters. Larger specimens are sometimes used as mounts as well. The frogs' ability to swallow creatures whole provides a bullywug hunting band an easy means of carrying prey back to their villages.

```statblock
"name": "Bullywug"
"size": "Medium"
"type": "humanoid"
"subtype": "bullywug"
"alignment": "Neutral Evil"
"ac": !!int "15"
"ac_class": "[hide armor](/3-Mechanics/CLI/items/hide-armor.md), [shield](/3-Mechanics/CLI/items/shield.md)"
"hp": !!int "11"
"hit_dice": "2d8 + 2"
"stats":
- !!int "12"
- !!int "12"
- !!int "13"
- !!int "7"
- !!int "10"
- !!int "7"
"speed": "20 ft., swim 40 ft."
"skillsaves":
  "Stealth": !!int "3"
"senses": "passive Perception 10"
"languages": "Bullywug"
"cr": "1/4"
"traits":
- "desc": "The bullywug can breathe air and water."
  "name": "Amphibious"
- "desc": "The bullywug can communicate simple concepts to frogs and toads when it\
    \ speaks in Bullywug."
  "name": "Speak with Frogs and Toads"
- "desc": "The bullywug has advantage on Dexterity ([Stealth](/3-Mechanics/CLI/rules/skills.md#Stealth))\
    \ checks made to hide in swampy terrain."
  "name": "Swamp Camouflage"
- "desc": "The bullywug's long jump is up to 20 feet and its high jump is up to 10\
    \ feet, with or without a running start."
  "name": "Standing Leap"
"actions":
- "desc": "The bullywug makes two melee attacks: one with its bite and one with its\
    \ spear."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d4 + 1|text(3) (1d4 + 1) bludgeoning damage."
  "name": "Bite"
- "desc": "Melee or Ranged Weapon Attack: dice: d20+3 (+3) to hit, reach 5 ft.\
    \ or range 20/60 ft., one target. Hit: dice:1d6 + 1|text(4) (1d6 + 1) piercing\
    \ damage, or dice:1d8 + 1|text(5) (1d8 + 1) piercing damage if used with two\
    \ hands to make a melee attack."
  "name": "Spear"
"source":
- "MM"
- "HotDQ"
- "WDMM"
- "GoS"
- "EGW"
- "WBtW"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Bullywug.webp"
```
^statblock

```encounter-table
name: Bullywug
creatures:
 - 1: Bullywug
```

## Environment

swamp