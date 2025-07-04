---
obsidianUIMode: preview
cssclass: json5e-monster
statblock: inline
tags:
- compendium/src/5e/mm
- monster/cr/8
- monster/environment/underdark
- monster/size/large
- monster/type/aberration
aliases: ["Cloaker"]
NoteIcon: monster
BestiaryType: aberration
SourceType: Bestiary
BookSource: Monster Manual p. 41. Available in the SRD.
---
# [Cloaker](3-Mechanics\CLI\bestiary\aberration/cloaker.md)
*Source: Monster Manual p. 41. Available in the SRD.*  

Cloakers earned their names for the resemblance they bear to dark leathery cloaks. Lurking in remote dungeons and caves, these stealthy predators wait to slay lone or injured prey stumbling through the darkness.

## Camouflaged Lurkers

Like a stingray, a cloaker's body is composed of cartilage and muscle. With its tail and fins unfurled, it flies through darkness and lurks among the shadows of caverns the same way a stingray glides through water and hides on the ocean floor. Parallel rows of round, black eyespots run along its back like buttons, and the ivory-colored claws on its cowl resemble bone clasps.

When a cloaker unfurls and moves to attack, it reveals its pale underside and makes its true nature evident. Red eyes glow above rows of sharp teeth, and a long pendulous tail whips behind it.

## Opportunistic Predators

When hunting, cloakers glide through the shadows at a safe distance behind groups of other creatures traversing the Underdark. They follow parties of humanoids to prey on the wounded after a battle, or pursue herds of Underdark beasts, attacking the sick, the weak, or the straggling.

Cloakers strike quickly and consume their meals as swiftly as possible, enveloping and devouring their victims. While it feeds, a cloaker uses its swift, whiplike tail for defense, although it rarely takes a stand against dangerous foes or groups of creatures. As an added defense, cloakers can create illusory duplicates of themselves.

## Haunting Moan

Cloakers' thoughts are alien to other life-forms, and they communicate with one another through subsonic moans inaudible to most creatures. At higher intensities, a cloaker's haunting moan becomes audible, evoking sensations of doom and dread in creatures that hear it.

## Cloaker Conclaves

Cloakers prefer isolation, but they sometimes convene with other cloakers for defense or to exchange information about new dangers, suitable hunting grounds, or developments that might affect their habitats. When this convergence is complete, the cloakers separate again.

```statblock
"name": "Cloaker"
"size": "Large"
"type": "aberration"
"alignment": "Chaotic Neutral"
"ac": !!int "14"
"ac_class": "natural armor"
"hp": !!int "78"
"hit_dice": "12d10 + 12"
"stats":
- !!int "17"
- !!int "15"
- !!int "12"
- !!int "13"
- !!int "12"
- !!int "14"
"speed": "10 ft., fly 40 ft."
"skillsaves":
  "Stealth": !!int "5"
"senses": "darkvision 60 ft., passive Perception 11"
"languages": "Deep Speech, Undercommon"
"cr": "8"
"traits":
- "desc": "While attached to a creature, the cloaker takes only half the damage dealt\
    \ to it (rounded down). and that creature takes the other half."
  "name": "Damage Transfer"
- "desc": "While the cloaker remains motionless without its underside exposed, it\
    \ is indistinguishable from a dark leather cloak."
  "name": "False Appearance"
- "desc": "While in bright light, the cloaker has disadvantage on attack rolls and\
    \ Wisdom ([Perception](/3-Mechanics/CLI/rules/skills.md#Perception)) checks that\
    \ rely on sight."
  "name": "Light Sensitivity"
"actions":
- "desc": "The cloaker makes two attacks: one with its bite and one with its tail."
  "name": "Multiattack"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 5 ft., one creature.\
    \ Hit: dice:2d6 + 3|text(10) (2d6 + 3) piercing damage, and if the target\
    \ is Large or smaller, the cloaker attaches to it. If the cloaker has advantage\
    \ against the target, the cloaker attaches to the target's head, and the target\
    \ is [blinded](/3-Mechanics/CLI/rules/conditions.md#blinded) and unable to breathe\
    \ while the cloaker is attached. While attached, the cloaker can make this attack\
    \ only against the target and has advantage on the attack roll. The cloaker can\
    \ detach itself by spending 5 feet of its movement. A creature, including the\
    \ target, can take its action to detach the cloaker by succeeding on a DC 16 Strength\
    \ check."
  "name": "Bite"
- "desc": "Melee Weapon Attack: dice: d20+6 (+6) to hit, reach 10 ft., one creature.\
    \ Hit: dice:1d8 + 3|text(7) (1d8 + 3) slashing damage."
  "name": "Tail"
- "desc": "Each creature within 60 feet of the cloaker that can hear its moan and\
    \ that isn't an aberration must succeed on a DC 13 Wisdom saving throw or become\
    \ [frightened](/3-Mechanics/CLI/rules/conditions.md#frightened) until the end\
    \ of the cloaker's next turn. If a creature's saving throw is successful, the\
    \ creature is immune to the cloaker's moan for the next 24 hours."
  "name": "Moan"
- "desc": "The cloaker magically creates three illusory duplicates of itself if it\
    \ isn't in bright light. The duplicates move with it and mimic its actions, shifting\
    \ position so as to make it impossible to track which cloaker is the real one.\
    \ If the cloaker is ever in an area of bright light, the duplicates disappear.\n\
    \nWhenever any creature targets the cloaker with an attack or a harmful spell\
    \ while a duplicate remains, that creature rolls randomly to determine whether\
    \ it targets the cloaker or one of the duplicates. A creature is unaffected by\
    \ this magical effect if it can't see or if it relies on senses other than sight.\n\
    \nA duplicate has the cloaker's AC and uses its saving throws. If an attack hits\
    \ a duplicate, or if a duplicate fails a saving throw against an effect that deals\
    \ damage, the duplicate disappears."
  "name": "Phantasms (Recharges after a Short or Long Rest)"
"source":
- "MM"
- "PotA"
- "WDMM"
- "EGW"
- "CRCotN"
- "JttRC"
- "KftGV"
- "PaBTSO"
- "DoDk"
- "VEoR"
- "QftIS"
"image": "https://raw.githubusercontent.com/5etools-mirror-2/5etools-img/main/bestiary/tokens/MM/Cloaker.webp"
```
^statblock

```encounter-table
name: Cloaker
creatures:
 - 1: Cloaker
```

## Environment

underdark