---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/3
- monster/environment/mountain
- monster/environment/underdark
- monster/size/medium
- monster/type/fiend
aliases: ["Hell Hound"]
NoteIcon: monster
BestiaryType: fiend
SourceType: Bestiary
BookSource: Monster Manual p. 182, Storm Lord's Wrath. Available in the SRD and the Basic Rules.
---
# [Hell Hound](3-Mechanics\CLI\bestiary\fiend/hell-hound.md)
*Source: Monster Manual p. 182, Storm Lord's Wrath. Available in the SRD and the Basic Rules.*  

Monstrous, fire-breathing fiends that take the form of powerful dogs, hell hounds are found on the battlefields of Acheron and throughout the Lower Planes. On the Material Plane, hell hounds are most commonly seen in service to devils, fire giants, and other evil creatures that use them as guard animals and companions.

## Burning Hunger

Hell hounds hunt in packs, feeding on any creature that appears edible. They avoid potentially dangerous foes in favor of targeting the weakest prey with their savage bite and fiery breath, demonstrating a relentless determination as they pursue that prey to the bitter end.

When hell hounds feed, the flesh they consume stokes the infernal fires that burn within them. When a hell hound dies, that fire consumes the creature's remains in a billowing eruption of smoke and blazing embers, leaving nothing behind but scorched tufts of black fur.

## Evil to the Core

Hell hounds are smarter than mundane beasts, and their lawful nature makes them good at following orders. However, a hell hound's evil disposition means that the creature can't be trained to be anything other than a ruthless killer. If a hell hound isn't allowed to indulge its malevolent hunger, it quickly abandons or turns against its master.

```statblock
"name": "Hell Hound"
"size": "Medium"
"type": "fiend"
"alignment": "Lawful Evil"
"ac": !!int "15"
"ac_class": "natural armor"
"hp": !!int "45"
"hit_dice": "7d8 + 14"
"stats":
- !!int "17"
- !!int "12"
- !!int "14"
- !!int "6"
- !!int "13"
- !!int "6"
"speed": "50 ft."
"skillsaves":
  "Perception": !!int "5"
"damage_immunities": "fire"
"senses": "darkvision 60 ft., passive Perception 15"
"languages": "understands Infernal but can't speak it"
"cr": "3"
"traits":
- "desc": "The hound has advantage on Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception))\
    \ checks that rely on hearing or smell."
  "name": "Keen Hearing and Smell"
- "desc": "The hound has advantage on an attack roll against a creature if at least\
    \ one of the hound's allies is within 5 feet of the creature and the ally isn't\
    \ [incapacitated](/3-Mechanics/CLI/rules/conditions.md#incapacitated)."
  "name": "Pack Tactics"
"actions":
- "desc": "Melee Weapon Attack: dice: d20+5 (+5) to hit, reach 5 ft., one target.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) piercing damage plus dice:2d6|text(7)\
    \ (2d6) fire damage."
  "name": "Bite"
- "desc": "The hound exhales fire in a 15-foot cone. Each creature in that area must\
    \ make a DC 12 Dexterity saving throw, taking dice:6d6|text(21) (6d6) fire\
    \ damage on a failed save, or half as much damage on a successful one."
  "name": "Fire Breath (Recharge 5-6)"
"source":
- "MM"
- "CoS"
- "PotA"
- "SKT"
- "TftYP"
- "ToA"
- "WDH"
- "WDMM"
- "SLW"
- "BGDIA"
- "IMR"
- "EGW"
- "MOT"
- "CM"
- "KftGV"
- "PSI"
- "SatO"
- "BMT"
- "HFStCM"
- "DoDk"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Hell%20Hound.webp"
```
^statblock

```encounter-table
name: Hell Hound
creatures:
 - 1: Hell Hound
```

## Environment

underdark, mountain